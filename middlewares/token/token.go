package token

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaims struct {
	Id    int    `json:"id"`
	Name  string `json:"name`
	Email string `json:"email"`
	jwt.StandardClaims
}

var KeyForAuthor = "forAuthor"
var KeyForAdmin = "forAdmin"

func GenerateJWT(authorClaims JwtClaims) string {
	jsonData := jwt.NewWithClaims(jwt.SigningMethodHS256, authorClaims)
	if authorClaims.Email == "admin@gmail.com" {
		token, err := jsonData.SignedString([]byte(KeyForAdmin))
		if err != nil {
			return err.Error()
		}
		return token
	} else {
		token, err := jsonData.SignedString([]byte(KeyForAuthor))
		if err != nil {
			return err.Error()
		}
		return token
	}
}

func GetAuthorIdFromJWT(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)
	fmt.Println(claims)
	return int(claims.Id)
}

var MiddlewareAdmin = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(KeyForAdmin),
})

var MiddlewareAuthor = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(KeyForAuthor),
})
