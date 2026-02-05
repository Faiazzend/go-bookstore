package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"github.com/Faiazzend/go-bookstore/pkg/routes"

)

func main() {
	router := gin.Default()
	routes.RegisterBookStoreRoutes(router)
	log.Fatal(router.Run(":8080"))

}