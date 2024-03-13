package controller

import (
	"fmt"
	"hacktiv-digitalent-golang/assignment/assignment-2/model"
	"hacktiv-digitalent-golang/assignment/assignment-2/repository"
	"hacktiv-digitalent-golang/assignment/assignment-2/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ordersController struct {
	ordersRepository repository.IOrdersRepository
}

func NewOrdersRepository(ordersRepository repository.IOrdersRepository) *ordersController {
	return &ordersController{
		ordersRepository: ordersRepository,
	}
}


// Crete New Orders
func (oc *ordersController) Create (ctx *gin.Context) {
	var newOrders model.Orders

	err := ctx.ShouldBindJSON(&newOrders)
	if err != nil {
		var res model.Response = model.Response{
			Success: false,
			Message: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	createOrders, err := oc.ordersRepository.Create(newOrders)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, createOrders, ""))
}

//  Get All Orders
func (oc *ordersController) GetAll (ctx *gin.Context) {

	orders, err := oc.ordersRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, orders, ""))
}

//Update Orders
func (oc *ordersController) Update(ctx *gin.Context) {
	order_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	var updatedOrder model.Orders
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	_, err = oc.ordersRepository.GetByID(order_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	updatedOrder.Orders_id = order_id

	updatedOrder, err = oc.ordersRepository.Update(updatedOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	var orders model.Orders
	fmt.Println(orders.Ordered_at)
	
	updatedOrder.Ordered_at = time.Now()

	ctx.JSON(http.StatusOK, util.CreateResponse(true, updatedOrder, ""))
}

// Delete Orders
func (oc *ordersController) Delete (ctx *gin.Context) {
	order_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	_, err = oc.ordersRepository.GetByID(order_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	err = oc.ordersRepository.Delete(order_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, "Data Deleted Successfully"))
}



