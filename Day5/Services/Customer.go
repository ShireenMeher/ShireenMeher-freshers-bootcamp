package Services

import (
	"Day5/Config"
	"Day5/Models"
	"time"
)

type Customer = Models.Customer

func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Model(&Customer{}).Create(customer).Error ; err != nil {
		return err
	}
	return nil
}

func GetOrdersByCustomer (orders *[]Order, customerID uint) (err error) {
	if err = Config.DB.Model(&Order{}).Where(&Order{CustomerID : customerID}).Find(orders).Error ;
			err != nil {
		return err
	}
	return nil
}

func GetCustomerByID (customer *Customer, id string) (err error) {
	if err = Config.DB.Model(&Customer{}).Where("id = ?", id).First(customer).Error ; err != nil {
		return err
	}
	return nil
}

func CheckFeasibility (customerID uint) (feasibility bool, err error) {
	var orders []Order
	err = GetOrdersByCustomer(&orders, customerID)
	if err != nil {
		return feasibility, err
	}
	if len(orders) == 0 {
		return true, nil
	}
	lastCreated := orders[len(orders)-1].CreatedAt
	if time.Now().Nanosecond() < lastCreated.Add(time.Minute*5).Nanosecond() {
		return false, nil
	}
	return  true, nil
}

