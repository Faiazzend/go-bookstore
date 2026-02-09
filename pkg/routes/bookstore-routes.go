package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Faiazzend/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(c *gin.Engine) {
	c.POST("/book", controllers.CreateBook)
	c.GET("/book", controllers.GetAllBook)
	c.GET("/book/:id", controllers.GetBookByID)
	c.PUT("/book/:id", controllers.UpdateBook)
	c.DELETE("/book/:id", controllers.DeleteBook)


}

