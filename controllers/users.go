package controllers

import (
	"gatorshare/middleware"
	"gatorshare/models"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (base *Controller) Register(ctx *gin.Context) {
	var userdata models.User

	log.Print("Got request to add new User")
	err := ctx.ShouldBindJSON(&userdata);
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, userdata, err)
		return
	}

	userId, err := models.AddNewUser(base.DB, &userdata)
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