package middleware

import (
  "github.com/golang-jwt/jwt"
  "time"
)

func CreateToken(userid uint, key string) (string, error) {
  var err error

  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userid
  atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

  token, err := at.SignedString([]byte(key))
  if err != nil {
     return "", err
  }

  return token, nil
}