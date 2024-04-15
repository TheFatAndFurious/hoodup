package app

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/handlers"
	"goserver.com/middlewares"
)

// This is where we define the struct of App, we'll combine the Router and DB as to make the database accessible to all routes.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Here we will list all the routes available, we might need to break it up a little bit later if there are too many routes.
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/users", handlers.CreateUserHandler(a.DB)).Methods("POST")
	// a.Router.HandleFunc("/users", handlers.GetAllUsersHandler(a.DB)).Methods("GET")
	a.Router.HandleFunc("/users/{id}", handlers.UpdateUser(a.DB)).Methods("POST")
	a.Router.Handle("/commander", middlewares.ProtectRoute(handlers.AdminHandler(a.DB))).Methods("GET")
	a.Router.HandleFunc("/login", handlers.LoginHandler(a.DB)).Methods("GET")
	a.Router.HandleFunc("/login", handlers.LoginVerificator(a.DB)).Methods("POST")
	a.Router.HandleFunc("/users/{id}", handlers.UpdateUserHandler(a.DB)).Methods("GET")
}

