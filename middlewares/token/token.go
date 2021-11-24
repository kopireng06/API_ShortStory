package token

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JwtAuthorClaims struct {
	Id    int    `json:"id"`
	Name  string `json:"name`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateAuthorJWT(authorClaims JwtAuthorClaims) string {
	jsonData := jwt.NewWithClaims(jwt.SigningMethodHS256, authorClaims)
	token, err := jsonData.SignedString([]byte("secret"))
	if err != nil {
		return err.Error()
	}
	return token
}

var MiddlewareToken = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
