package editor

import (
	"encoding/json"
	"fmt"
	"io"
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

var f, _ = utils.GetFromRoot("cli/web/editor.html")
var f2, _ = utils.GetFromRoot("cli/web/index.html")
var templates = template.Must(template.ParseFiles(f, f2))

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

	p, _ := utils.GetFromRoot("cli/web/assets")
	imgs, _ := utils.GetFromRoot("articles/imgs")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(p))))
	http.Handle("/articles/imgs/", http.StripPrefix("/articles/imgs/", http.FileServer(http.Dir(imgs))))
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
