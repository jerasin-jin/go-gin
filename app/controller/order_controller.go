package controller

import (
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	svc service.OrderServiceInterface
}

type OrderControllerInterface interface {
	CreateOrder(c *gin.Context)
}

func OrderControllerInit(orderSvc service.OrderServiceInterface) *OrderController {
	return &OrderController{
		svc: orderSvc,
	}
}

// @Summary CreateOrder
// @Schemes
// @Description CreateOrder
// @Tags Order
//
// @Param request body request.OrderRequest true "query params"
//
//	@Success		200	{object}	model.Order
//
// @Security Bearer
//
// @Router /orders [post]
func (o OrderController) CreateOrder(c *gin.Context) {
	o.svc.CreateOrder(c)
}
