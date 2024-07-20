package controller

import (
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type ProductCategoryControllerInterface interface {
	GetListProductCategory(c *gin.Context)
	AddProductCategory(c *gin.Context)
	// GetUserById(c *gin.Context)
	// UpdateUserData(c *gin.Context)
	// DeleteUser(c *gin.Context)
}

type ProductCategoryController struct {
	svc service.ProductCategoryServiceInterface
}

func ProductCategoryControllerInit(productCategorySvc service.ProductCategoryServiceInterface) *ProductCategoryController {
	return &ProductCategoryController{
		svc: productCategorySvc,
	}
}

// @Summary Create product category
// @Schemes
// @Description Create Product Category
// @Tags Product Category
//
// @Param request body request.ProductCategoryRequest true "query params"
//
//	@Success		200	{object}	model.ProductCategory
//
// @Security Bearer
//
// @Router /product/category [post]
func (p ProductCategoryController) AddProductCategory(c *gin.Context) {
	p.svc.AddProductCategory(c)
}

// @Summary Get List product category
// @Schemes
// @Description Get List Product Category
// @Tags Product Category
//
//	@Success		200	{object}	model.ProductCategory
//
// @Security Bearer
//
// @Router /product/category [get]
func (p ProductCategoryController) GetListProductCategory(c *gin.Context) {
	p.svc.GetListProductCategory(c)
}
