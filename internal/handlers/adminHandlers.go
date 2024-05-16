package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"goserver.com/internal/models"
)


func AdminHandler( db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(basePath)
	tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "adminContent.html")))
users, err := models.GetAllUsers(db)
if err!= nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
	data := struct {
	Title string
	Users []models.User
} {
	Title: "Admin Panel",
	Users: users,
}
err = tmpl.ExecuteTemplate(w, "base.html", data)
if err!= nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}}
}


func EditorHandler( db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		article := models.Article{}
		if vars["id"] != "" {
			id := vars["id"]
			articleId, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			article, err = models.FetchSingleArticle(db, articleId)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Redirect(w, r, "/error?code=404", http.StatusFound)
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		tmp := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "quill.html")))	
		authors, err := models.GetAllUsers(db)	
		if err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}	
		data := struct {
			Title string
			Authors []models.User
			Article models.Article
			Content template.HTML
		} {
			Title: "Editeur",
			Authors: authors,
			Article: article,
			Content: template.HTML(article.Content),
		}
		
		err = tmp.ExecuteTemplate(w, "base.html", data)
		if err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func ArticlesPanelHandler( db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "articlesPanel.html")))    
        articles, err := models.FetchAllArticles(db)    
        if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }    
        data := struct {
            Title string
            Articles []models.Article
        } {
            Title: "Articles",
            Articles: articles,
        }
        
        err = tmp.ExecuteTemplate(w, "base.html", data)
        if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
	}
}