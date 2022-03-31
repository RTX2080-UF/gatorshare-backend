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

func (base *Controller) GetTag(ctx *gin.Context) {
	var tag models.Tag
	tid_str := ctx.Params.ByName("tagId")
	
	tid, err := strconv.Atoi(tid_str)
    if err != nil {
		errCustom := errors.New("invalid tag Id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return    
	}

	err = models.GetTag(base.DB, &tag, uint(tid))
	if err != nil {
		errCustom := errors.New("unable to retrieve tags for given id").Error()
		middleware.RespondJSON(ctx, http.StatusNotFound, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, tag, nil)
	}
}

func (base *Controller) AddTag(ctx *gin.Context) {
	var tag Tag

	log.Print("Got request to add new comment")
	err := ctx.ShouldBindJSON(&tag);
	if err != nil {
		errCustom := errors.New("invalid tag object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	tagDbObj := TagRequestToDBModel(tag, uid)
	tagId, err := models.AddNewTag(base.DB, &tagDbObj)
	if err != nil {
		errCustom := errors.New("unable to add new tag").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, tagId, nil)
	}
}



func (base *Controller) UpdateTag(ctx *gin.Context) {
	var tag models.Tag
	id := ctx.Params.ByName("id")
	
	tagId, err := strconv.Atoi(id)
    if err != nil {
		errCustom := errors.New("invalid tag id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return    
	}
	
	err = models.GetTag(base.DB, &tag, uint(tagId))
	if err != nil {
		errCustom := errors.New("unable to find tag with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return	
	}

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	if tag.CreatorId != uid {
		errCustom := errors.New("user is not the tag author")
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom.Error(), errCustom)
		return
	}

	err = ctx.ShouldBindJSON(&tag);
	if err != nil {
		errCustom := errors.New("invalid tag object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	err = models.UpdateTag(base.DB, &tag)
	if err != nil {
		errCustom := errors.New("unable to update the tag").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, tag, nil)
	}
}

func (base *Controller) DeleteTag(ctx *gin.Context) {
	var tag models.Tag
	var tagId uint
	id := ctx.Params.ByName("id")
	
	tagIdParam, err := strconv.Atoi(id)
    if err != nil {
		errCustom := errors.New("invalid tag id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
    }

	if tagIdParam <= 0  {
		errCustom := errors.New("invalid tag id provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
    } else {
		tagId = uint(tagIdParam)
	}

	err = models.GetTag(base.DB, &tag, tagId)
	if err != nil {
		errCustom := errors.New("unable to find tag with given id").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
	}
	ctx.BindJSON(&tag)

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	if tag.CreatorId != uid {
		errCustom := errors.New("user is not the tag author").Error()
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, err)
		return
	}

	err = models.DeleteTag(base.DB, tagId)
	if err != nil {
		errCustom := errors.New("unable to delete the tag").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, tag, nil)
	}
}

func (base *Controller) FollowTagsByUser(ctx *gin.Context) {
	tid_str := ctx.Params.ByName("tagId")
	tid, err := strconv.Atoi(tid_str)
    if err != nil {
		errCustom := errors.New("invalid tag Id provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return    
	}

	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	usertagMapping, err := models.FollowTagsByUser(base.DB, uid, uint(tid))
	if err != nil {
		errCustom := errors.New("unable to associate tag with given id").Error()
		middleware.RespondJSON(ctx, http.StatusNotFound, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, usertagMapping, nil)
	}
}

func (base *Controller) PopularTags(ctx *gin.Context){
	
	
}