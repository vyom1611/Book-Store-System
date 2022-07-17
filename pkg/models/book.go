package models

import (
	"github.com/jinzhu/gorm"
	"main.go/pkg/config"
)

var db *gorm.DB

//Specifying the struct schema for book type
type Book struct {
	gorm.Model			
	Name string			`json:"name"`
	Author string		`json:"author"`
	Publication string	`json:"publication"`

}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//Setting up book instance creator function
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//Function for retrieving all Books instances from database
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}