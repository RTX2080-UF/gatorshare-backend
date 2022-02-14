package controllers

import (
	"gatorshare/middleware"
	"gatorshare/models"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (base *Controller) Listpost(ctx *gin.Context) {
	var posts []models.Post
	uid_str := ctx.Params.ByName("userId")
	
	uid, err := strconv.Atoi(uid_str)
    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, posts, err)
		return    
	}

	err = models.GetAllpost(base.DB, &posts, uid)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusNotFound, posts, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, posts, nil)
	}
}

func (base *Controller) AddNewpost(ctx *gin.Context) {
	var post models.Post

	log.Print("Got request to add new post")
	err := ctx.ShouldBindJSON(&post);
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return
	}

	postId, err := models.AddNewpost(base.DB, &post)
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

    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
		return
    }

	err = models.GetOnepost(base.DB, &post, postId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, post, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}

func (base *Controller) UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Params.ByName("id")
	
	PostId, err := strconv.Atoi(id)
    if err != nil {
		middleware.RespondJSON(c, http.StatusBadRequest, post, err)
		return    
	}

	err = c.ShouldBindJSON(&post);
	if err != nil {
		middleware.RespondJSON(c, http.StatusBadRequest, post, err)
		return
	}

	err = models.GetOnepost(base.DB, &post, PostId)
	if err != nil {
		middleware.RespondJSON(c, http.StatusBadRequest, post, err)
		return	
	}
	c.BindJSON(&post)
	
	err = models.UpdatePost(base.DB, &post, PostId)
	if err != nil {
		middleware.RespondJSON(c, http.StatusBadGateway, post, err)
	} else {
		middleware.RespondJSON(c, http.StatusOK, post, nil)
	}
}

func (base *Controller) Deletepost(ctx *gin.Context) {
	var post models.Post
	id := ctx.Params.ByName("id")
	
	PostId, err := strconv.Atoi(id)
    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
    }

	err = models.GetOnepost(base.DB, &post, PostId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, post, err)
	}
	ctx.BindJSON(&post)

	err = models.Deletepost(base.DB, &post, PostId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, post, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}