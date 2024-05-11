package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"goserver.com/models"
)

//TODO: Add validation for all user inputs
func CreateArticle(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	article, err := models.FormToArticle(r.Form, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if article.Id == 0  {

		if err := models.PersistArticle(db, article); err!= nil {
		log.Fatalf("Failed to persist user: %v", err)
		return
		}
 
	} else {
		if err := models.UpdateArticle(db, article); err!= nil {
        log.Fatalf("Failed to persist user: %v", err)
		return
        }
    }
	http.Redirect(w, r, "/commander/articles", http.StatusSeeOther)
	}
}


func DisplayAllArticlesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "allArticles.html")))
		articles, err := models.FetchAllArticles(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := struct {
            Title string
            Articles []models.Article
		} {
			Title: "Articles",
            Articles: articles,
		}
		err = tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
}}

func DisplayArticleHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "singleArticle.html")))
		vars := mux.Vars(r)
		id := vars["id"]
		article := models.Article{}
		articleId, err := strconv.Atoi(id)
		if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
		article, err = models.FetchSingleArticle(db, articleId)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Redirect(w, r, "/error?code=404", http.StatusFound)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := struct {
			Title string
			Content template.HTML
            Article models.Article
        } {
			Title: article.Title,
			Content: template.HTML(article.Content),
            Article: article,
        }
		err = tmpl.ExecuteTemplate(w, "base.html", data)
		if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

		}
	}



