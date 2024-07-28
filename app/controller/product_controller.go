package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type ProductControllerInterface interface {
	CreateProduct(c *gin.Context)
	GetAllProducts(c *gin.Context)
}

type ProductController struct {
	svc service.ProductServiceInterface
}

func ProductControllerInit(productSvc service.ProductServiceInterface) *ProductController {
	return &ProductController{
		svc: productSvc,
	}
}

// @Summary Create Product
// @Schemes
// @Description Create Product
// @Tags Product
//
// @Param request body request.Product true "query params"
//
//	@Success		200	{object}	response.CreateDataResponse
//
// @Security Bearer
//
// @Router /products [post]
func (p ProductController) CreateProduct(c *gin.Context) {
	p.svc.CreateProduct(c)
}

// @Summary Get List Products
// @Schemes
// @Description Get List Products
// @Tags Product
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.ProductPagination
//
// @Security Bearer
//
// @Router /products [get]
func (p ProductController) GetAllProducts(c *gin.Context) {
	query := CreatePagination(c)
	product := response.Product{}

	p.svc.GetPaginationProduct(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, product)
}
