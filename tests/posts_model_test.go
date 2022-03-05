package test

import (
	"fmt"
	"gatorshare/models"
	"math/rand"
	"testing"
)

func TestAddNewpost(t *testing.T) {
  	user := &models.User{
		Username: "TestUser1" + fmt.Sprint(rand.Intn(10000)),
		Firstname: "Test User",
		Email: "TestUser1" + fmt.Sprint(rand.Intn(10000)) + "@gatorshare.com",
		Lastname: "1",
		Password: "Test",
	} 

	models.AddNewUser(testobj.DB, user)

	post := &models.Post{
		Title: "Test post 1",
		UserID: user.ID,
		Description: "Test Message",
		UserLimit: 4,
		Status: 2,
	}

	res, _ := models.AddNewpost(testobj.DB, post)
	if res == 0 {
		t.Error("Unable to create post!")
  	} else {
		t.Log("Succesfully able to create post")	  
	}

	testobj.DB.Delete(user)
	testobj.DB.Delete(post)
	print(res)  

}
