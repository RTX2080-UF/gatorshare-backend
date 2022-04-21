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

func BootstrapTags(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))

	userBaseObj = models.User{
		Username:  "Test_User_Tags" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email:     "TestUserTags" + fmt.Sprint(rnum) + "@gatorshare.com",
		Lastname:  "Tags",
		Password:  "Test",
	}

	_, err := models.AddNewUser(testobj.DB, &userBaseObj)

	postBaseObj = models.Post{
		Title:       "Test post",
		UserID:      userBaseObj.ID,
		Description: "Test Message",
		UserLimit:   4,
		Status:      2,
	}

	_, err = models.AddNewPost(testobj.DB, &postBaseObj)
	commentBaseObj = models.Comment{
		PostID:  postBaseObj.ID,
		UserID:  userBaseObj.ID,
		Message: "Test comment" + fmt.Sprint(rnum),
	}

	_, err = models.AddNewComment(testobj.DB, &commentBaseObj)

	if err != nil || userBaseObj.ID == 0 || postBaseObj.ID == 0 || commentBaseObj.ID == 0 {
		t.Error("Unable to create base Object aborting!")
	} else {
		t.Log("Succesfully able to create base object")
	}

	//   testobj.DB.Delete(comment)
	//   testobj.DB.Delete(post)
	//   testobj.DB.Delete(user)
}

func CreateTag(t *testing.T) {
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

func GetOneTag(t *testing.T) {
	var retrievedTag models.Tag
	res := models.GetTag(testobj.DB, &retrievedTag, tagBaseObj.ID)


	if retrievedTag.ID == 0 || res != nil  {
		t.Error("Unable to get tag!")
		return
	}
	
	t.Log("Succesfully retrieved tag")
}

func CreateTags(t *testing.T) {
	numTags := 10
	
	for i:=0; i<numTags; i++ {
		rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
		tagObjArr = append(tagObjArr, models.Tag{ 
			Name: "Test Tag Bulk" + fmt.Sprint(i) + fmt.Sprint(rnum),
			Description: "Bulk Tags" + fmt.Sprint(i) + "Description" + fmt.Sprint(rnum),
			Votes: numTags - i,
			CreatorId: userBaseObj.ID,
		})
	}

	// t.Logf("----%v----", tagObjArr)
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

func GetPopularTag(t *testing.T) {
	numTags := 10

	var popularTagsArr []models.Tag
	res := models.GetPopularTags(testobj.DB, &popularTagsArr, numTags)
	if res != nil || len(popularTagsArr) != numTags {
		t.Error("Failed to get popular tags!")
	}
		
	t.Log("Succesfully retrieved popular tags")
}

func UpdateTag(t *testing.T) {
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

func DeleteTag(t *testing.T) {
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

func TestAllTags(t *testing.T){
    t.Run("Bootstrap", BootstrapTags)
    t.Run("CreateOne", CreateTag)
    t.Run("GetOne", GetOneTag)
    t.Run("CreateBulk", CreateTags)
    t.Run("GetBulk", GetPopularTag)
    t.Run("UpdateTag", UpdateTag)
    t.Run("DeleteTag", DeleteTag)
}