package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/niconiahi/gorkerd/src/routes/home"
)

//go:embed src/routes/home/route.js src/routes/home/root.js
var js embed.FS

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		print("incoming request from GET /\n")
		h := home.Handler{}
		t := template.Must(h.GetHtml())
		// h.GetJs()
		t.Execute(w, h.GetData())
	})

	mux.HandleFunc("GET /route", func(w http.ResponseWriter, r *http.Request) {
		jsFile, err := js.ReadFile("src/routes/home/route.js")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(jsFile)

	})
	// mux.Handle("GET /js", http.FileServer(http.FS(js)))

	// jsFile, err := js.ReadFile("src/routes/home/route.js")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
