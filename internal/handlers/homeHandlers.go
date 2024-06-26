package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

// var projectRootPath = "/var/www/"



	func HomepageHandler(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "homepage.html")))            
            data := struct {
                Title string
            } {
                Title: "Homepage",
            }
            err := tmpl.ExecuteTemplate(w, "base.html", data)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        }
	}

	
	
	