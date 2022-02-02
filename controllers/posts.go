package controllers

import (
	"gatorshare/helpers"
	"gatorshare/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Listpost(ctx *gin.Context) {
	var posts []models.Post
	uid_str := ctx.Params.ByName("userId")
	
	uid, err := strconv.Atoi(uid_str)
    if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, posts, err)
		return    
	}

	err = models.GetAllpost(&posts, uid)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusNotFound, posts, err)
	} else {
		helpers.RespondJSON(ctx, http.StatusOK, posts, nil)
	}
}

func AddNewpost(ctx *gin.Context) {
	var post models.Post

	log.Print("Got request to add new post")
	err := ctx.ShouldBindJSON(&post);
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return
	}

	postId, err := models.AddNewpost(&post)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadGateway, post, err)
	} else {
		helpers.RespondJSON(ctx, http.StatusOK, postId, nil)
	}
}

func GetOnepost(ctx *gin.Context) {
	postIdStr := ctx.Params.ByName("id")
	var post models.Post
	postId, err := strconv.Atoi(postIdStr)

    if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return
    }

	err = models.GetOnepost(&post, postId)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadGateway, post, err)
	} else {
		helpers.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Params.ByName("id")
	
	PostId, err := strconv.Atoi(id)
    if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, post, err)
		return    
	}

	err = c.ShouldBindJSON(&post);
	if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, post, err)
		return
	}

	err = models.GetOnepost(&post, PostId)
	if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, post, err)
		return	
	}
	c.BindJSON(&post)
	
	err = models.UpdatePost(&post, PostId)
	if err != nil {
		helpers.RespondJSON(c, http.StatusBadGateway, post, err)
	} else {
		helpers.RespondJSON(c, http.StatusOK, post, nil)
	}
}

func Deletepost(ctx *gin.Context) {
	var post models.Post
	id := ctx.Params.ByName("id")
	
	PostId, err := strconv.Atoi(id)
    if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, post, err)
    }

	err = models.GetOnepost(&post, PostId)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, post, err)
	}
	ctx.BindJSON(&post)

	err = models.Deletepost(&post, PostId)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadGateway, post, err)
	} else {
		helpers.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}