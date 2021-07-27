package Controller

import (
	"Day5/Models"
	"Day5/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Customer = Models.Customer
type Order = Models.Order

func CreateCustomer(c *gin.Context) {
	var customer Customer
	c.BindJSON(&customer)
	err := Services.CreateCustomer(&customer)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func GetOrdersByCustomer(c *gin.Context) {
	var orders []Order
	var customerID uint
	c.BindJSON(customerID)
	err := Services.GetOrdersByCustomer(&orders, customerID)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}


