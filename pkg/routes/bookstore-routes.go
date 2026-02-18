package routes

import (
	"github.com/Faiazzend/go-bookstore/pkg/auth"
	"github.com/Faiazzend/go-bookstore/pkg/controllers"
	"github.com/Faiazzend/go-bookstore/pkg/middleware"
	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(c *gin.Engine) {
	c.POST("/login", controllers.Login)

	authorized := c.Group("/")
	authorized.Use(middleware.JWTAuth(auth.JWTSecret()))
	{
		authorized.POST("/book", controllers.CreateBook)
		authorized.GET("/book", controllers.GetAllBook)
		authorized.GET("/book/:id", controllers.GetBookByID)
		authorized.PUT("/book/:id", controllers.UpdateBook)
		authorized.DELETE("/book/:id", controllers.DeleteBook)
	}
}
