package service

import (
	"fmt"
	"net/http"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderServiceModel struct {
	BaseRepository repository.BaseRepositoryInterface
}

type OrderServiceInterface interface {
	CreateOrder(c *gin.Context)
}

func OrderServiceInit(baseRepo repository.BaseRepositoryInterface) *OrderServiceModel {
	return &OrderServiceModel{
		BaseRepository: baseRepo,
	}
}

func (o OrderServiceModel) calDetailOrder(tx *gorm.DB, request []request.OrderItem, orderID uint) {
	fmt.Printf("Orders = %+v\n", request)
	var err error
	for _, value := range request {
		order := model.OrderDetail{
			ProductID:         uint(value.Id),
			Name:              value.Name,
			Description:       value.Description,
			Price:             value.Price,
			Amount:            value.Amount,
			ProductCategoryID: uint(value.ProductCategoryId),
			OrderID:           orderID,
		}

		product := model.Product{}
		err = o.BaseRepository.FindOne(tx, &product, "id = ?", value.Id)
		if err != nil {
			log.Error("error ShouldBindJSON", err)
			pkg.PanicException(constant.DataNotFound)
		}

		if product.Amount < value.Amount {
			log.Error("error ShouldBindJSON", err)
			pkg.PanicException(constant.DataNotFound)
		}

		updateProduct := map[string]interface{}{
			"Amount": product.Amount - value.Amount,
		}

		err = o.BaseRepository.Update(tx, value.Id, &product, &updateProduct)
		if err != nil {
			log.Error("error ShouldBindJSON", err)
			pkg.PanicException(constant.DataNotFound)
		}

		fmt.Printf("Updating product with ID %d: %+v\n", value.Id, &updateProduct)

		o.BaseRepository.Create(tx, &order)
	}
}

func sumTotal(request []request.OrderItem) (int, int) {
	var amountTotal int
	var priceTotal int
	for _, value := range request {
		amountTotal += value.Amount
		priceTotal += int(value.Price)
	}

	return amountTotal, priceTotal
}

func (o OrderServiceModel) CreateOrder(c *gin.Context) {
	defer pkg.PanicHandler(c)

	o.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var request request.OrderRequest
		var err error

		// Validate Request Body
		err = c.ShouldBindJSON(&request)
		if err != nil {
			log.Error("error ShouldBindJSON", err)
			pkg.PanicException(constant.BadRequest)
		}

		// fmt.Printf("OrderRequest = %+v\n", request)
		// fmt.Printf("%+v\n", request)

		amountTotal, priceTotal := sumTotal(request.Orders)
		order := model.Order{
			TotalPrice:  priceTotal,
			TotalAmount: amountTotal,
		}

		fmt.Printf("order = %+v\n", order)

		err = o.BaseRepository.Create(tx, &order)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		o.calDetailOrder(tx, request.Orders, order.ID)

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
		return nil
	})

}
