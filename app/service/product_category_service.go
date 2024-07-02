package service

import (
	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ProductCategoryServiceInterface interface{}

type ProductCategoryService struct {
	ProductCategoryRepo repository.ProductCategoryRepositoryInterface
}

func NewProductCategoryService(productCategoryRepo repository.ProductCategoryRepositoryInterface) ProductCategoryServiceInterface {
	return &ProductCategoryService{
		ProductCategoryRepo: productCategoryRepo,
	}
}

func (p ProductCategoryService) AddProductCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data product category")

	var request model.ProductCategoryRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	result, err := p.ProductCategoryRepo.FindOneProduct(request)

	if err != nil {

	}
}
