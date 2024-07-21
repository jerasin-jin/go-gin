package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserController struct {
	svc service.UserServiceInterface
}

func UserControllerInit(userService service.UserServiceInterface) *UserController {
	return &UserController{
		svc: userService,
	}
}

// @Summary Get List users
// @Schemes
// @Description Get List users
// @Tags User
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.UserPagination
//
// @Security Bearer
//
// @Router /user [get]
func (u UserController) GetAllUserData(c *gin.Context) {
	query := CreatePagination(c)
	user := response.User{}

	u.svc.GetPaginationUser(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, user)
}

// @Summary Create user
// @Schemes
// @Description Create user
// @Tags User
//
// @Param request body request.UserRequest true "query params"
//
//	@Success		200	{object}	model.User
//
// @Security Bearer
//
// @Router /user [post]
func (u UserController) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

// @Summary Get user By Id
// @Schemes
// @Description Get user By Id
// @Tags User
// @Param userID  path int true "User ID"
//
//	@Success		200	{object}	response.User
//
// @Security Bearer
//
// @Router /user/{userID} [get]
func (u UserController) GetUserById(c *gin.Context) {
	u.svc.GetUserById(c)
}

// @Summary Update user By Id
// @Schemes
// @Description Update user By Id
// @Tags User
// @Param userID  path int true "User ID"
// @Param request body request.UserRequest true "query params"
//
//	@Success		200	{object}	model.User
//
// @Security Bearer
//
// @Router /user/{userID} [put]
func (u UserController) UpdateUserData(c *gin.Context) {
	u.svc.UpdateUserData(c)
}

// @Summary Update user By Id
// @Schemes
// @Description Update user By Id
// @Tags User
// @Param userID  path int true "User ID"
//
//	@Success		200	{object}	model.User
//
// @Security Bearer
//
// @Router /user/{userID} [delete]
func (u UserController) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}
