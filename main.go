package main

import (
	"fmt"

	"github.com/Jerasin/app/config"
	"github.com/Jerasin/app/router"
	docs "github.com/Jerasin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@host		localhost:3000
//	@BasePath	/api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	config.EnvConfig()

	PORT := config.GetEnv("PORT", "3000")
	baseModule := router.NewBaseModule()
	app := router.RouterInit(baseModule)
	docs.SwaggerInfo.BasePath = "/api"

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	appInfo := fmt.Sprintf("0.0.0.0:%s", PORT)
	fmt.Println("appInfo", appInfo)
	app.Run(appInfo) // listen and serve on 0.0.0.0:8080
}
