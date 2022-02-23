package main

import (
	helper "gatorshare/middleware"
	"gatorshare/models"
	"gatorshare/routes"
)


func main() {
  envsrc := helper.LoadEnv(".env")
  models.Init(envsrc)
  DB := models.GetDB()
  router := routes.InitializeRoutes(DB)
  hostport := helper.GetEnv("PORT", "8080", envsrc)
  router.Run(":"+hostport)
}