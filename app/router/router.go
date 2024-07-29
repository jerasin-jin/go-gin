package router

import (
	"github.com/Jerasin/app/middleware"
	"github.com/Jerasin/app/module"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type BaseModuleInit struct {
	UserModule            *module.UserModule
	AuthModule            *module.AuthModule
	ProductCategoryModule *module.ProductCategoryModule
	ProductModule         *module.ProductModule
}

func NewBaseModule() BaseModuleInit {
	userInit := module.UserModuleInit()
	authInit := module.AuthModuleInit()
	productCategoryInit := module.ProductCategoryModuleInit()
	productInit := module.ProductModuleInit()

	return BaseModuleInit{
		UserModule:            userInit,
		AuthModule:            authInit,
		ProductCategoryModule: productCategoryInit,
		ProductModule:         productInit,
	}
}

func RouterInit(init BaseModuleInit) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	api := router.Group("/api")

	user := api.Group("/users")
	user.Use(middleware.AuthorizeJwt())
	user.GET("", init.UserModule.UserCtrl.GetAllUsers)
	user.POST("", init.UserModule.UserCtrl.CreateUser)
	user.GET("/:userID", init.UserModule.UserCtrl.GetUserById)
	user.PUT("/:userID", init.UserModule.UserCtrl.UpdateUserData)
	user.DELETE("/:userID", init.UserModule.UserCtrl.DeleteUser)

	auth := api.Group("/auth")

	auth.POST("/register", init.AuthModule.AuthCtrl.Register)
	auth.POST("/login", init.AuthModule.AuthCtrl.Login)
	auth.POST("/refresh/token", init.AuthModule.AuthCtrl.RefreshToken)

	product := api.Group("/products")
	product.POST("", init.ProductModule.ProductCtrl.CreateProduct)
	product.GET("", init.ProductModule.ProductCtrl.GetAllProducts)
	product.GET("/:productID", init.ProductModule.ProductCtrl.GetProductById)
	product.PUT("/:productID", init.ProductModule.ProductCtrl.UpdateProductData)
	product.DELETE("/:productID", init.ProductModule.ProductCtrl.DeleteProduct)

	productCategory := product.Group("/categories")
	productCategory.Use(middleware.AuthorizeJwt())
	productCategory.POST("", init.ProductCategoryModule.ProductCategoryCtrl.CreateProductCategory)
	productCategory.GET("", init.ProductCategoryModule.ProductCategoryCtrl.GetListProductCategory)
	productCategory.GET("/:productCategoryID", init.ProductCategoryModule.ProductCategoryCtrl.GetProductCategoryById)
	productCategory.PUT("/:productCategoryID", init.ProductCategoryModule.ProductCategoryCtrl.UpdateProductCategoryData)
	return router
}
