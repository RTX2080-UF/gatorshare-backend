package test

import (
	"fmt"
	"gatorshare/models"
	"math/big"
	"crypto/rand"
	"testing"
)

func TestCreateNewComment(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))

	user := &models.User{
		Username: "TestUser1" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email: "TestUser1" + fmt.Sprint(rnum) + "@gatorshare.com",
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

	models.AddNewpost(testobj.DB, post)
	comment := &models.Comment {
		PostID: post.ID,
		UserID: user.ID,
		Message: "Test comment" + fmt.Sprint(rnum),
  	} 

  res, _ := models.AddNewcomment(testobj.DB, comment)
  if res == 0 {
	  t.Error("Unable to create post!")
	} else {
	  t.Log("Succesfully able to create post")	  
  }
 
  testobj.DB.Delete(user)
  testobj.DB.Delete(post)
  testobj.DB.Delete(comment)
  print(res)
}

func createComment(t *testing.T)(comment *models.Comment){
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
	user := &models.User{
		Username: "TestUser123" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email: "TestUser17" + fmt.Sprint(rnum)+ "@gatorshare.com",
		Lastname: "1",
		Password: "Test",
	} 

	res, _ := models.AddNewUser(testobj.DB, user)

	if res != 0 {
		// print(res)
		post := &models.Post{
			Title: "Test new post",
			UserID: user.ID,
			Description: "Test Message"+fmt.Sprint(rnum),
			UserLimit: 4,
			Status: 2,
		}
		res, _ := models.AddNewpost(testobj.DB, post)
		if res != 0 {
			comment := &models.Comment{
				UserID: user.ID,
				PostID: post.ID,
				Message: "Test Commment"+fmt.Sprint(rnum),
			}
			res, _ := models.AddNewcomment(testobj.DB, comment)
			if res != 0 {
				return comment
			}else {
				t.Error("Unable to create comment")
			}


		} else {
			t.Error("Unable to create post!")
		}

	}else{
		t.Error("Cannot return post User not created!")
	}
	return 
}

func TestGetOnecomment(t *testing.T){
	comment := createComment(t)
	if comment.ID != 0{
		res := models.GetOnecomment(testobj.DB, comment, int(comment.ID))
		if res != nil {
			t.Log("Succesfully able to return comment")	
		} else {
			t.Error("Unable to return comment!")
		}
	}
}