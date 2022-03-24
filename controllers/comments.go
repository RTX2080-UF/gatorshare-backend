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

func (base *Controller) GetAllcomment(ctx *gin.Context) {
	var comments []models.Comment
	pid_str := ctx.Params.ByName("postId")
	
	pid, err := strconv.Atoi(pid_str)
    if err != nil {
		errCustom := errors.New("invalid post Id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return    
	}

	err = models.GetAllcomment(base.DB, &comments, pid)
	if err != nil {
		errCustom := errors.New("unable to retrieve comment for given post").Error()
		middleware.RespondJSON(ctx, http.StatusNotFound, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comments, nil)
	}
}

func (base *Controller) AddNewcomment(ctx *gin.Context) {
	var comment Comment

	log.Print("Got request to add new comment")
	err := ctx.ShouldBindJSON(&comment);
	if err != nil {
		errCustom := errors.New("invalid comment object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	commentDbObj := CommentRequestToDBModel(comment, uid)
	CommentId, err := models.AddNewcomment(base.DB, &commentDbObj)
	if err != nil {
		errCustom := errors.New("unable to add new comment").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, CommentId, nil)
	}
}

func (base *Controller) GetOnecomment(ctx *gin.Context) {
	commentIdStr := ctx.Params.ByName("id")
	var comment models.Comment
	CommentId, err := strconv.Atoi(commentIdStr)

    if err != nil {
		errCustom := errors.New("invalid comment id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	err = models.GetOnecomment(base.DB, &comment, CommentId)
	if err != nil {
		errCustom := errors.New("unable to retrieve comment with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}

func (base *Controller) Updatecomment(ctx *gin.Context) {
	var comment models.Comment
	id := ctx.Params.ByName("id")
	
	CommentId, err := strconv.Atoi(id)
    if err != nil {
		errCustom := errors.New("invalid comment id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return    
	}
	
	err = models.GetOnecomment(base.DB, &comment, CommentId)
	if err != nil {
		errCustom := errors.New("unable to find comment with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return	
	}

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	if comment.UserID != uid {
		errCustom := errors.New("user is not the comment author").Error()
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
		return
	}

	err = ctx.ShouldBindJSON(&comment);
	if err != nil {
		errCustom := errors.New("invalid comment object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	err = models.Updatecomment(base.DB, &comment)
	if err != nil {
		errCustom := errors.New("unable to update the comment").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}

func (base *Controller) Deletecomment(ctx *gin.Context) {
	var comment models.Comment
	id := ctx.Params.ByName("id")
	
	CommentId, err := strconv.Atoi(id)
    if err != nil {
		errCustom := errors.New("invalid comment id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
    }

	err = models.GetOnecomment(base.DB, &comment, CommentId)
	if err != nil {
		errCustom := errors.New("unable to find comment with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
	}
	ctx.BindJSON(&comment)

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	if comment.UserID != uid {
		errCustom := errors.New("user is not the comment author").Error()
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
		return
	}

	err = models.Deletecomment(base.DB, &comment, CommentId)
	if err != nil {
		errCustom := errors.New("unable to delete the comment").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}