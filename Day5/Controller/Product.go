package Controller

import (
	"Day5/Models"
	"Day5/Services"
	"Day5/Views"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product = Models.Product

func PostProduct (c *gin.Context) {
	var product Product
	c.BindJSON(&product)
	err := Services.PostProduct(&product)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(http.StatusOK, Views.ProductPostView(&product, 1))
	}
}

func UpdateProduct (c *gin.Context) {
	var product Product
	id := c.Params.ByName("id")
	err := Services.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithError(500, err)
	}
	if &product == nil {
		c.AbortWithError(404, errors.New("no records found to update"))
	}
	c.BindJSON(&product)
	err = Services.UpdateProduct(&product)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(http.StatusOK, Views.ProductPostView(&product, 1))
	}
}

func GetAllProducts (c *gin.Context) {
	var products []Product
	err := Services.GetAlLProducts(&products)
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(http.StatusOK, Views.GetProductsView(&products))
	}
}

func GetProductByID (c *gin.Context) {
	var product Product
	id := c.Params.ByName("id")
	err := Services.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithError(500, err)
	} else if &product == nil {
		c.AbortWithError(404, errors.New("no records found to update"))
	} else {
		c.JSON(http.StatusOK, Views.ProductPostView(&product, 2))
	}
}
