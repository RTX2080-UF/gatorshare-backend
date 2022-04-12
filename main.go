package main

import (
	helper "gatorshare/middleware"
	"gatorshare/models"
	"gatorshare/routes"
)

var EnvSrc = helper.LoadEnv(".env")

func main() {
  models.Init(EnvSrc)
  DB := models.GetDB()
  router := routes.InitializeRoutes(DB, EnvSrc)
  hostport := helper.GetEnv("PORT", "8080", EnvSrc)
  router.Run(":"+hostport)
}