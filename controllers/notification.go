package controllers

import (
	"errors"
	"gatorshare/middleware"
	"gatorshare/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (base *Controller) GetNewNotifications(ctx *gin.Context) {
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}
	var notifications []models.Notification

	err := models.GetNewNotifications(base.DB, &notifications, uid)
	if err != nil {
		errCustom := errors.New("unable to get new notifications").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}
 
	if (len(notifications) == 0) {
		errCustom := errors.New("no new notifications for user")
		middleware.RespondJSON(ctx, http.StatusNoContent, errCustom.Error(), errCustom)
		return
	}
	
	middleware.RespondJSON(ctx, http.StatusOK, notifications, nil)
}

func (base *Controller) UpdateNotifications(ctx *gin.Context) {
	uid := middleware.GetUidFromToken(ctx)
	if uid == 0 {
		return
	}

	err := models.UpdateNotifications(base.DB, uid)
	if err != nil {
		errCustom := errors.New("unable to update notifications").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, true, nil)
	}
}
