package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/handlers"
)


func ArticlesRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/articles", handlers.CreateArticle(db)).Methods("POST")
	router.HandleFunc("/articles", handlers.DisplayAllArticlesHandler(db)).Methods("GET")
	router.HandleFunc("/article/{id}", handlers.DisplayArticleHandler(db)).Methods("GET")
}