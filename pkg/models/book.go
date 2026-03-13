package models

import (
	"github.com/TechYash-bit/proj1Go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (Book) TableName() string {
	return "book_store"
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var Books Book
	db := db.Where("ID=?", id).Find(&Books)
	return &Books, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(&book)
	return book
}
