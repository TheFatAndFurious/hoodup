package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// This is where we define the struct of App, we'll combine the Router and DB as to make the database accessible to all routes.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) InitializeRoutes() {
	ArticlesRoutes(a.Router, a.DB)
	UsersRoutes(a.Router, a.DB)
	LoginRoutes(a.Router, a.DB)
	AdminRoutes(a.Router, a.DB)
	ErrorRoutes(a.Router, a.DB)
	HomepageRoutes(a.Router, a.DB)
	ContactRoutes(a.Router, a.DB)
	ServerSentEventsRoutes(a.Router, a.DB)
}