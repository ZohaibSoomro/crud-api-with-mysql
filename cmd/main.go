package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/crud-api-with-mysql/config"
	"github.com/zohaibsoomro/crud-api-with-mysql/models"
	"github.com/zohaibsoomro/crud-api-with-mysql/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	config.Connect()
	config.GetDb().AutoMigrate(&models.Book{})
	// http.Handle("/", router)
	fmt.Println("Server started...")
	panic(http.ListenAndServe(":8080", router))
}
