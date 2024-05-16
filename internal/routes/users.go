package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/internal/handlers"
)

func UsersRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/users", handlers.CreateUserHandler(db)).Methods("POST")
	// a.Router.HandleFunc("/users", handlers.GetAllUsersHandler(db)).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUserHandler(db)).Methods("GET")
}