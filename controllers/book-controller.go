package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/crud-api-with-mysql/models"
)

var GetAllBooks = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "request Method not supported", http.StatusNotAcceptable)
		return
	}
	books := models.GetAllBooksFromDb()
	b, err := json.MarshalIndent(books, "", "")
	if err != nil {
		http.Error(w, "some error occurred", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

var GetBookByIsbn = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "request Method not supported", http.StatusNotAcceptable)
		return
	}
	params := mux.Vars(r)
	book := models.GetBookbyIsbnFromDb(params["isbn"])

	b, err := json.MarshalIndent(book, "", "")
	if err != nil {
		http.Error(w, "some error occurred", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "request Method not supported", http.StatusNotAcceptable)
		return
	}
	var book models.Book
	bytes, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		http.Error(w, "Invalid json object", http.StatusNotAcceptable)
		return
	}
	book.CreateBookInDb()
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

var UpdateBook = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "request Method not supported", http.StatusNotAcceptable)
		return
	}
	params := mux.Vars(r)
	var book models.Book
	bytes, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		http.Error(w, "Invalid json object", http.StatusNotAcceptable)
		return
	}
	b, _ := models.UpdateBookbyIsbnInDb(params["isbn"], book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, _ = json.MarshalIndent(b, "", "")
	w.Write(bytes)
}

var DeleteBook = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "request Method not supported", http.StatusNotAcceptable)
		return
	}
	params := mux.Vars(r)
	models.DeleteBookbyIsbnFromDb(params["isbn"])
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	books, _ := json.MarshalIndent(models.GetAllBooksFromDb(), "", "")
	w.Write(books)
}
