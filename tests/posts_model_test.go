package test

import (
	"crypto/rand"
	"fmt"
	"gatorshare/models"
	"math/big"
	"testing"
)

func TestAddNewpost(t *testing.T) {
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

	res, _ := models.AddNewpost(testobj.DB, post)
	if res == 0 {
		t.Error("Unable to create post!")
  	} else {
		t.Log("Succesfully able to create post")	  
	}

	// testobj.DB.Delete(user)
	// testobj.DB.Delete(post)
	print(res)  

}
func createPost(t *testing.T)(post *models.Post){
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
	user := &models.User{
		Username: "TestUser17" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email: "TestUser17" + fmt.Sprint(rnum)+ "@gatorshare.com",
		Lastname: "1",
		Password: "Test",
	} 

	res, _ := models.AddNewUser(testobj.DB, user)

	if res != 0 {
		print(res)
		post := &models.Post{
			Title: "Test new post",
			UserID: user.ID,
			Description: "Test Message",
			UserLimit: 4,
			Status: 2,
		}
		res, _ := models.AddNewpost(testobj.DB, post)
		if res == 0 {
			t.Error("Unable to create post!")
		} else {
			return post
		}

	}else{
		t.Error("Cannot return post User not created!")
	}
	return 
}

func TestGetOnepost(t *testing.T){

	post := createPost(t)
	if(post.ID  != 0){
		res:= models.GetOnepost(testobj.DB, post, int(post.ID))
		print(res)
		if post.ID != 0 && res == nil {
			t.Log("Succesfully able to return post")
		}else{
			t.Error("Unable to return post!")
		}
	}
}
func createPosts(t *testing.T)(post *[]models.Post, userID uint){
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
	user := &models.User{
		Username: "TestUser17" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email: "TestUser17" + fmt.Sprint(rnum)+ "@gatorshare.com",
		Lastname: "1",
		Password: "Test",
	} 

	res, _ := models.AddNewUser(testobj.DB, user)

	if res != 0 {
		print(res)
		post := &models.Post{
			Title: "Test new post",
			UserID: user.ID,
			Description: "Test Message",
			UserLimit: 4,
			Status: 2,
		}
		post1 := &models.Post{
			Title: "Test new post",
			UserID: user.ID,
			Description: "Test Message 1",
			UserLimit: 4,
			Status: 2,	
		}
		post2 := &models.Post{
			Title: "Test new post 1",
			UserID: user.ID,
			Description: "Test Message 2",
			UserLimit: 4,
			Status: 2,	
		}

		res, _ := models.AddNewpost(testobj.DB, post)
		res1, _ := models.AddNewpost(testobj.DB, post1)
		res2, _ := models.AddNewpost(testobj.DB, post2)


		if res == 0 && res1 == 0 && res2 == 0 {
			t.Error("Unable to create post!")
		} else {
		posts := &[]models.Post{*post,*post1,*post2}	
		return posts,user.ID
		}

	}else{
		t.Error("Cannot return post User not created!")
	}
	return 
}
func TestGetAllpost(t *testing.T){
	
	post, ID := createPosts(t)

	if((*post)[0].UserID == ID ){
		res := models.GetAllpost(testobj.DB, post, int(ID))
		if res == nil {
			t.Log("Succesfully able to return posts")
		}else {
			t.Error("Unable to return posts!")
		}
	}
}
func TestDeletepost(t *testing.T){
	post := createPost(t)
	if(post.ID != 0){
		res := models.Deletepost(testobj.DB,post,int(post.ID))
		if res == nil {
			t.Log("Succesfully able to delete post")
		}else {
			t.Error("Unable to delete post!")
		}
	}
}