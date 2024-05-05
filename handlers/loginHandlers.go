package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)


func LoginHandler( db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "login.html")))			
		data := struct {
			Title string
		} {
			Title: "Login",
		}
		err := tmp.ExecuteTemplate(w, "base.html", data)
		if err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}	