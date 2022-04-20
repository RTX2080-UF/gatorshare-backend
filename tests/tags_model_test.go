package test

import (
	"crypto/rand"
	"fmt"
	"gatorshare/models"
	"math/big"
	mrand "math/rand"
	"testing"
	"time"
)

var userBaseObj models.User
var postBaseObj models.Post
var commentBaseObj	models.Comment
var tagBaseObj	models.Tag
var tagObjArr []models.Tag

func TestBootstrapTags(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))

	userBaseObj := models.User{
		Username:  "TestUser99ab" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email:     "TestUser99ab" + fmt.Sprint(rnum) + "@gatorshare.com",
		Lastname:  "1",
		Password:  "Test",
	}

	_, err := models.AddNewUser(testobj.DB, &userBaseObj)

	postBaseObj := models.Post{
		Title:       "Test post 1",
		UserID:      userBaseObj.ID,
		Description: "Test Message",
		UserLimit:   4,
		Status:      2,
	}

	_, err = models.AddNewPost(testobj.DB, &postBaseObj)
	commentBaseObj := models.Comment{
		PostID:  postBaseObj.ID,
		UserID:  userBaseObj.ID,
		Message: "Test comment" + fmt.Sprint(rnum),
	}

	_, err = models.AddNewComment(testobj.DB, &commentBaseObj)

	if err != nil || userBaseObj.ID == 0 || postBaseObj.ID == 0 || commentBaseObj.ID == 0 {
		t.Error("Unable to create base Object aborting!")
	} else {
		t.Log("Succesfully able to create post")
	}

	//   testobj.DB.Delete(comment)
	//   testobj.DB.Delete(post)
	//   testobj.DB.Delete(user)
}

func TestCreateTag(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
	if (userBaseObj.ID == 0) {
		t.Error("Test failed no associated user object exist")
	}

	tagBaseObj = models.Tag {
		CreatorId:  userBaseObj.ID,
		Description: "Tag created for testing",
		Name: "Test Tag" + fmt.Sprint(rnum),
	}

	res, err := models.AddNewTag(testobj.DB, &tagBaseObj)
	if tagBaseObj.ID == 0 || res == 0 || err != nil {
		t.Error("Unable to create tag")
		return
	}

	t.Log("Succesfully created tag")
}

func TestGetOneTag(t *testing.T) {
	var retrievedTag models.Tag
	res := models.GetTag(testobj.DB, &retrievedTag, tagBaseObj.ID)

	if retrievedTag.ID == 0 || res != nil || retrievedTag != tagBaseObj {
		t.Error("Unable to get tag!")
		return
	}
	
	t.Log("Succesfully retrieved tag")
}

func TestCreateTags(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
	numTags := 10

	for i:=0; i<numTags; i++ {
		tagObjArr = append(tagObjArr, models.Tag{ 
			Name: "Test Tag Bulk" + fmt.Sprint(i),
			Description: "Bulk Tags" + fmt.Sprint(i) + "Description" + fmt.Sprint(rnum),
			Votes: numTags - i,
			CreatorId: userBaseObj.ID,
		})
	}

	tagIds, err := models.InsertTags(testobj.DB, tagObjArr)
	if len(tagIds) != numTags || err != nil {
		t.Error("Unable to create Tags")
		return
	}

	for i:=0; i<numTags; i++ { 
		if tagIds[i] <= 0 {
			t.Error("Invalid tag Id generated")
			return
		}
	}

	t.Log("Succesfully created bulk tags")
}

func TestGetPopularTag(t *testing.T) {
	numTags := 10

	var popularTagsArr []models.Tag
	res := models.GetPopularTags(testobj.DB, &popularTagsArr, numTags)
	if res != nil || len(popularTagsArr) != numTags {
		t.Error("Failed to get popular tags!")
	}
		
	t.Log("Succesfully retrieved popular tags")
}

func TestUpdateTag(t *testing.T) {
	tagObjCpy := tagBaseObj
    mrand.Seed(time.Now().UnixNano())

	tagObjCpy.Name = "Tag name modified"
	tagObjCpy.Votes = mrand.Intn(1000)
	tagObjCpy.Description = "Updated test obj"

	res := models.UpdateTag(testobj.DB, &tagObjCpy)
	if res != nil || tagObjCpy == tagBaseObj {
		t.Error("Unable to update tag!")
		return
	}

	t.Log("Succesfully able to update tag")
}

func TestDeleteTag(t *testing.T) {
	res := models.DeleteTag(testobj.DB, tagBaseObj.ID)
	if res != nil {
		t.Error("Unable to delete tag!")
		return
	}
	
	var retrievedTag models.Tag
	res = models.GetTag(testobj.DB, &retrievedTag, tagBaseObj.ID)
	if res == nil {
		t.Error("Unable to delete tag!")
		return
	}

	t.Log("Succesfully able to delete comment")
}