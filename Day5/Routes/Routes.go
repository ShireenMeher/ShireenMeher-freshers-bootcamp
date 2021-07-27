package Routes

import (
	"Day5/Controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	group := r.Group("")
	{
		group.GET("products", Controller.GetAllProducts)
		group.POST("products", Controller.GetAllProducts)
		group.GET("products/:id", Controller.GetProductByID)
		group.PATCH("products/:id", Controller.UpdateProduct)
		group.POST("orders", Controller.NewOrder)
		group.PATCH("orders", Controller.UpdateOrder)
		group.POST("customers", Controller.CreateCustomer)
		group.POST("orders/customers/:id", Controller.GetOrdersByCustomer)
	}
	return r
}
