package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/internal/handlers"
)
func LoginRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/login", handlers.LoginHandler(db)).Methods("GET")
    router.HandleFunc("/login", handlers.LoginVerificatorHandler(db)).Methods("POST")
}