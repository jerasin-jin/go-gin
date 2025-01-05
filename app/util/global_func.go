package util

import (
	"fmt"
	"slices"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomError struct {
	message string
}

func (c CustomError) Error() string {
	return c.message
}

func FindElementByCondition[T any](slice []T, condition func(T) bool) (*T, error) {
	index := slices.IndexFunc(slice, condition)

	if index != -1 {
		return &slice[index], nil
	} else {
		pkg.PanicException(constant.BadRequest)
		return nil, CustomError{
			message: "FindElementByCondition Error",
		}
	}

}

func GetUserId(c *gin.Context) (any, error) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		pkg.PanicException(constant.Unauthorized)
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := pkg.NewAuthService().ValidateToken(tokenString)

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println("claims", claims)
		return claims["username"], nil
	} else {
		return nil, err
	}
}
