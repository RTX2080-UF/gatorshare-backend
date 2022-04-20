package controllers

import (
	"errors"
	"gatorshare/middleware"
	"gatorshare/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (base *Controller) GetProfileGeneric(ctx *gin.Context) {
	var userData models.User

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	userIdStr := ctx.Params.ByName("id")
	userId, err := strconv.Atoi(userIdStr)
	log.Print("Got request to get User profile", userId)

    if err != nil {
		errCustom := errors.New("invalid user id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	err = models.GetUserProfile(base.DB, &userData, uint(userId))
	userData.Password = ""
	userData.Bookmarks = ""
	userData.Email = ""

	if err != nil {
		errCustom := errors.New("unable to retrieve user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userData, err)
	}
}


func (base *Controller) GetProfile(ctx *gin.Context) {
	var userData models.User
	
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	log.Print("Got request to get User profile", uid)

	err := models.GetUserProfile(base.DB, &userData, uid)

	if err != nil {
		errCustom := errors.New("unable to retrieve user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		if userData.ID != uid {
			errCustom := errors.New("profile doesn't belong to the given user").Error()
			middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
			return
		}	
		middleware.RespondJSON(ctx, http.StatusOK, userData, err)
	}
}


func (base *Controller) UpdateProfile(ctx *gin.Context) {
	var newUserData UpdateUserProfile
	var currentUserData models.User

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	log.Print("Got request to Update User profile ", uid)

	err := models.GetUserProfile(base.DB, &currentUserData, uid)
	if err != nil {
		errCustom := errors.New("unable to retrieve user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return	
	}

	if currentUserData.ID != uid {
		errCustom := errors.New("profile doesn't belong to the given user")
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom.Error(), errCustom)
		return
	}
	
	err = ctx.ShouldBindJSON(&newUserData);
	if err != nil {
		errCustom := errors.New("invalid profile object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}
	
	// fmt.Printf("%+v\n",newUserData)
	if (newUserData.OldPassword != "" && !middleware.CheckPasswordHash(currentUserData.Password, newUserData.OldPassword)) {
		errCustom := errors.New("unable to update password, password didn't match")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
	}

	updatedUserData := models.User {
		Firstname: newUserData.Firstname,
		Lastname: newUserData.Lastname,
		Email: newUserData.Email,
		Zipcode: newUserData.Zipcode,
		Avatar: newUserData.Avatar,
	}
	updatedUserData.ID = currentUserData.ID

	if (newUserData.OldPassword != "") {
		updatedUserData.Password, err = middleware.HashPassword(newUserData.Password)
	} else {
		updatedUserData.Password = currentUserData.Password
	}

	err = models.UpdateUserProfile(base.DB, &updatedUserData)
	if err != nil {
		errCustom := errors.New("unable to update user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}
	
	
	notif_message := "User " + updatedUserData.Username + " profile has been updated"
	middleware.SendMail(
		"Updates", 
		updatedUserData.Firstname, 
		updatedUserData.Email, 
		"Your profile has been updated",
		notif_message,
		"")

	var notification = models.Notification {
		UserID: updatedUserData.ID,
		Description: notif_message,
	}

	_, err = models.AddNotification(base.DB, &notification)
	if (err != nil) {
		log.Printf("unable to add notification %v",err)
	}

	middleware.RespondJSON(ctx, http.StatusOK, updatedUserData, nil)
}


func (base *Controller) DeleteUser(ctx *gin.Context) {
	var userData models.User

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	log.Print("Got request to delete User profile", uid)

	err := models.GetUserProfile(base.DB, &userData, uid)
	if err != nil {
		errCustom := errors.New("unable to retrieve user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
	}
	ctx.BindJSON(&userData)

	if userData.ID != uid {
		errCustom := errors.New("profile doesn't belong to the given user").Error()
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
		return
	}

	err = models.DeleteUser(base.DB, int(uid))
	if err != nil {
		errCustom := errors.New("unable to delete user with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userData, nil)
	}
}


func (base *Controller) GetFollowers(ctx *gin.Context) {
	var userData models.User
	
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	log.Print("Got request to get User profile", uid)
	
	err := models.GetUserProfile(base.DB, &userData, uid)

	userIdStr := ctx.Params.ByName("userId")
	userId, err := middleware.ConvertStrToInt(userIdStr)
	if (err != nil) {
		middleware.RespondJSON(ctx, http.StatusBadGateway, err.Error(), err)
		return
	}
	
	var follower []models.Follower 
	err = models.GetFollowers(base.DB, &follower, userId)
	if (err != nil) {
		errCustom := errors.New("unable to get user followers").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	middleware.RespondJSON(ctx, http.StatusOK, follower, nil)
}


func (base *Controller) FollowUser(ctx *gin.Context) {
	var userData models.User
	
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	log.Print("Got request to get User profile", uid)
	
	err := models.GetUserProfile(base.DB, &userData, uid)

	if err != nil {
		errCustom := errors.New("unable to retrieve user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		if userData.ID != uid {
			errCustom := errors.New("profile doesn't belong to the given user").Error()
			middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
			return
		}
	}

	userIdStr := ctx.Params.ByName("userId")
	followerId, err := middleware.ConvertStrToInt(userIdStr)
	if (err != nil) {
		middleware.RespondJSON(ctx, http.StatusBadGateway, err.Error(), err)
		return
	}

	var followee models.User
	err = models.GetUserProfile(base.DB, &followee,  uint(followerId))
	if (err != nil) {
		errCustom := errors.New("User with given follower id doesn't exist")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
	}

	relationId, err := models.FollowUserByUser(base.DB, uid, uint(followerId))
	if (err != nil) {
		errCustom := errors.New("unable to follow user").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	notif_message := "User " + userData.Username + " followed you recently check html link below to see the notification"
	middleware.SendMail(
		"Notification", 
		followee.Firstname, 
		followee.Email, 
		"You got a new follower",
		notif_message,
		"")

	var notification = models.Notification {
		UserID: followee.ID,
		Description: notif_message,
	}

	_, err = models.AddNotification(base.DB, &notification)
	if (err != nil) {
		log.Printf("unable to add notification %v",err)
	}

	middleware.RespondJSON(ctx, http.StatusOK, relationId, nil)
}


func (base *Controller) AddFeedback(ctx *gin.Context) {
	var feedback models.FeedBack

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	err := ctx.ShouldBindJSON(&feedback);
	if err != nil {
		errCustom := errors.New("invalid feedback object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}
	fid, err := models.AddFeedback(base.DB, &feedback)
	if err != nil {
		errCustom := errors.New("unable to add user feedback").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, fid, nil)
	}
}

func (base *Controller) GetFeedback(ctx *gin.Context) {
	var feedback models.FeedBack
	id := ctx.Params.ByName("userId")
	// uid := middleware.GetUidFromToken(ctx)
	uid, err1 := strconv.Atoi(id)
	if err1 != nil  {
		return
	}
	err := models.GetFeedback(base.DB, &feedback, uid)

	if err != nil {
		errCustom := errors.New("unable to retrieve user feedback with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		if int(feedback.ID) != uid {
			errCustom := errors.New("feedback doesn't belong to the given user").Error()
			middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
			return
		}	
		middleware.RespondJSON(ctx, http.StatusOK, feedback, err)
	}
}