package controllers

import (
	"crypto/rand"
	"errors"
	"fmt"
	"gatorshare/middleware"
	"gatorshare/models"
	"log"
	"math/big"
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
)

func (base *Controller) Register(ctx *gin.Context) {
	var userdata UserProfile

	log.Print("Got request to add new User")
	err := ctx.ShouldBindJSON(&userdata)
	if err != nil {
		errCustom := errors.New("invalid user object provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}

	userdataDb := UserRequestToDBModel(userdata)
	// if (!middleware.IsValidPassword(userdata.Password)) {
	// 	errCustom := errors.New("user password doesn't satisfy minimum requirement").Error()
	// 	middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	// 	return
	// }

	userdataDb.Password, err = middleware.HashPassword(userdataDb.Password)
	if err != nil {
		errCustom := errors.New("invalid user password provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	userId, err := models.AddNewUser(base.DB, &userdataDb)
	if err != nil {
		errCustom := errors.New("unable to register user").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
	} else {
		middleware.RespondJSON(ctx, http.StatusOK, userId, nil)
	}
}

func (base *Controller) Login(ctx *gin.Context) {

	var loginDetails Login
	if err := ctx.ShouldBindJSON(&loginDetails); err != nil {
		log.Println("Login", err.Error())
		errCustom := errors.New("invalid input provided")
		middleware.RespondJSON(ctx, http.StatusUnprocessableEntity, errCustom, err)
		return
	}
	log.Print("Got request to get User profile")

	userObj, err := models.GetUserDetailByUsername(base.DB, loginDetails.Username)
	isPasswordValid := middleware.CheckPasswordHash(userObj.Password, loginDetails.Password)
	if err != nil || !isPasswordValid {
		if err != nil {
			log.Println("Login", err.Error())
		}
		errCustom := errors.New("unable to authenticate user")
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom, errCustom)
		return
	}

	token, err := middleware.CreateToken(userObj.ID)
	if err != nil {
		errCustom := errors.New("unable to generate token")
		log.Println("Login", err.Error())
		middleware.RespondJSON(ctx, http.StatusUnprocessableEntity, errCustom, err)
		return
	}

	userLikedTags, _ := models.GetUserLikedTags(base.DB, userObj.ID)

	userObj.Password = ""
	var respose = LoginResponse{
		Data:  userObj,
		Token: token,
		Tag: userLikedTags,
	}

	middleware.RespondJSON(ctx, http.StatusOK, respose, nil)
}

func (base *Controller) RefreshToken(ctx *gin.Context) {
	token := middleware.ExtractToken(ctx)
	if token == "" {
		errCustom := errors.New("Unable to unravel token, invalid token provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
	}

	newToken := middleware.RefreshToken(token)
	if newToken == "" {
		errCustom := errors.New("Token expired, please login again")
		middleware.RespondJSON(ctx, http.StatusUnauthorized, errCustom.Error(), errCustom)
		return
	}

	middleware.RespondJSON(ctx, http.StatusOK, newToken, nil)
}

func (base *Controller) ResetPassword(ctx *gin.Context) {
	emailStr := ctx.DefaultQuery("email", "")
	_, err := mail.ParseAddress(emailStr)
	envSrc := middleware.LoadEnv(".env")
	hostPath := middleware.GetEnv("FRONTEND_HOST_PATH", "localhost:8080", envSrc)

	if err != nil {
		errCustom := errors.New("Invalid email address provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), err)
		return
	}

	userDetails, err := models.GetUserDetailByEmail(base.DB, emailStr)
	if err != nil {
		errCustom := errors.New("User doesn't exist")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
	}

	rnum, _ := rand.Int(rand.Reader, big.NewInt(100000))
	randStr, err := middleware.HashPassword(fmt.Sprint(time.Now().UnixNano()) + rnum.String() + fmt.Sprint(userDetails.ID))

	var resetObj = models.ResetPassword{
		UserID:       userDetails.ID,
		Status:       true,
		UniqueRndStr: randStr,
	}

	response, err := models.UpdatResetPassword(base.DB, resetObj)
	if err != nil {
		errCustom := errors.New("Unable to reset password")
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	resetUrl := "<a href='" + scheme + "://" + hostPath + "/v1/users/changePassword?email=" + userDetails.Email + "&token=" + randStr + "'> Reset Password </a>"

	response = middleware.SendMail(
		"Admin",
		userDetails.Firstname,
		userDetails.Email,
		"Password reset for Gatorshare ",
		"You have requested password reset for"+userDetails.Username+"Please follow link below to reset your passowrd",
		resetUrl)

	if !response {
		errCustom := errors.New("Unable to generate password reset link")
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom.Error(), errCustom)
		return
	}

	middleware.RespondJSON(ctx, http.StatusOK, response, nil)
}

func (base *Controller) UpdatePassword(ctx *gin.Context) {
	var updatePass UpdatePasswordReq 
	err := ctx.ShouldBindJSON(&updatePass);
	if err != nil {
		errCustom := errors.New("invalid updatePass request").Error()
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom, err)
		return
	}
	
	emailStr := updatePass.Email
	tokenStr := updatePass.Token

	_, err = mail.ParseAddress(emailStr)
	if err != nil {
		errCustom := errors.New("Invalid email address provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), err)
		return
	}

	if len(tokenStr) >= 64 {
		errCustom := errors.New("Invalid token provided")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
	}

	userData, err := models.GetUserDetailByEmail(base.DB, emailStr)
	if err != nil {
		errCustom := errors.New("Unable to get user details")
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom.Error(), err)
		return
	}

	resp, err := models.VerifyPasswordReset(base.DB, userData.ID, tokenStr)
	if err != nil {
		errCustom := errors.New("Unable to authenticate user")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), err)
		return
	} else if !resp {
		errCustom := errors.New("Password reset request is invalid")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), errCustom)
		return
	}

	newPassword := ctx.PostForm("password")
	if len(newPassword) > 128 || len(newPassword) <= 6 {
		errCustom := errors.New("Password must be between 6 and 128 characters ")
		middleware.RespondJSON(ctx, http.StatusBadRequest, errCustom.Error(), err)
		return
	}

	newPassword, err = middleware.HashPassword(newPassword)
	if err != nil {
		errCustom := errors.New("invalid user password provided").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	err = models.UpdatPassword(base.DB, userData.ID, newPassword)
	if err != nil {
		errCustom := errors.New("Unable to update password").Error()
		middleware.RespondJSON(ctx, http.StatusBadGateway, errCustom, err)
		return
	}

	middleware.RespondJSON(ctx, http.StatusOK, resp, nil)
}
