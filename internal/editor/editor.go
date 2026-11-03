package editor

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"text/template"
	"valterbarros/blog-tools/internal/content"
	"valterbarros/blog-tools/internal/utils"
)

//go:embed web/assets/*
var staticFs embed.FS

//go:embed web/*.html
var templatesFS embed.FS

var templates = template.Must(template.ParseFS(templatesFS, "web/*.html"))

type Page struct {
	Title string
	Body  []byte
}

func renderTemplate(w http.ResponseWriter, tmpl string, p any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func editorHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	post, _ := utils.GetFromRoot("articles/_posts/" + name)

	d, err := utils.GetFileData(post)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	p := &Page{
		Title: "Editor",
		Body:  []byte(d),
	}

	renderTemplate(w, "editor", p)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	ps, err := content.ListPosts()

	if err != nil {
		http.NotFound(w, r)
		return
	}

	list := strings.Split(ps, "\n")

	renderTemplate(w, "index", list)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("sending imagem")

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "file is very big", http.StatusBadRequest)
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "file is ausent", http.StatusBadRequest)
	}

	defer file.Close()

	b, af := utils.GetDate()
	name := filepath.Base(b + af + header.Filename)

	basePath, _ := utils.GetFromRoot("articles/imgs")
	dstPath := strings.ReplaceAll(filepath.Join(basePath, name), " ", "-")

	// TODO: verify if that can be turn into a utils
	dst, errS := os.Create(dstPath)

	if errS != nil {
		http.Error(w, "Error when saving file", http.StatusBadRequest)
	}

	defer dst.Close()

	// TODO: verify if that can be turn into a utils
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error when copying file", http.StatusBadRequest)
	}

	encoder := json.NewEncoder(w)
	url := "{{\"" + filepath.Join("articles/imgs", filepath.Base(dstPath)) + "\" | relative_url }}"
	encoder.Encode(map[string]string{"url": url})
}

func RunServer() error {
	const PORT = "12769"

	signals := make(chan os.Signal, 1)
	// closes server on siging -9
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	http.HandleFunc("/ws", handleWebSocket)

	imgs, _ := utils.GetFromRoot("articles/imgs")
	http.Handle("/articles/imgs/", http.StripPrefix("/articles/imgs/", http.FileServer(http.Dir(imgs))))

	assets, errFs := fs.Sub(staticFs, "web/assets")
	if errFs != nil {
		return errFs
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets))))

	http.HandleFunc("/editor", editorHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", listHandler)

	err := http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		log.Fatal("Error to start server: ", err)
		return err
	}

	return nil
}
