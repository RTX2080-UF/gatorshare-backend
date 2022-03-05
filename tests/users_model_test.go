package test

import (
	"fmt"
	"gatorshare/models"
	"math/big"
	"crypto/rand"
	"testing"
)

func TestRegisterNewUserUnique(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))

	user := &models.User{
	  Username: "TestUser1" + fmt.Sprint(rnum),
	  Firstname: "Test User",
	  Email: "TestUser1" + fmt.Sprint(rnum) + "@gatorshare.com",
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
