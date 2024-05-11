package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"goserver.com/data"
	"goserver.com/routes"
)

// func initEnv(){
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading.env file")
// 	}
// }

func NewRouter() *mux.Router {
	r := mux.NewRouter() 

	return r
}


func main() {


	// initEnv()

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

	fs := http.FileServer(http.Dir("/var/www/static/"))
	pub := http.FileServer(http.Dir("../public/"))
	
	myApp.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	myApp.Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", pub))

	// data.CronTest()

	//go:generate npm run build

	http.ListenAndServe(":8080", myApp.Router)
}
