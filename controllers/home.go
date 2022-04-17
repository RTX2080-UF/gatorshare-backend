package controllers

import (
	"errors"
	"gatorshare/middleware"
	"gatorshare/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (base *Controller) GetUserHome(ctx *gin.Context) {
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	var posts []models.Post

	tagPosts, err := models.GetUserHomePosts(base.DB, uid)
	if err != nil {
		errCustom := errors.New("unable to retrieve post for user").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}
 
	if (len(tagPosts) == 0) {
		errCustom := errors.New("no new posts for user")
		middleware.RespondJSON(ctx, http.StatusNoContent, errCustom.Error(), errCustom)
		return
	}

	for _, element := range tagPosts {
		posts = append(posts, element.Post)
	}
	
	middleware.RespondJSON(ctx, http.StatusOK, posts, nil)
}