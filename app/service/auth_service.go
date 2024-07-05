package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	Login(c *gin.Context)
}

type AuthServiceModel struct {
	UserRepository repository.UserRepositoryInterface
}

func AuthServiceInit(userRepo repository.UserRepositoryInterface) *AuthServiceModel {
	return &AuthServiceModel{
		UserRepository: userRepo,
	}
}

func (authSvc AuthServiceModel) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request request.LoginRequest
	var user model.User

	fmt.Println("request", request)

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.CustomPanicException(constant.InvalidRequest, err.Error())
	}

	query := make(map[interface{}]interface{})
	query["username"] = request.Username

	fields := make(map[string]interface{})

	result, err := authSvc.UserRepository.FindOneUser(user, query, fields)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.PanicException(constant.DataNotFound)
		}

		pkg.PanicException(constant.InvalidRequest)
	}

	fmt.Println("result", result)

	isError := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))

	fmt.Println("isError", isError)

	if isError != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	// var res response.User
	// copier.Copy(&res, &result)

	jwt := pkg.NewAuthService()

	token := jwt.GenerateToken(result.Username)

	var response = make(map[string]interface{})
	response["token"] = token
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
