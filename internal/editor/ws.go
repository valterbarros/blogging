package editor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"valterbarros/blog-tools/internal/markdown"
	"valterbarros/blog-tools/internal/utils"

	"github.com/gorilla/websocket"
)

// DTO
type Action struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	File  string `json:"file"`
}

func parseAction(msg []byte) (*Action, error) {
	var a Action

	if err := json.Unmarshal(msg, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// DTO

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TODO: maybe move to markdown package
func processImageUrl(text string) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", nil
	}

	tpl := template.New("name")
	tpl.Funcs(template.FuncMap{
		"relative_url": func(path string) string {
			// ignores the relative_url pipeline only jekyll needs it
			return path
		},
	})
	_, errP := tpl.Parse(text)

	if errP != nil {
		return "", fmt.Errorf("Parse template: %w", errP)
	}

	b := new(bytes.Buffer)
	err := tpl.Execute(b, nil)

	if err != nil {
		return "", fmt.Errorf("Error executing: %w", err)
	}

	return b.String(), nil
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Erro ao fazer upgrade: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected via websocket!")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Erro to read message: %v\n", err)
			break
		}

		ac, err := parseAction(msg)

		switch ac.Type {
		case "change":
			text, err := processImageUrl(ac.Value)

			if err != nil {
				log.Printf("Error to write response: %v\n", err)
				continue
			}

			result := markdown.RemoveMetadata(text)
			result = markdown.MarkToHTML([]byte(result))
			// send markdown in html
			err = conn.WriteMessage(messageType, []byte(result))
		case "update":
			path, _ := utils.GetFromRoot("articles/_posts/" + ac.File)
			utils.SaveFile(path, ac.Value)
		}

		if err != nil {
			log.Printf("Error to write response: %v\n", err)
			break
		}
	}
}
