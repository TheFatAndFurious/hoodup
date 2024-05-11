package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func ErrorHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	errorCode := r.URL.Query().Get("code")
	tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "error.html")))
		var errorMessage string
	switch errorCode {
		case "404": 
		errorMessage = "Page inexistante"
		default:
		errorMessage = "Erreur inconnue"
	}
	data := struct {
	Title string
	Error string
} {
	Title: "404 my guy",
	Error: errorMessage,
}
err := tmpl.ExecuteTemplate(w, "base.html", data)
if err!= nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}}
}