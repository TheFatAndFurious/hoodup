package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/handlers"
	"goserver.com/middlewares"
)

func AdminRoutes(router *mux.Router, db *sql.DB) {
	router.Handle("/commander", middlewares.ProtectRoute(handlers.AdminHandler(db))).Methods("GET")
	router.HandleFunc("/commander/editor/", handlers.EditorHandler(db)).Methods("GET")
	router.HandleFunc("/commander/editor/{id}", handlers.EditorHandler(db)).Methods("GET")
	router.HandleFunc("/commander/articles", handlers.ArticlesPanelHandler(db)).Methods("GET")
}