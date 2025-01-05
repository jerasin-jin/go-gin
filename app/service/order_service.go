package service

import (
	"fmt"
	"net/http"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/Jerasin/app/util"
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

// func (o OrderServiceModel) calDetailOrder(tx *gorm.DB, request []request.OrderItem) ([]model.OrderDetail, int, float64) {
// 	fmt.Printf("Orders = %+v\n", request)
// 	var err error
// 	amountTotal := 0
// 	priceTotal := 0.0
// 	var orderDetailList []model.OrderDetail

// 	for _, value := range request {

// 		product := model.Product{}
// 		err = o.BaseRepository.FindOne(tx, &product, "id = ?", value.Id)
// 		if err != nil {
// 			log.Error("error ShouldBindJSON", err)
// 			pkg.PanicException(constant.DataNotFound)
// 		}

// 		order := model.OrderDetail{
// 			ProductID:         uint(value.Id),
// 			Name:              product.Name,
// 			Description:       product.Description,
// 			Price:             product.Price,
// 			Amount:            value.Amount,
// 			ProductCategoryID: uint(product.ProductCategoryID),
// 			OrderID:           orderID,
// 		}

// 		if product.Amount < value.Amount {
// 			log.Error("error ShouldBindJSON", err)
// 			pkg.PanicException(constant.DataNotFound)
// 		}

// 		updateProduct := map[string]interface{}{
// 			"Amount": product.Amount - value.Amount,
// 		}

// 		err = o.BaseRepository.Update(tx, value.Id, &product, &updateProduct)
// 		if err != nil {
// 			log.Error("error ShouldBindJSON", err)
// 			pkg.PanicException(constant.DataNotFound)
// 		}

// 		fmt.Printf("Updating product with ID %d: %+v\n", value.Id, &updateProduct)

// 		amountTotal += value.Amount
// 		priceTotal += product.Price
// 		// o.BaseRepository.Save(tx, &order)

// 		orderDetailList := append(orderDetailList, order)
// 	}

// 	return orderDetailList, amountTotal, priceTotal
// }

// func (o OrderServiceModel) getProduct(tx *gorm.DB, productID []uint) (model.Product, error) {
// 	product := model.Product{}
// 	err := o.BaseRepository.Find(tx, &product, "id = ?", productID)

// 	if err != nil {
// 		log.Error("error ShouldBindJSON", err)
// 		return product, err
// 	}

// 	return product, nil
// }

// goroutine ใช้ใน transaction Database ไม่ได้
func (o OrderServiceModel) CreateOrder(c *gin.Context) {
	defer pkg.PanicHandler(c)

	o.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var body request.OrderRequest
		var err error
		var productIDS []uint
		var products []model.Product
		var orderDetails []model.OrderDetail
		var totalPrice float64
		var totalAmount int

		// Validate Request Body
		err = c.ShouldBindJSON(&body)
		if err != nil {
			log.Error("error ShouldBindJSON", err)
			pkg.PanicException(constant.BadRequest)
		}

		for _, value := range body.Orders {
			productIDS = append(productIDS, uint(value.ProductId))
		}

		log.Debugf("productIDS = %+v\n", productIDS)

		err = o.BaseRepository.Find(tx, &products, "id IN ?", productIDS)
		if err != nil {
			fmt.Println("err", err)
			pkg.PanicException(constant.BadRequest)
		}

		// fmt.Printf("products = %+v type = %T \n", products, products)
		// fmt.Println(len(products))
		// fmt.Println(len(productIDS))

		if len(products) != len(productIDS) || len(products) == 0 {
			pkg.PanicException(constant.DataNotFound)
		}

		for index, product := range products {
			// fmt.Printf("product = %+v\n", product)

			if product.Amount <= 0 {
				pkg.PanicException(constant.BadRequest)
			}

			item, err := util.FindElementByCondition(body.Orders, func(o request.OrderItem) bool {
				return o.ProductId == int(product.ID)
			})
			if err != nil {
				pkg.PanicException(constant.BadRequest)
			}

			orderDetail := model.OrderDetail{
				ProductID: product.ID,
				Price:     product.Price,
				Amount:    item.Amount,
			}

			orderDetails = append(orderDetails, orderDetail)
			totalPrice += product.Price * float64(item.Amount)
			totalAmount += item.Amount

			// fmt.Println("item", item)
			// fmt.Println("*item", *item)

			// fmt.Println("item.Amount", item.Amount)
			// fmt.Printf("FindElementByCondition = %+v type = %T \n", *item, *item)

			products[index].Amount -= item.Amount
			// fmt.Printf("product = %+v\n", product)
		}

		fmt.Printf("products 2 = %+v\n", products)

		username, err := util.GetUserId(c)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}
		fmt.Println("username", username)
		var user model.User
		o.BaseRepository.FindOne(tx, &user, "username = ?", username)

		order := model.Order{
			TotalPrice:  totalPrice,
			TotalAmount: totalAmount,
			CreatedBy:   user.ID,
		}
		err = o.BaseRepository.Save(tx, &order)
		fmt.Println(err)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}
		fmt.Println("create Order Success")

		for i, value := range orderDetails {
			fmt.Println("value", value)
			orderDetails[i].OrderID = order.ID
		}

		err = o.BaseRepository.Save(tx, &orderDetails)
		fmt.Println(err)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		o.BaseRepository.Save(tx, &products)

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
		return nil
	})

}
