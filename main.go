package main

import (
	"gatorshare/middleware"
	"gatorshare/models"
	"gatorshare/routes"
)


func main() {
  envsrc := middleware.LoadEnv(".env")
  models.Init(envsrc)
  DB := models.GetDB()
  router := routes.InitializeRoutes(DB)
  hostport := middleware.GetEnv("PORT", "8080", envsrc)
  router.Run(":"+hostport)
}