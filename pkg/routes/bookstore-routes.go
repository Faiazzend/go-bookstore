package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Faiazzend/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(c *gin.Engine) {
	c.POST("/book", controllers.CreateBook)
	c.GET("/book", controllers.GetBook)
	c.GET("/book/:id", controllers.GetBookById)
	c.PUT("/book/:id", controllers.UpdateBook)
	c.DELETE("/book/:id", controllers.DeleteBook)


}

