package routes

import (
	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/crud-api-with-mysql/controllers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{isbn}", controllers.GetBookByIsbn).Methods("GET")
	router.HandleFunc("/books/create", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/update/{isbn}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/delete/{isbn}", controllers.DeleteBook).Methods("DELETE")

}
