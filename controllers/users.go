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
	var userdata UserRegister

	log.Print("Got request to add new User")
	err := ctx.ShouldBindJSON(&userdata);
	if err != nil {
		errCustom := errors.New("invalid user object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	userdataDb := UserRequestToDBModel(userdata) 
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

	if userData.ID != uid {
		errCustom := errors.New("profile doesn't belong to the given user").Error()
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
		return
	}

	if err != nil {
		errCustom := errors.New("unable to retrieve user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userData, err)
	}
}

func (base *Controller) UpdateProfile(ctx *gin.Context) {
	var userData models.User
	
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	log.Print("Got request to Update User profile", uid)

	err := models.GetUserProfile(base.DB, &userData, uid)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, userData, err)
		return	
	}

	if userData.ID != uid {
		errCustom := errors.New("profile doesn't belong to the given user").Error()
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
		return
	}
	
	err = ctx.ShouldBindJSON(&userData);
	if err != nil {
		errCustom := errors.New("unable to retrieve user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}
	
	err = models.UpdateUserProfile(base.DB, &userData)
	if err != nil {
		errCustom := errors.New("unable to update user profile with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userData, nil)
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
