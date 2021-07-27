package Models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type Order struct {
	gorm.Model
	Slug string `gorm:"unique_index" json:"slug_id"`
	OrderStatus int `gorm:"default:0" json:"order_status"`
	Customer Customer `gorm:"foreignKey:CustomerId:" json:"customer"`
	CustomerID uint `gorm:"default:null" json:"customer_id"`
	Product Product `gorm:"foreignKey:CustomerId:" json:"product"`
	ProductID uint `gorm:"default:null" json:"product_id"`
	Quantity int `gorm:"default:1" json:"quantity"`
}

func (order *Order) GetOrderStatusAsString() string {
	switch order.OrderStatus {
	case 0:
		return "order placed"
	case 1:
		return "processed"
	case 2:
		return "failed"
	default:
		return "unknown"
	}
}

func (order *Order) BeforeCreate() (err error) {
	idString := strconv.Itoa(int(order.ID))
	idString = idString + strings.Repeat("0", 6- len(idString))
	order.Slug = "CST" + idString
	return
}

