package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"goserver.com/data"
	"goserver.com/internal/handlers"
	"goserver.com/internal/routes"
	"goserver.com/internal/utils"
)



func NewRouter() *mux.Router {
	r := mux.NewRouter() 

	return r
}


func main() {


	utils.InitEnv()

	handlers.InitializeHandlers()

	db, err := data.DbInitialize()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := data.InitializeSchema(db); err!= nil {
		log.Fatalf("failed to initialize schema: %v", err)
    }

	myApp := &routes.App{
		Router: NewRouter(),
        DB:     db,
    }

	
	myApp.InitializeRoutes()

	fs := http.FileServer(http.Dir("/app/static/"))
	pub := http.FileServer(http.Dir("app/public/"))
	
	myApp.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	myApp.Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", pub))

	// data.CronTest()

	//go:generate npm run build

	http.ListenAndServe(":8080", myApp.Router)
}
