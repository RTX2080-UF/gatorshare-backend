package controllers

import (
	"gatorshare/middleware"
	"gatorshare/models"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (base *Controller) GetAllcomment(ctx *gin.Context) {
	var comments []models.Comment
	uid_str := ctx.Params.ByName("userId")
	
	uid, err := strconv.Atoi(uid_str)
    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, comments, err)
		return    
	}

	err = models.GetAllcomment(base.DB, &comments, uid)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusNotFound, comments, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comments, nil)
	}
}

func (base *Controller) AddNewcomment(ctx *gin.Context) {
	var comment models.Comment

	log.Print("Got request to add new comment")
	err := ctx.ShouldBindJSON(&comment);
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, comment, err)
		return
	}

	CommentId, err := models.AddNewcomment(base.DB, &comment)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, comment, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, CommentId, nil)
	}
}

func (base *Controller) GetOnecomment(ctx *gin.Context) {
	commentIdStr := ctx.Params.ByName("id")
	var comment models.Comment
	CommentId, err := strconv.Atoi(commentIdStr)

    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, comment, err)
		return
    }

	err = models.GetOnecomment(base.DB, &comment, CommentId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, comment, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}

func (base *Controller) Updatecomment(c *gin.Context) {
	var comment models.Comment
	id := c.Params.ByName("id")
	
	CommentId, err := strconv.Atoi(id)
    if err != nil {
		middleware.RespondJSON(c, http.StatusBadRequest, comment, err)
		return    
	}

	err = c.ShouldBindJSON(&comment);
	if err != nil {
		middleware.RespondJSON(c, http.StatusBadRequest, comment, err)
		return
	}

	err = models.GetOnecomment(base.DB, &comment, CommentId)
	if err != nil {
		middleware.RespondJSON(c, http.StatusBadRequest, comment, err)
		return	
	}
	c.BindJSON(&comment)
	
	err = models.Updatecomment(base.DB, &comment, CommentId)
	if err != nil {
		middleware.RespondJSON(c, http.StatusBadGateway, comment, err)
	} else {
		middleware.RespondJSON(c, http.StatusOK, comment, nil)
	}
}

func (base *Controller) Deletecomment(ctx *gin.Context) {
	var comment models.Comment
	id := ctx.Params.ByName("id")
	
	CommentId, err := strconv.Atoi(id)
    if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, comment, err)
    }

	err = models.GetOnecomment(base.DB, &comment, CommentId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadRequest, comment, err)
	}
	ctx.BindJSON(&comment)

	err = models.Deletecomment(base.DB, &comment, CommentId)
	if err != nil {
		middleware.RespondJSON(ctx, http.StatusBadGateway, comment, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}