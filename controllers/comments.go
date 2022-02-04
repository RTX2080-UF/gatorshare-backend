package controllers

import (
	"gatorshare/helpers"
	"gatorshare/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllcomment(ctx *gin.Context) {
	var comments []models.Comment
	uid_str := ctx.Params.ByName("userId")
	
	uid, err := strconv.Atoi(uid_str)
    if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, comments, err)
		return    
	}

	err = models.GetAllcomment(&comments, uid)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusNotFound, comments, err)
	} else {
		helpers.RespondJSON(ctx, http.StatusOK, comments, nil)
	}
}

func AddNewcomment(ctx *gin.Context) {
	var comment models.Comment

	log.Print("Got request to add new comment")
	err := ctx.ShouldBindJSON(&comment);
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, comment, err)
		return
	}

	CommentId, err := models.AddNewcomment(&comment)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadGateway, comment, err)
	} else {
		helpers.RespondJSON(ctx, http.StatusOK, CommentId, nil)
	}
}

func GetOnecomment(ctx *gin.Context) {
	commentIdStr := ctx.Params.ByName("id")
	var comment models.Comment
	CommentId, err := strconv.Atoi(commentIdStr)

    if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, comment, err)
		return
    }

	err = models.GetOnecomment(&comment, CommentId)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadGateway, comment, err)
	} else {
		helpers.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}

func Updatecomment(c *gin.Context) {
	var comment models.Comment
	id := c.Params.ByName("id")
	
	CommentId, err := strconv.Atoi(id)
    if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, comment, err)
		return    
	}

	err = c.ShouldBindJSON(&comment);
	if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, comment, err)
		return
	}

	err = models.GetOnecomment(&comment, CommentId)
	if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, comment, err)
		return	
	}
	c.BindJSON(&comment)
	
	err = models.Updatecomment(&comment, CommentId)
	if err != nil {
		helpers.RespondJSON(c, http.StatusBadGateway, comment, err)
	} else {
		helpers.RespondJSON(c, http.StatusOK, comment, nil)
	}
}

func Deletecomment(ctx *gin.Context) {
	var comment models.Comment
	id := ctx.Params.ByName("id")
	
	CommentId, err := strconv.Atoi(id)
    if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, comment, err)
    }

	err = models.GetOnecomment(&comment, CommentId)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadRequest, comment, err)
	}
	ctx.BindJSON(&comment)

	err = models.Deletecomment(&comment, CommentId)
	if err != nil {
		helpers.RespondJSON(ctx, http.StatusBadGateway, comment, err)
	} else {
		helpers.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}