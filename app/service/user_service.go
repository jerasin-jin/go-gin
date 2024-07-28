package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/Jerasin/app/response"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	GetPaginationUser(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.User)
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	// GetUser(c *gin.Context, user model.User, query map[interface{}]interface{}, field response.User) model.User
}

type UserServiceModel struct {
	BaseRepository repository.BaseRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func UserServiceInit(baseRepo repository.BaseRepositoryInterface, userRepo repository.UserRepositoryInterface) *UserServiceModel {
	return &UserServiceModel{
		BaseRepository: baseRepo,
		UserRepository: userRepo,
	}
}

func (u UserServiceModel) CreateUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	err := u.BaseRepository.Create(&request)
	if err != nil {
		pkg.PanicDatabaseException(err, c)
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))

}

func (u UserServiceModel) UpdateUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var err error
	log.Info("start to execute program update user data by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var request request.UpdateUserRequest
	if err = c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	var user model.User
	err = u.BaseRepository.FindOne(&user, "id = ?", userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	updateError := u.BaseRepository.Update(userID, &user, &request)

	if updateError != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.UpdateResponse()))
}

func (u UserServiceModel) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var user model.User
	err := u.BaseRepository.FindOne(&user, "id = ?", userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, user))
}

func (u UserServiceModel) GetPaginationUser(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.User) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data user")
	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)

	var users []model.User
	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      users,
	}

	data, err := u.BaseRepository.Pagination(paginationModel)
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	totalPage, err := u.BaseRepository.TotalPage(&users, pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	fmt.Println("count", totalPage)

	var res []response.User
	pkg.ModelDump(&res, data)
	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize))
}

func (u UserServiceModel) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))
	var user model.User
	err := u.BaseRepository.Delete(&user, userID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.DeleteResponse()))
}

// func (u UserServiceModel) GetUser(c *gin.Context, user model.User, query map[interface{}]interface{}, field response.User) model.User {
// 	defer pkg.PanicHandler(c)

// 	fields := structs.Map(field)

// 	log.Info("start to execute get data user")
// 	result, err := u.UserRepository.FindOneUser(user, query, fields)
// 	if err != nil {
// 		log.Error("Happened Error when try delete data user from DB. Error:", err)
// 		pkg.PanicException(constant.UnknownError)
// 	}

// 	return result
// }
