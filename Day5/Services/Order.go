package Services

import (
	"Day5/Config"
	"Day5/Models"
	"strconv"
	"sync"
)

var mutex sync.Mutex

type Order = Models.Order

func NewOrder (order *Order) (err error) {
	defer mutex.Unlock()
	err = GetProductByID(& order.Product, strconv.Itoa(int(order.ProductID)))
	if err != nil {
		return err
	}
	err = GetCustomerByID(& order.Customer, strconv.Itoa(int(order.CustomerID)))
	if err != nil {
		return err
	}
	mutex.Lock()
	_ = GetProductByID(& order.Product, strconv.Itoa(int(order.ProductID)))
	if order.Product.Stock < order.Quantity {
		order.OrderStatus = 2
		if err = Config.DB.Create(order).Error ; err != nil {
			return err
		}
		return nil
	}
	order.Product.Stock = order.Product.Stock - order.Quantity
	UpdateProduct(& order.Product)
	order.OrderStatus = 0
	if err = Config.DB.Create(order).Error ; err != nil {
		return err
	}
	return nil
}

func GetOrderByID (order *Order ,id string) (err error) {
	if err = Config.DB.Model(&Order{}).Where("id = ?", id).First(order).Error ; err != nil {
		return err
	}
	return nil
}

func UpdateOrder (order *Order) (err error) {
	Config.DB.Save(order)
	return nil
}




