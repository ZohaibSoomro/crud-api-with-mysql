package models

import (
	"github.com/zohaibsoomro/crud-api-with-mysql/config"
	"gorm.io/gorm"
)

type Book struct {
	*gorm.Model `json:"-"`
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
}

func (b *Book) CreateBookInDb() *Book {
	db := config.GetDb()
	db.Create(b)
	return b
}

func GetAllBooksFromDb() []Book {
	var books []Book
	db := config.GetDb()
	db.Select([]string{"ID", "isbn", "title", "author"}).Find(&books)
	return books
}

func GetBookbyIsbnFromDb(isbn string) Book {
	var book Book
	db := config.GetDb()
	// db.Where("isbn = ?", isbn).Find(&book)

	db.Select([]string{"ID", "isbn", "title", "author"}).First(&book, "isbn=?", isbn)
	return book
}

func UpdateBookbyIsbnInDb(isbn string, b Book) (Book, *gorm.DB) {
	var book Book
	db := config.GetDb()
	DeleteBookbyIsbnFromDb(isbn)
	b.CreateBookInDb()
	b.ISBN = isbn
	db.Save(&b)
	// db.Model(b).Where("isbn=?", isbn).Update("title","koi mil gya")
	// db.Where("isbn = ?", isbn).Set("title")
	return book, db
}
func DeleteBookbyIsbnFromDb(isbn string) Book {
	var book Book
	db := config.GetDb()
	db.Where("isbn = ?", isbn).Delete(&book)
	return book
}
