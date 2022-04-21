package test

import (
	"crypto/rand"
	"fmt"
	"gatorshare/models"
	// "log"
	"math/big"
	"testing"
)


func BootstrapNotification(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))

	userBaseObj = models.User{
		Username:  "Test_User_Notification" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email:     "TestUserNotifications" + fmt.Sprint(rnum) + "@gatorshare.com",
		Lastname:  "Notifications",
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

func CreateNotification(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))
	if (userBaseObj.ID == 0) {
		t.Error("Test failed no associated user object exist")
	}

	notificationBaseObj = models.Notification {
		UserID:  userBaseObj.ID,
		Description: "Notification created for testing" + fmt.Sprint(rnum),
		ReadStatus: false,
	}

	_, err := models.AddNotification(testobj.DB, &notificationBaseObj)
	if  notificationBaseObj.ID == 0 || err != nil {
		t.Error("Unable to create Notification")
		return
	}

	t.Log("Succesfully created Notification")
}

func GetOneNotifications(t *testing.T) {
	var retrievedNotification []models.Notification
	res := models.GetNewNotifications(testobj.DB, &retrievedNotification, userBaseObj.ID)

	if len(retrievedNotification) == 0 || res != nil  {
		t.Error("Unable to get Notification!")
		return
	}
	
	t.Log("Succesfully retrieved Notification")
}

func UpdateNotifications(t *testing.T) {
	var retrievedNotification []models.Notification
	res := models.UpdateNotifications(testobj.DB, userBaseObj.ID)
	if res != nil  {
		t.Error("Failed to updateNotifications!")
	}

	res = models.GetNewNotifications(testobj.DB, &retrievedNotification, userBaseObj.ID)
	if len(retrievedNotification) != 0 || res != nil  {
		t.Error("Unable to update Notifications!")
		return
	}
		
	t.Log("Succesfully updated Notifications")
}


func TestAllNotification(t *testing.T){
    t.Run("Bootstrap", BootstrapTags)
    t.Run("Create Notification", CreateNotification)
    t.Run("Get New Notification", GetOneNotifications)
    t.Run("Update Notification", UpdateNotifications)
}