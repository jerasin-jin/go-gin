// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package module

import (
	"github.com/Jerasin/app/controller"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/service"
	"github.com/Jerasin/app/util"
	"github.com/google/wire"
)

// Injectors from auth_injector.go:

func AuthModuleInit() *AuthModule {
	gormDB := util.InitDbClient()
	baseRepository := repository.BaseRepositoryInit(gormDB)
	userRepository := repository.UserRepositoryInit(baseRepository)
	userServiceModel := service.UserServiceInit(baseRepository, userRepository)
	authServiceModel := service.AuthServiceInit(baseRepository, userRepository, userServiceModel)
	authController := controller.AuthControllerInit(authServiceModel)
	authModule := NewAuthModule(authController, authServiceModel, userRepository, userServiceModel)
	return authModule
}

// Injectors from order_injector.go:

func OrderModuleInit() *OrderModule {
	gormDB := util.InitDbClient()
	baseRepository := repository.BaseRepositoryInit(gormDB)
	orderServiceModel := service.OrderServiceInit(baseRepository)
	orderController := controller.OrderControllerInit(orderServiceModel)
	orderModule := NewOrderModule(orderServiceModel, orderController)
	return orderModule
}

// Injectors from permission_info_injector.go:

func PermissionInfoModuleInit() *PermissionInfoModule {
	gormDB := util.InitDbClient()
	baseRepository := repository.BaseRepositoryInit(gormDB)
	permissionInfoServiceModel := service.PermissionInfoServiceInit(baseRepository)
	permissionInfoController := controller.PermissionInfoControllerInit(permissionInfoServiceModel)
	permissionInfoModule := NewPermissionInfoModule(permissionInfoServiceModel, permissionInfoController)
	return permissionInfoModule
}

// Injectors from product_category_injector.go:

func ProductCategoryModuleInit() *ProductCategoryModule {
	gormDB := util.InitDbClient()
	baseRepository := repository.BaseRepositoryInit(gormDB)
	productCategoryRepository := repository.ProductCategoryRepositoryInit(baseRepository)
	productCategoryServiceModel := service.ProductCategoryServiceInit(baseRepository, productCategoryRepository)
	productCategoryController := controller.ProductCategoryControllerInit(productCategoryServiceModel)
	productCategoryModule := NewProductCategoryModule(productCategoryController, productCategoryServiceModel, productCategoryRepository)
	return productCategoryModule
}

// Injectors from product_injector.go:

func ProductModuleInit() *ProductModule {
	gormDB := util.InitDbClient()
	baseRepository := repository.BaseRepositoryInit(gormDB)
	productRepository := repository.ProductRepositoryInit(baseRepository)
	productServiceModel := service.ProductServiceInit(productRepository, baseRepository)
	productController := controller.ProductControllerInit(productServiceModel)
	productModule := NewProductModule(productRepository, productServiceModel, productController)
	return productModule
}

// Injectors from role_info_injector.go:

func RoleInfoModuleInit() *RoleInfoModule {
	gormDB := util.InitDbClient()
	baseRepository := repository.BaseRepositoryInit(gormDB)
	roleInfoServiceModel := service.RoleInfoServiceInit(baseRepository)
	roleInfoController := controller.RoleInfoControllerInit(roleInfoServiceModel)
	roleInfoModule := NewRoleInfoModule(roleInfoServiceModel, roleInfoController)
	return roleInfoModule
}

// Injectors from user_injector.go:

func UserModuleInit() *UserModule {
	gormDB := util.InitDbClient()
	baseRepository := repository.BaseRepositoryInit(gormDB)
	userRepository := repository.UserRepositoryInit(baseRepository)
	userServiceModel := service.UserServiceInit(baseRepository, userRepository)
	userController := controller.UserControllerInit(userServiceModel)
	userModule := NewUserModule(userRepository, userServiceModel, userController)
	return userModule
}

// Injectors from wallet_injector.go:

func WalletModuleInit() *WalletModule {
	gormDB := util.InitDbClient()
	baseRepository := repository.BaseRepositoryInit(gormDB)
	walletServiceModel := service.WalletServiceInit(baseRepository)
	walletController := controller.WalletControllerInit(walletServiceModel)
	walletModule := NewWalletModule(walletServiceModel, walletController)
	return walletModule
}

// auth_injector.go:

var authSvcSet = wire.NewSet(service.AuthServiceInit, wire.Bind(new(service.AuthServiceInterface), new(*service.AuthServiceModel)))

var authCtrlSet = wire.NewSet(controller.AuthControllerInit, wire.Bind(new(controller.AuthControllerInterface), new(*controller.AuthController)))

type AuthModule struct {
	AuthCtrl controller.AuthControllerInterface
	AuthSvc  service.AuthServiceInterface
	UserRepo repository.UserRepositoryInterface
	UserSvc  service.UserServiceInterface
}

func NewAuthModule(
	authCtrl controller.AuthControllerInterface,
	authSvc service.AuthServiceInterface,
	userRepo repository.UserRepositoryInterface,
	userSvc service.UserServiceInterface,
) *AuthModule {
	return &AuthModule{
		AuthSvc:  authSvc,
		AuthCtrl: authCtrl,
		UserRepo: userRepo,
		UserSvc:  userSvc,
	}
}

// order_injector.go:

var OrderSvcSet = wire.NewSet(service.OrderServiceInit, wire.Bind(new(service.OrderServiceInterface), new(*service.OrderServiceModel)))

var OrderCtrlSet = wire.NewSet(controller.OrderControllerInit, wire.Bind(new(controller.OrderControllerInterface), new(*controller.OrderController)))

type OrderModule struct {
	// OrderRepo repository.OrderRepositoryInterface
	OrderSvc  service.OrderServiceInterface
	OrderCtrl controller.OrderControllerInterface
}

func NewOrderModule(

	OrderService service.OrderServiceInterface,
	OrderCtrl controller.OrderControllerInterface,
) *OrderModule {
	return &OrderModule{

		OrderSvc:  OrderService,
		OrderCtrl: OrderCtrl,
	}
}

// permission_info_injector.go:

var PermissionInfoSvcSet = wire.NewSet(service.PermissionInfoServiceInit, wire.Bind(new(service.PermissionInfoServiceInterface), new(*service.PermissionInfoServiceModel)))

var PermissionInfoCtrlSet = wire.NewSet(controller.PermissionInfoControllerInit, wire.Bind(new(controller.PermissionInfoControllerInterface), new(*controller.PermissionInfoController)))

type PermissionInfoModule struct {
	PermissionInfoSvc  service.PermissionInfoServiceInterface
	PermissionInfoCtrl controller.PermissionInfoControllerInterface
}

func NewPermissionInfoModule(
	PermissionInfoService service.PermissionInfoServiceInterface,
	PermissionInfoCtrl controller.PermissionInfoControllerInterface,
) *PermissionInfoModule {
	return &PermissionInfoModule{
		PermissionInfoSvc:  PermissionInfoService,
		PermissionInfoCtrl: PermissionInfoCtrl,
	}
}

// product_category_injector.go:

var productCategorySvcSet = wire.NewSet(service.ProductCategoryServiceInit, wire.Bind(new(service.ProductCategoryServiceInterface), new(*service.ProductCategoryServiceModel)))

var productCategoryCtrlSet = wire.NewSet(controller.ProductCategoryControllerInit, wire.Bind(new(controller.ProductCategoryControllerInterface), new(*controller.ProductCategoryController)))

var productCategoryRepoSet = wire.NewSet(repository.ProductCategoryRepositoryInit, wire.Bind(new(repository.ProductCategoryRepositoryInterface), new(*repository.ProductCategoryRepository)))

type ProductCategoryModule struct {
	ProductCategoryCtrl controller.ProductCategoryControllerInterface
	ProductCategorySvc  service.ProductCategoryServiceInterface
	ProductCategoryRepo repository.ProductCategoryRepositoryInterface
}

func NewProductCategoryModule(
	productCategoryCtrl controller.ProductCategoryControllerInterface,
	productCategorySvc service.ProductCategoryServiceInterface,
	productCategoryRepo repository.ProductCategoryRepositoryInterface,
) *ProductCategoryModule {
	return &ProductCategoryModule{
		ProductCategoryCtrl: productCategoryCtrl,
		ProductCategorySvc:  productCategorySvc,
		ProductCategoryRepo: productCategoryRepo,
	}
}

// product_injector.go:

var productSvcSet = wire.NewSet(service.ProductServiceInit, wire.Bind(new(service.ProductServiceInterface), new(*service.ProductServiceModel)))

var productCtrlSet = wire.NewSet(controller.ProductControllerInit, wire.Bind(new(controller.ProductControllerInterface), new(*controller.ProductController)))

var productRepoSet = wire.NewSet(repository.ProductRepositoryInit, wire.Bind(new(repository.ProductRepositoryInterface), new(*repository.ProductRepository)))

type ProductModule struct {
	ProductRepo repository.ProductRepositoryInterface
	ProductSvc  service.ProductServiceInterface
	ProductCtrl controller.ProductControllerInterface
}

func NewProductModule(productRepo repository.ProductRepositoryInterface,
	productService service.ProductServiceInterface,
	productCtrl controller.ProductControllerInterface,
) *ProductModule {
	return &ProductModule{
		ProductRepo: productRepo,
		ProductSvc:  productService,
		ProductCtrl: productCtrl,
	}
}

// role_info_injector.go:

var RoleInfoSvcSet = wire.NewSet(service.RoleInfoServiceInit, wire.Bind(new(service.RoleInfoServiceInterface), new(*service.RoleInfoServiceModel)))

var RoleInfoCtrlSet = wire.NewSet(controller.RoleInfoControllerInit, wire.Bind(new(controller.RoleInfoControllerInterface), new(*controller.RoleInfoController)))

type RoleInfoModule struct {
	RoleInfoSvc  service.RoleInfoServiceInterface
	RoleInfoCtrl controller.RoleInfoControllerInterface
}

func NewRoleInfoModule(
	RoleInfoService service.RoleInfoServiceInterface,
	RoleInfoCtrl controller.RoleInfoControllerInterface,
) *RoleInfoModule {
	return &RoleInfoModule{
		RoleInfoSvc:  RoleInfoService,
		RoleInfoCtrl: RoleInfoCtrl,
	}
}

// user_injector.go:

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserServiceInterface), new(*service.UserServiceModel)))

var userCtrlSet = wire.NewSet(controller.UserControllerInit, wire.Bind(new(controller.UserControllerInterface), new(*controller.UserController)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit, wire.Bind(new(repository.UserRepositoryInterface), new(*repository.UserRepository)))

type UserModule struct {
	UserRepo repository.UserRepositoryInterface
	UserSvc  service.UserServiceInterface
	UserCtrl controller.UserControllerInterface
}

func NewUserModule(userRepo repository.UserRepositoryInterface,
	userService service.UserServiceInterface,
	userCtrl controller.UserControllerInterface,
) *UserModule {
	return &UserModule{
		UserRepo: userRepo,
		UserSvc:  userService,
		UserCtrl: userCtrl,
	}
}

// wallet_injector.go:

var WalletSvcSet = wire.NewSet(service.WalletServiceInit, wire.Bind(new(service.WalletServiceInterface), new(*service.WalletServiceModel)))

var WalletCtrlSet = wire.NewSet(controller.WalletControllerInit, wire.Bind(new(controller.WalletControllerInterface), new(*controller.WalletController)))

type WalletModule struct {
	WalletSvc  service.WalletServiceInterface
	WalletCtrl controller.WalletControllerInterface
}

func NewWalletModule(
	WalletService service.WalletServiceInterface,
	WalletCtrl controller.WalletControllerInterface,
) *WalletModule {
	return &WalletModule{
		WalletSvc:  WalletService,
		WalletCtrl: WalletCtrl,
	}
}
