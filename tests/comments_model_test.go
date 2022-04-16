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
		Username: "TestUser99ab" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email: "TestUser99ab" + fmt.Sprint(rnum) + "@gatorshare.com",
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

	models.AddNewPost(testobj.DB, post)
	comment := &models.Comment {
		PostID: post.ID,
		UserID: user.ID,
		Message: "Test comment" + fmt.Sprint(rnum),
  	} 

  res, _ := models.AddNewComment(testobj.DB, comment)
  if res == 0 {
	  t.Error("Unable to create post!")
	} else {
	  t.Log("Succesfully able to create post")	  
  }
 
//   testobj.DB.Delete(comment)
//   testobj.DB.Delete(post)
//   testobj.DB.Delete(user)
  print(res)
}

func createComment(t *testing.T)(comment *models.Comment){
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
	user := &models.User{
		Username: "TestUser123" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email: "TestUser123" + fmt.Sprint(rnum)+ "@gatorshare.com",
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
		res, _ := models.AddNewPost(testobj.DB, post)
		if res != 0 {
			comment := &models.Comment{
				UserID: user.ID,
				PostID: post.ID,
				Message: "Test Commment"+fmt.Sprint(rnum),
			}
			res, _ := models.AddNewComment(testobj.DB, comment)
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
		res := models.GetOneComment(testobj.DB, comment, uint(comment.ID))
		print(res)
		if  comment.ID != 0 && res == nil{
			t.Log("Succesfully able to return comment")	
		} else {
			t.Error("Unable to return comment!")
		}
	}
}
func createcomments(t *testing.T)(comment *[]models.Comment, postID uint){
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
	user := &models.User{
		Username: "TestUserxyz" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email: "TestUserxyz" + fmt.Sprint(rnum)+ "@gatorshare.com",
		Lastname: "1",
		Password: "Test",
	} 

	res, _ := models.AddNewUser(testobj.DB, user)

	if res != 0 {
		post := &models.Post{
			Title: "Test new post",
			UserID: user.ID,
			Description: "Test Message",
			UserLimit: 4,
			Status: 2,
		}

		res, _ := models.AddNewPost(testobj.DB, post)

		if res == 0  {
			t.Error("Unable to create post!")
		} else {
			comment := &models.Comment{
				UserID: user.ID,
				PostID: post.ID,
				Message: "Test Commment"+fmt.Sprint(rnum),
			}
			comment1 := &models.Comment{
				UserID: user.ID,
				PostID: post.ID,
				Message: "Test Commment1"+fmt.Sprint(rnum),
			}
			comment2 := &models.Comment{
				UserID: user.ID,
				PostID: post.ID,
				Message: "Test Commment2"+fmt.Sprint(rnum),
			}
			res, _ := models.AddNewComment(testobj.DB, comment)
			res1, _ := models.AddNewComment(testobj.DB, comment1)
			res2, _ := models.AddNewComment(testobj.DB, comment2)
			if res != 0 && res1 !=0 && res2!=0 {
				comments := &[]models.Comment{*comment,*comment1,*comment2}	
				return comments,post.ID
			}else {
				t.Error("Unable to create comments")
			}	
		}

	}else{
		t.Error("Cannot return post User not created!")
	}
	return 
}
func TestGetAllcomment(t *testing.T){
	comments,ID := createcomments(t)
	if((*comments)[0].PostID != ID){
		res := models.GetAllComment(testobj.DB,comments,uint(ID))
		if res == nil {
			t.Log("Succesfully able to return comments")
		}else {
			t.Error("Unable to return comments!")
		}
	}
}
func TestDeletecomment(t *testing.T){
	comment := createComment(t)
	if comment.ID != 0 && comment.PostID != 0 && comment.UserID != 0{
		res := models.DeleteComment(testobj.DB, comment, uint(comment.ID))
		if res == nil {
			t.Log("Succesfully able to delete comment")
		}else {
			t.Error("Unable to delete comment!")
		}
	}
}


func TestUpdatecomment(t *testing.T){
	comment := createComment(t)
	if(comment.ID != 0){
		res := models.UpdateComment(testobj.DB, comment)
		if res == nil {
			t.Log("Succesfully able to update comment")
		}else {
			t.Error("Unable to update comment!")
		}
	}
//   testobj.DB.Delete(comment)
}