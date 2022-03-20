package test

import (
	"fmt"
	"gatorshare/models"
	"math/big"
	"crypto/rand"
	"testing"
)

func TestRegisterNewUserUnique(t *testing.T){
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
func createUser(t *testing.T)(users *models.User, username string){
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
	  return user, user.Username
  	}
	return
}
func TestGetUserProfile(t *testing.T){
	user, _:= createUser(t)
	if(user.ID != 0){
		res := models.GetUserProfile(testobj.DB,user,int(user.ID))
		if(res == nil){
			t.Log("Succesfully able to return user")
		}else {
			t.Error("Unable to return user!")
		}
	}
}


func TestGetUserDetailByUsername(t *testing.T){
	user, username:= createUser(t)
	if(username != "" && user.ID != 0){
		res, _ := models.GetUserDetailByUsername(testobj.DB, string(username))
		if(res.ID != 0){
			t.Log("Succesfully able to return user")
		}else {
			t.Error("Unable to return user!")
		}
	}
}

func TestDeleteUser(t *testing.T){
	user, _ := createUser(t)
	if(user.ID != 0){
		res := models.DeleteUser(testobj.DB,int(user.ID))
		if res == nil {
			t.Log("Succesfully able to delete user")
		}else {
			t.Error("Unable to delete user!")
		}
	}
}

