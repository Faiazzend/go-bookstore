package controllers

import (
	
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/Faiazzend/go-bookstore/pkg/models"

)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}
	createdBook := book.CreateBook()
	c.JSON(http.StatusCreated, createdBook)
}

func GetAllBook(c *gin.Context) {
	books := models.GetAllBook()
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book, result := models.GetBookById(int64(bookID))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book not found",
		})
		return
	}
	c.JSON(http.StatusOK, book)

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := models.DeleteBook(bookID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"message": "Book deleted successfully",

	})

}

func UpdateBook(c *gin.Context) {
	var updateBook models.Book

	id := c.Param("id")
	bookID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookDetails, db := models.GetBookById(bookID)

    if db.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if bookDetails.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	c.JSON(http.StatusOK, bookDetails)

	

}
