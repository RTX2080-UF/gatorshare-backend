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

func (base *Controller) Listpost(ctx *gin.Context) {
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	var posts []models.Post
	err := models.GetAllpost(base.DB, &posts, uid)
	if err != nil {
		errCustom := errors.New("no post exist for given user").Error()
		middleware.RespondJSON(ctx, http.StatusNotFound, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, posts, nil)
	}
}

func (base *Controller) AddNewpost(ctx *gin.Context) {
	var post Post
	log.Print("Got request to add new post")
	err := ctx.ShouldBindJSON(&post);
	if err != nil {
		errCustom := errors.New("invalid post request").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	post_model := PostRequestToDBModel(post, uid)

	postId, err := models.AddNewpost(base.DB, &post_model)
	if err != nil {
		errCustom := errors.New("unable to add new post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, postId, nil)
	}
}

func (base *Controller) GetOnepost(ctx *gin.Context) {
	postIdStr := ctx.Params.ByName("id")
	var post models.Post
	postId, err := strconv.Atoi(postIdStr)
	log.Print("Got request to get post")

    if err != nil {
		errCustom := errors.New("invalid post id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	err = models.GetOnepost(base.DB, &post, postId)
	post.User.Password = ""  
	if err != nil {
		errCustom := errors.New("not able to retrieve post with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, err)
	}
}

func (base *Controller) UpdatePost(ctx *gin.Context) {
	var post models.Post
	id := ctx.Params.ByName("id")

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	PostId, err := strconv.Atoi(id)
    if err != nil {
		errCustom := errors.New("invalid post Id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return    
	}
	
	err = models.GetOnepost(base.DB, &post, PostId)
	if err != nil {
		errCustom := errors.New("unable to get post with given ID").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return	
	}

	if (post.User.ID != uid) {
		errCustom := errors.New("user is not the author of post").Error()
		middleware.RespondJSON(ctx, http.StatusForbidden, errCustom, err)
		return
	}
	
	err = ctx.ShouldBindJSON(&post);
	if err != nil {
		errCustom := errors.New("invalid post object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}
	
	err = models.UpdatePost(base.DB, &post)
	if err != nil {
		errCustom := errors.New("unable to update the post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}

func (base *Controller) Deletepost(ctx *gin.Context) {
	var post models.Post
	id := ctx.Params.ByName("id")
	
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	PostId, err := strconv.Atoi(id)
    if err != nil {
		errCustom := errors.New("invalid post id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	err = models.GetOnepost(base.DB, &post, PostId)
	if err != nil {
		errCustom := errors.New("unable to retrieve post with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}
	ctx.BindJSON(&post)

	if (post.User.ID != uid) {
		errCustom := errors.New("user is not the author of post").Error()
		middleware.RespondJSON(ctx, http.StatusForbidden, errCustom, err)
		return
	}

	err = models.Deletepost(base.DB, &post, PostId)
	if err != nil {
		errCustom := errors.New("unable to delete the post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}