package controllers

import (
	"errors"
	"fmt"
	"gatorshare/middleware"
	"gatorshare/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (base *Controller) ListPost(ctx *gin.Context) {
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	var posts []models.Post
	err := models.GetAllPost(base.DB, &posts, uid)
	if err != nil {
		errCustom := errors.New("no post exist for given user").Error()
		middleware.RespondJSON(ctx, http.StatusNotFound, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, posts, nil)
	}
}

func (base *Controller) AddNewPost(ctx *gin.Context) {
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

	postId, err := models.AddNewPost(base.DB, &post_model)
	if err != nil {
		errCustom := errors.New("unable to add new post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	var tagsList []models.Tag
	for _, element := range post.Tags {
		var tagObj = models.Tag{CreatorId : uid , Name : element}	
		tagsList = append(tagsList, tagObj)	
	}
	tagsIDs, err := models.InsertTags(base.DB, tagsList)
	if err != nil {
		errCustom := errors.New("unable to add tags present in the post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}
	
	err = models.AddPostTags(base.DB, postId, tagsIDs)
	if err != nil {
		errCustom := errors.New("unable to associate tags with the post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	middleware.RespondJSON(ctx, http.StatusOK, postId, nil)
}

func (base *Controller) GetOnePost(ctx *gin.Context) {
	postIdStr := ctx.Params.ByName("id")
	var post models.Post
	postId, err := strconv.Atoi(postIdStr)
	log.Print("Got request to get post")

    if err != nil {
		errCustom := errors.New("invalid post id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	err = models.GetOnePost(base.DB, &post, postId)
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
	
	err = models.GetOnePost(base.DB, &post, PostId)
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

func (base *Controller) DeletePost(ctx *gin.Context) {
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

	err = models.GetOnePost(base.DB, &post, PostId)
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

	err = models.DeletePost(base.DB, &post, PostId)
	if err != nil {
		errCustom := errors.New("unable to delete the post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, post, nil)
	}
}

func (base *Controller) ReactToPost(ctx *gin.Context) {
	postIdStr := ctx.PostForm("postid")
	reaction := ctx.PostForm("reaction")

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	
	postId, err := middleware.ConvertStrToInt(postIdStr)
	if err != nil {
		errCustom := errors.New("Invalid post Id provided").Error()
		middleware.RespondJSON(ctx, http.StatusForbidden, errCustom, err)
		return
	}

	var post_reaction = models.UserPost {
		UserID: uid,
		PostID: uint(postId),
		Reaction: models.ReactionType(reaction),
	}

	reactionId, err := models.ReactToPost(base.DB, &post_reaction)
	if err != nil {
		errCustom := errors.New("unable to add reaction to post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	} 
	
	var userData models.User
	err = models.GetUserProfile(base.DB, &userData, uid)
	if err != nil {
		log.Println("Unable to retrieve userdetails")
	}

	var post models.Post
	err = models.GetOnePost(base.DB, &post, int(postId))
	if err != nil {
		log.Println("Unable to retrieve post")
	}

	notif_message := "User " + userData.Username + " reacted on your post"
	middleware.SendMail(
		"Notification", 
		post.User.Firstname, 
		post.User.Email, 
		"Your post got a new reaction",
		notif_message,
		"")

	var notification = models.Notification {
		UserID: userData.ID,
		Description: notif_message,
	}

	_, err = models.AddNotification(base.DB, &notification)
	if (err != nil) {
		log.Printf("unable to add notification %v",err)
	}

	middleware.RespondJSON(ctx, http.StatusOK, reactionId, nil)
}

func (base *Controller) GetPostReaction(ctx *gin.Context) {
	postIdStr := ctx.Params.ByName("postId")

	postId, err := middleware.ConvertStrToInt(postIdStr)
	if err != nil {
		errCustom := errors.New("Invalid post Id provided").Error()
		middleware.RespondJSON(ctx, http.StatusForbidden, errCustom, err)
		return
	}

	var reactionList[] models.UserPost
	err = models.GetReactions(base.DB, postId, &reactionList)
	if err != nil {
		errCustom := errors.New("unable to get reaction for post").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, reactionList, nil)
	}
}


func (base *Controller) SearchPost(ctx *gin.Context){
	var verifiedPosts []models.Post
	var tagIds []uint
	var tagList SearchPostReq
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	
	err := ctx.ShouldBindJSON(&tagList);
	if err != nil {
		errCustom := errors.New("invalid tag object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	tagIds = models.SearchTagIdHelper(base.DB, tagList.TagNames)
	if(len(tagIds) == 0){
		errCustom := errors.New("Unable to fetch Tags")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
	}else{
		err = models.SearchPost(base.DB, &verifiedPosts, tagIds)
		if err != nil {
			errCustom := errors.New("unable to associate post with given tag id").Error()
			middleware.RespondJSON(ctx, http.StatusNotFound, errCustom, err)
			return
		} 
		fmt.Printf("posts------------ %v", verifiedPosts)
		middleware.RespondJSON(ctx, http.StatusOK, verifiedPosts, nil)
	}

}