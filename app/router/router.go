package router

import (
	"github.com/Jerasin/app/middleware"
	"github.com/Jerasin/app/module"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type BaseModuleInit struct {
	UserModule *module.UserModule
	AuthModule *module.AuthModule
}

func NewBaseModule() BaseModuleInit {
	userInit := module.UserModuleInit()
	authInit := module.AuthModuleInit()
	return BaseModuleInit{
		UserModule: userInit,
		AuthModule: authInit,
	}
}

func RouterInit(init BaseModuleInit) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	api := router.Group("/api")

	user := api.Group("/user")

	user.Use(middleware.AuthorizeJwt())
	user.GET("", init.UserModule.UserCtrl.GetAllUserData)
	user.POST("", init.UserModule.UserCtrl.AddUserData)
	user.GET("/:userID", init.UserModule.UserCtrl.GetUserById)
	user.PUT("/:userID", init.UserModule.UserCtrl.UpdateUserData)
	user.DELETE("/:userID", init.UserModule.UserCtrl.DeleteUser)

	auth := api.Group("/auth")

	auth.POST("/login", init.AuthModule.AuthCtrl.Login)

	return router
}
