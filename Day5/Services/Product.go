package Services

import (
	"Day5/Config"
	"Day5/Models"
)

type Product = Models.Product

func GetAlLProducts (products *[]Product) (err error) {
	if err = Config.DB.Model(&Product{}).Find(products).Error ; err != nil {
		return err
	}
	return nil
}

func PostProduct (product *Product) (err error) {
	if err = Config.DB.Model(&Product{}).Create(product).Error ; err != nil {
		return err
	}
	return nil
}

func GetProductByID (product *Product, id string) (err error) {
	if err = Config.DB.Model(&Product{}).Where("id = ?", id).First(product).Error ; err != nil {
		return err
	}
	return nil
}

func UpdateProduct (product *Product) (err error) {
	Config.DB.Model(&Product{}).Save(product)
	return nil
}
