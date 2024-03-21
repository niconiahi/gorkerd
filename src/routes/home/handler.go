package home

import (
	"embed"
	"html/template"
	// "net/http"
)

type Event struct {
	Name string
}

type Loader struct {
	Event Event
}

type Handler struct{}

type Data struct {
	date string
}

func (h *Handler) GetData() Data {
	return Data{
		date: "blind",
	}
}

func (h *Handler) GetFiles() []string {
	return []string{
		"root.html",
		"route.js",
	}
}

//go:embed root.html
var html embed.FS

func (h *Handler) GetHtml() (*template.Template, error) {
	return template.ParseFS(html, "root.html")
}

//go:embed route.js
var js embed.FS

func (h *Handler) GetJs() {
	print("getting JS\n")
}
