package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"goserver.com/handlers"
)

func ServerSentEventsRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/sse", handlers.EventsHandler(db)).Methods("GET")
}