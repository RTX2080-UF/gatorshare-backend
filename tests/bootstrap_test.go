package test

import (
	"gatorshare/middleware"
	"gatorshare/models"
	"testing"

	"gorm.io/gorm"
)

type TestSuit struct {
	DB *gorm.DB
} 

var testobj = TestSuit{}

func TestMain (t *testing.T) {
	envsrc := middleware.LoadEnv(".env")
  	models.Init(envsrc)
  	testobj.DB = models.GetDB()
	// os.Exit(t.Run())
}

var userBaseObj models.User
var postBaseObj models.Post
var commentBaseObj	models.Comment
var tagBaseObj	models.Tag
var tagObjArr []models.Tag
var notificationBaseObj models.Notification