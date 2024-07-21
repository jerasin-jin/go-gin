package service

import (
	"fmt"
	"net/http"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/response"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ProductCategoryServiceInterface interface {
	AddProductCategory(c *gin.Context)
	GetPaginationProductCategory(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.ProductCategory)
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

func (p ProductCategoryServiceModel) GetPaginationProductCategory(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.ProductCategory) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get list product category")

	log.Info("start to execute get all data user")
	offset := (page - 1) * pageSize
	limit := pageSize
	fields := structs.Map(field)
	fmt.Println("query", search)
	fmt.Println("fields", fields)

	data, err := p.ProductCategoryRepo.PaginationProductCategory(limit, offset, search, sortField, sortValue, fields)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	totalPage, err := p.ProductCategoryRepo.TotalPage(pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	fmt.Println("count", totalPage)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}
