package common

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/izhujiang/appengine/module/auth"
	"github.com/izhujiang/appengine/module/core"
)

func ValidateJWT(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func CurrentUser(c *gin.Context) (auth.User, error) {
	// err := ValidateJWT(c)
	// if err != nil {
	// 	return auth.User{}, err
	// }

	token, _ := getToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	// fmt.Println("claims: ", claims)
	// fmt.Println("userId: ", claims["id"])
	userId, err := core.StringToUUID(claims["id"].(string))
	if err != nil {
		return auth.User{}, err
	}

	user, err := auth.FindUserById(userId)
	return user, nil
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	// Exacts token from Request.Header
	var tokenString string
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		tokenString = splitToken[1]
	}

	token, err := auth.ParseJWT(tokenString)
	return token, err
}
