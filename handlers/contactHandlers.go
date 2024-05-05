package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)


	func ContactHandler(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "contact.html")))            
            data := struct {
                Title string
            } {
                Title: "Contact",
            }
            err := tmpl.ExecuteTemplate(w, "base.html", data)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        }
	}

	
	
	