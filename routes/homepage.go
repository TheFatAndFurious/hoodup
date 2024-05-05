package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/handlers"
)

func HomepageRoutes (r *mux.Router, db *sql.DB) {
	r.HandleFunc("/", handlers.HomepageHandler(db)).Methods("GET")
}