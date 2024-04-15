package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"goserver.com/app"
	"goserver.com/data"
	router "goserver.com/internal"
)


func main() {

	db, err := data.DbInitialize()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := data.InitializeSchema(db); err!= nil {
		log.Fatalf("failed to initialize schema: %v", err)
    }

	myApp := &app.App{
		Router: router.NewRouter(),
        DB:     db,
    }

	myApp.InitializeRoutes()

	fs := http.FileServer(http.Dir("./static/"))
	myApp.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//go:generate npm run build

	http.ListenAndServe(":8080", myApp.Router)
}
