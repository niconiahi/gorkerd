package home

import (
	"embed"
	"html/template"
)

type Handler struct{}

type Loader struct {
	date string
}

func (h *Handler) RunLoader() Loader {
	return Loader{
		date: "blind",
	}
}

//go:embed route.js
var Javascript embed.FS

//go:embed root.html
var html embed.FS

func (h *Handler) CreateRoot() (*template.Template, error) {
	return template.ParseFS(html, "root.html")
}
