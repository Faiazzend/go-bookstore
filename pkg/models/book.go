package models

import (
	"gorm.io/gorm"
	"github.com/Faiazzend/go-bookstore/pkg/config"

)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func Init() {
	config.InitDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book