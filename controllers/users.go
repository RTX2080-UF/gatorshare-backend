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

func (base *Controller) Register(ctx *gin.Context) {
	var userdata UserProfile

	log.Print("Got request to add new User")
	err := ctx.ShouldBindJSON(&userdata);
	if err != nil {
		errCustom := errors.New("invalid user object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	userdataDb := UserRequestToDBModel(userdata) 
	// if (!middleware.IsValidPassword(userdata.Password)) {
	// 	errCustom := errors.New("user password doesn't satisfy minimum requirement").Error()
	// 	middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	// 	return
	// }

	userdataDb.Password, err = middleware.HashPassword(userdataDb.Password)
	if (err != nil) {
		errCustom := errors.New("invalid user password provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	userId, err := models.AddNewUser(base.DB, &userdataDb)
	if err != nil {
		errCustom := errors.New("unable to register user").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userId, nil)
	}
}

func (base *Controller) Login(ctx *gin.Context) {

	var loginDetails Login
	if err := ctx.ShouldBindJSON(&loginDetails); err != nil {
		log.Println("Login", err.Error())
		errCustom := errors.New("invalid input provided")
		middleware.RespondJSON(ctx, http.StatusUnprocessableEntity, errCustom ,err)
	   	return
	}
	log.Print("Got request to get User profile")
	
	userObj, err := models.GetUserDetailByUsername(base.DB, loginDetails.Username)
	isPasswordValid := middleware.CheckPasswordHash(userObj.Password, loginDetails.Password)
	if err != nil || !isPasswordValid {
		if err != nil {
			log.Println("Login", err.Error())
		}
		errCustom := errors.New("unable to authenticate user")
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, errCustom)
		return
	}

	token, err := middleware.CreateToken(userObj.ID)
	if err != nil {
		errCustom := errors.New("unable to generate token")
		log.Println("Login", err.Error())
		middleware.RespondJSON(ctx, http.StatusUnprocessableEntity, errCustom, err)
		return
	}

	middleware.RespondJSON(ctx, http.StatusOK, token, nil)
}

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
	if (!middleware.CheckPasswordHash(currentUserData.Password, newUserData.OldPassword)) {
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
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, updatedUserData, nil)
	}
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
	followerId, err := strconv.Atoi(userIdStr)
    if err != nil {
		errCustom := errors.New("invalid user id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	isExist, err := models.CheckUserExists(base.DB, uint(followerId))
	if (!isExist || err != nil) {
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

	middleware.RespondJSON(ctx, http.StatusOK, relationId, nil)
}


func (base *Controller) RefreshToken(ctx *gin.Context) {
	token := middleware.ExtractToken(ctx)
	if token == "" {
		errCustom := errors.New("Unable to unravel token, invalid token provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
	}

	newToken := middleware.RefreshToken(token)
	if newToken == "" {
		errCustom := errors.New("Token expired, please login again")
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom.Error(), errCustom)
		return
	}

	middleware.RespondJSON(ctx, http.StatusOK, newToken, nil)
}