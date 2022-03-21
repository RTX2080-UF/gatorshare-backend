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
		middleware.RespondJSON(ctx, http.StatusBadRequest, userdata, err)
		return
	}

	userdataDb := UserRequestToDBModel(userdata) 
	userdataDb.Password, err = middleware.HashPassword(userdataDb.Password)
	if (err != nil) {
		middleware.RespondJSON(ctx, http.StatusBadGateway, userdata, err)
		return
	}

	userId, err := models.AddNewUser(base.DB, &userdataDb)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, userdata, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userId, nil)
	}
}

func (base *Controller) GetProfile(ctx *gin.Context) {
	var userData models.User
	userIdStr := ctx.Params.ByName("id")
	userId, err := strconv.Atoi(userIdStr)
	log.Print("Got request to get User profile", userId)

    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, userData, err)
		return
    }

	err = models.GetUserProfile(base.DB, &userData, userId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, nil, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userData, err)
	}
}

func (base *Controller) UpdateProfile(ctx *gin.Context) {
	var userData models.User
	userIdStr := ctx.Params.ByName("id")
	log.Print("Got request to Update User profile", userIdStr)

	userId, err := strconv.Atoi(userIdStr)
    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, userData, err)
		return    
	}
	
	err = models.GetUserProfile(base.DB, &userData, userId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, userData, err)
		return	
	}
	
	err = ctx.ShouldBindJSON(&userData);
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, userData, err)
		return
	}
	
	err = models.UpdateUserProfile(base.DB, &userData)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, userData, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userData, nil)
	}
}

func (base *Controller) DeleteUser(ctx *gin.Context) {
	var userData models.User
	userIdStr := ctx.Params.ByName("id")
	
	userId, err := strconv.Atoi(userIdStr)
    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, userData, err)
    }

	err = models.GetUserProfile(base.DB, &userData, userId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, userData, err)
	}
	ctx.BindJSON(&userData)

	err = models.DeleteUser(base.DB, userId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, userData, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userData, nil)
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