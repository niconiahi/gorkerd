package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/niconiahi/gorkerd/src/routes/home"
)

func main() {
	mux := http.NewServeMux()
	h := home.Handler{}

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			print("incoming request from GET /\n")
			t := template.Must(h.CreateRoot())
			t.Execute(w, h.RunLoader())
		}
	})

	mux.HandleFunc("GET /home/route", func(w http.ResponseWriter, r *http.Request) {
		file, err := home.Javascript.ReadFile("route.js")
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(file)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
