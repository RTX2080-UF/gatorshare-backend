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

func (base *Controller) GetAllComment(ctx *gin.Context) {
	var comments []models.Comment
	var pid uint
	pid_str := ctx.Params.ByName("postId")
	
	pidParam, err := strconv.Atoi(pid_str)
    if err != nil {
		errCustom := errors.New("invalid post Id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return    
	}

	if pidParam <= 0  {
		errCustom := errors.New("invalid tag id provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
    } else {
		pid = uint(pidParam)
	}


	err = models.GetAllComment(base.DB, &comments, uint(pid))
	if err != nil {
		errCustom := errors.New("unable to retrieve comment for given post").Error()
		middleware.RespondJSON(ctx, http.StatusNotFound, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comments, nil)
	}
}

func (base *Controller) AddNewComment(ctx *gin.Context) {
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
	CommentId, err := models.AddNewComment(base.DB, &commentDbObj)
	if err != nil {
		errCustom := errors.New("unable to add new comment").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	var userData models.User
	err = models.GetUserProfile(base.DB, &userData, uid)
	if err != nil {
		log.Println("Unable to retrieve userdetails")
	}

	var post models.Post
	err = models.GetOnePost(base.DB, &post, int(commentDbObj.PostID))
	if err != nil {
		log.Println("Unable to retrieve post")
	}

	notif_message := "User " + userData.Username + " commented to your post"
	middleware.SendMail(
		"Notification", 
		post.User.Firstname, 
		post.User.Email, 
		"Your post got a new comment",
		notif_message,
		"")

	var notification = models.Notification {
		UserID: post.User.ID,
		Description: notif_message,
	}

	_, err = models.AddNotification(base.DB, &notification)
	if (err != nil) {
		log.Printf("unable to add notification %v",err)
	}

	middleware.RespondJSON(ctx, http.StatusOK, CommentId, nil)
}

func (base *Controller) GetOneComment(ctx *gin.Context) {
	commentIdStr := ctx.Params.ByName("id")
	var comment models.Comment
	var commentId uint
	commentIdParam, err := strconv.Atoi(commentIdStr)

    if err != nil {
		errCustom := errors.New("invalid comment id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	if commentIdParam <= 0  {
		errCustom := errors.New("invalid tag id provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
    } else {
		commentId = uint(commentIdParam)
	}

	err = models.GetOneComment(base.DB, &comment, commentId)
	if err != nil {
		errCustom := errors.New("unable to retrieve comment with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}

func (base *Controller) UpdateComment(ctx *gin.Context) {
	var comment models.Comment
	id := ctx.Params.ByName("id")
	var commentId uint

	commentIdParam, err := strconv.Atoi(id)
    if err != nil {
		errCustom := errors.New("invalid comment id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return    
	}
	
	if commentIdParam <= 0  {
		errCustom := errors.New("invalid tag id provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
    } else {
		commentId = uint(commentIdParam)
	}

	err = models.GetOneComment(base.DB, &comment, commentId)
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

	err = models.UpdateComment(base.DB, &comment)
	if err != nil {
		errCustom := errors.New("unable to update the comment").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}

func (base *Controller) DeleteComment(ctx *gin.Context) {
	var comment models.Comment
	id := ctx.Params.ByName("id")
	var commentId uint

	commentIdParam, err := strconv.Atoi(id)
    if err != nil {
		errCustom := errors.New("invalid comment id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	if commentIdParam <= 0  {
		errCustom := errors.New("invalid tag id provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
    } else {
		commentId = uint(commentIdParam)
	}

	err = models.GetOneComment(base.DB, &comment, commentId)
	if err != nil {
		errCustom := errors.New("unable to find comment with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
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

	err = models.DeleteComment(base.DB, &comment, commentId)
	if err != nil {
		errCustom := errors.New("unable to delete the comment").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, comment, nil)
	}
}