package main

import (
	"log"

	"github.com/Faiazzend/go-bookstore/pkg/models"
	"github.com/Faiazzend/go-bookstore/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.Init()
	routes.RegisterBookStoreRoutes(router)
	log.Fatal(router.Run(":8080"))

}