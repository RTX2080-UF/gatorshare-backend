package main

import (
	"gatorshare/models"
	"gatorshare/routes"
)


func main() {
  models.Init()
  DB := models.GetDB()
  router := routes.InitializeRoutes(DB)
  router.Run(":8080")
}