package cyoa

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var htmlTemplate = &template.Template{}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	htmlTemplate = template.Must(template.ParseFiles(wd + "/pkg/templates/story.html"))
}

func ParseStoryJson(file *os.File) (IStory, error) {
	d := json.NewDecoder(file)
	var story IStory
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s IStory) http.Handler {
	return handler{s}
}

type handler struct {
	s IStory
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := htmlTemplate.Execute(w, chapter)
		if err != nil {
			log.Fatalf("%v", err)
			http.Error(w, "Something went wrong", http.StatusConflict)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}
