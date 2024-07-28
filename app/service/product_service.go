package service

import (
	"net/http"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ProductServiceInterface interface {
	CreateProduct(c *gin.Context)
	GetPaginationProduct(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.Product)
}

type ProductServiceModel struct {
	ProductRepository repository.ProductRepositoryInterface
	BaseRepository    repository.BaseRepositoryInterface
}

func ProductServiceInit(productRepo repository.ProductRepositoryInterface, baseRepo repository.BaseRepositoryInterface) *ProductServiceModel {
	return &ProductServiceModel{
		ProductRepository: productRepo,
		BaseRepository:    baseRepo,
	}
}

func (p ProductServiceModel) CreateProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var request model.Product
	var err error

	// Validate Request Body
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.Error("error ShouldBindJSON", err)
		pkg.PanicException(constant.BadRequest)
	}

	// Validate Duplicated Data
	var products []model.Product
	err = p.BaseRepository.IsExits(&products, "name = ?", request.Name)
	if err != nil {
		pkg.PanicException(constant.Duplicated)
	}

	if len(products) > 0 {
		pkg.PanicException(constant.ValidateError)
	}

	err = p.BaseRepository.Create(&request)
	if err != nil {
		pkg.PanicException(constant.BadRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
}

func (p ProductServiceModel) GetPaginationProduct(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.Product) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get list product")

	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)
	// fields := structs.Map(field)

	// p.BaseRepository

	var products []model.Product
	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      products,
	}

	data, err := p.BaseRepository.Pagination(paginationModel)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	totalPage, err := p.BaseRepository.TotalPage(&products, pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	var res []response.Product

	pkg.ModelDump(&res, data)
	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize))
}
