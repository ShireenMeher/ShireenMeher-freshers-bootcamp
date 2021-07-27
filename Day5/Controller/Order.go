package Controller

import (
	"Day5/Services"
	"Day5/Views"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewOrder(c *gin.Context) {
	var order Order
	c.BindJSON(&order)
	feasibility, err := Services.CheckFeasibility(order.CustomerID)
	if err !=nil {
		c.AbortWithError(500, err)
	}
	if !feasibility {
		c.AbortWithError(404, errors.New("please wait for 5min and book new order"))
	}
	err = Services.NewOrder(&order)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(http.StatusOK, Views.GetOrderView(&order))
	}
}

func UpdateOrder(c *gin.Context) {
	var order Order
	id := c.Params.ByName("id")
	err := Services.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithError(500, err)
	}
	if &order == nil {
		c.AbortWithError(404, errors.New("no records found to update"))
	}
	c.BindJSON(&order)
	err = Services.UpdateOrder(&order)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(http.StatusOK, Views.GetOrderView(&order))
	}
}


