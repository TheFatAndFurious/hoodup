package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"goserver.com/handlers"
	"goserver.com/middlewares"
)



func NewRouter() *mux.Router {

	r := mux.NewRouter() 
	

	r.Handle("/bro", middlewares.TimingMiddleware(middlewares.CompressResponse(http.HandlerFunc(handlers.HomeHandler)))).Methods("GET")
	r.Handle("/gzip", middlewares.TimingMiddleware(http.HandlerFunc(handlers.HomeHandler))).Methods("GET")
	return r
}