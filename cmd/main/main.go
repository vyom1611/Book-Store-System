package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	//Setting a mux router equal to the defined routers from bookstore-routes
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)      		//Base path attached to the defiend router
	log.Fatal(http.ListenAndServe("localhost:9700", r))	//Starting server from mux router
}