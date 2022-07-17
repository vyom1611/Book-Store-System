package routes

import (
	"github.com/gorilla/mux"
	"main.go/pkg/controllers"
)

//Routes
var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Defining the http methods for each route
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}