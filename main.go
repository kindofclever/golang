package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
    tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil)
    })

    http.HandleFunc("/portfolio", func(w http.ResponseWriter, r *http.Request) {
        portfolioTmpl := template.Must(template.ParseFiles(filepath.Join("templates", "portfolio.html")))
        portfolioTmpl.Execute(w, nil)
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":80", nil)
}