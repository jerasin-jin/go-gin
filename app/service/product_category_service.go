package service

import (
	"net/http"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ProductCategoryServiceInterface interface {
	AddProductCategory(c *gin.Context)
	GetListProductCategory(c *gin.Context)
}

type ProductCategoryServiceModel struct {
	ProductCategoryRepo repository.ProductCategoryRepositoryInterface
}

func ProductCategoryServiceInit(productCategoryRepo repository.ProductCategoryRepositoryInterface) *ProductCategoryServiceModel {
	return &ProductCategoryServiceModel{
		ProductCategoryRepo: productCategoryRepo,
	}
}

func (p ProductCategoryServiceModel) AddProductCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data product category")

	var request model.ProductCategory
	if err := c.ShouldBind(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := p.ProductCategoryRepo.Save(&request)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p ProductCategoryServiceModel) GetListProductCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get list product category")

	// var request model.ProductCategory
	// if err := c.ShouldBind(&request); err != nil {
	// 	log.Error("Happened error when mapping request from FE. Error", err)
	// 	pkg.PanicException(constant.InvalidRequest)
	// }

	data, err := p.ProductCategoryRepo.FindAllProductCategory()

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}
