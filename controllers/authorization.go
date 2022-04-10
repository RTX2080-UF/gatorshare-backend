package controllers

import (
	"errors"
	"gatorshare/middleware"
	"gatorshare/models"
	"log"
	"net/http"

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

func (base *Controller) ResetPassword(ctx *gin.Context) {
}