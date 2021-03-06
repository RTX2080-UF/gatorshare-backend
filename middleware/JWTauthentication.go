package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var signingKey = GetEnv("ACCESS_SECRET", "Gatorshare", true)

func CreateToken(userid uint) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func TokenValid(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractTokenID(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}

func ExtractToken(ctx *gin.Context) string {
	keys := ctx.Request.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.Request.Header["Authorization"]
	if len(bearerToken) > 0 && len(strings.Split(bearerToken[0], " ")) == 2 {
		return strings.Split(bearerToken[0], " ")[1]
	}

  	return ""
}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}

func GetUidFromToken(ctx *gin.Context) uint {
	token := ExtractToken(ctx)
	errCustom := errors.New("invalid token provided")

	if (token != "") {
		err := TokenValid(token)
		if (err != nil) {
			RespondJSON(ctx, http.StatusForbidden, errCustom, err)
			return 0
		}
	} else {
		RespondJSON(ctx, http.StatusForbidden, errCustom, nil)
		return 0
	}
	
	uid, err := ExtractTokenID(token)
	if err != nil {
		RespondJSON(ctx, http.StatusForbidden, errCustom, err)
		return 0
	}

	return uid
}

func RefreshToken(token string) string {
	err := TokenValid(token)
	if (err != nil) {
		return ""
	}

	uid, err := ExtractTokenID(token)
	if (err != nil) {
		return ""
	}

	newtoken, err := CreateToken(uid)
	if (err != nil) {
		return ""
	}

	return newtoken
}