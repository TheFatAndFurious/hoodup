package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/handlers"
)

func ContactRoutes (r *mux.Router, db *sql.DB) {
	r.HandleFunc("/contact", handlers.ContactHandler(db)).Methods("GET")
}