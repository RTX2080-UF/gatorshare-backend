package test

import (
	"fmt"
	"gatorshare/models"
	"math/rand"
	"testing"
)

func TestRegisterNewUserUnique(t *testing.T) {
	user := &models.User{
	  Username: "TestUser1" + fmt.Sprint(rand.Intn(10000)),
	  Firstname: "Test User",
	  Email: "TestUser1" + fmt.Sprint(rand.Intn(10000)) + "@gatorshare.com",
	  Lastname: "1",
	  Password: "Test",
  } 

  res, _ := models.AddNewUser(testobj.DB, user)
  if res == 0 {
	  t.Error("Unable to create post!")
	} else {
	  t.Log("Succesfully able to create post")	  
  }

  testobj.DB.Delete(user)
  print(res)
}
