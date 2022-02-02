package main

import (
	"gatorshare/models"
	"gatorshare/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/", func(responseCtx *gin.Context) {
    responseCtx.JSON(http.StatusOK, gin.H{"data": "Welcome to Go and Gin!"})    
  })

  routes.InitializeRoutes(router)
  models.ConnectDatabase()
  router.Run(":8080")
}