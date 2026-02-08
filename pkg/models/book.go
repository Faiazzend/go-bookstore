package models

import (
	"github.com/Faiazzend/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Init() {
	config.InitDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books 

}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, result
}

func DeleteBook(Id int64) *gorm.DB {
	return db.Delete(&Book{}, Id)
	
}