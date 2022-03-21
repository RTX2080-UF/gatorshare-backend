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

func getUidFromToken(ctx *gin.Context) uint {
	token := middleware.ExtractToken(ctx)
	errCustom := errors.New("invalid token provided")

	if (token != "") {
		err := middleware.TokenValid(token)
		if (err != nil) {
			middleware.RespondJSON(ctx, http.StatusForbidden, errCustom, err)
			return 0
		}
	} else {
		middleware.RespondJSON(ctx, http.StatusForbidden, errCustom, nil)
		return 0
	}
	
	uid, err := middleware.ExtractTokenID(token)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusForbidden, errCustom, err)
		return 0
	}

	return uid
}

func (base *Controller) Listpost(ctx *gin.Context) {
	uid := getUidFromToken(ctx)
	if uid == 0 {
		return
	}

	var posts []models.Post
	err := models.GetAllpost(base.DB, &posts, uid)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusNotFound, posts, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, posts, nil)
	}
}

func (base *Controller) AddNewpost(ctx *gin.Context) {
	var post Post
	log.Print("Got request to add new post")
	err := ctx.ShouldBindJSON(&post);
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return
	}

	uid := getUidFromToken(ctx)
	if uid == 0 {
		return
	}
	post_model := PostRequestToDBModel(post, uid)

	postId, err := models.AddNewpost(base.DB, &post_model)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, post, err)
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
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return
    }

	err = models.GetOnepost(base.DB, &post, postId)
	post.User.Password = ""  
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, nil, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, err)
	}
}

func (base *Controller) UpdatePost(ctx *gin.Context) {
	var post models.Post
	id := ctx.Params.ByName("id")

	uid := getUidFromToken(ctx)
	if uid == 0 {
		return
	}

	PostId, err := strconv.Atoi(id)
    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return    
	}
	
	err = models.GetOnepost(base.DB, &post, PostId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return	
	}

	if (post.User.ID != uid) {
		middleware.RespondJSON(ctx, http.StatusForbidden, post, err)
		return
	}
	
	err = ctx.ShouldBindJSON(&post);
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return
	}
	
	err = models.UpdatePost(base.DB, &post)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, post, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}

func (base *Controller) Deletepost(ctx *gin.Context) {
	var post models.Post
	id := ctx.Params.ByName("id")
	
	uid := getUidFromToken(ctx)
	if uid == 0 {
		return
	}

	PostId, err := strconv.Atoi(id)
    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
    }

	err = models.GetOnepost(base.DB, &post, PostId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
	}
	ctx.BindJSON(&post)

	if (post.User.ID != uid) {
		middleware.RespondJSON(ctx, http.StatusForbidden, post, err)
		return
	}

	err = models.Deletepost(base.DB, &post, PostId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, post, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}