package token

import (
	"api_short_story/app/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaims struct {
	Id    int    `json:"id"`
	Name  string `json:"name`
	Email string `json:"email"`
	Role  int    `json:"int"`
	jwt.StandardClaims
}

var configs = config.ReadJsonConfig()
var KeyForAuthor = configs.KeyJWT.Author
var KeyForAdmin = configs.KeyJWT.Admin

func GenerateJWT(authorClaims JwtClaims) string {
	jsonData := jwt.NewWithClaims(jwt.SigningMethodHS256, authorClaims)
	if authorClaims.Role == 0 {
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
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return int(id)
}

var MiddlewareAdmin = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(KeyForAdmin),
})

var MiddlewareAuthor = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(KeyForAuthor),
})
