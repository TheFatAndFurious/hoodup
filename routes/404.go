package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/handlers"
)


func ErrorRoutes(router *mux.Router, db *sql.DB) {
	router.Handle("/error", handlers.ErrorHandler()).Methods("GET")
}