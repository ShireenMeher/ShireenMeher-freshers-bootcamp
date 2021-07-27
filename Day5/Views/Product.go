package Views

import (
	"Day5/Models"
	"sync"
)

type ProductPost struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
	Message string `json:"message"`
}

type Order = Models.Order
type Product = Models.Product

var wg sync.WaitGroup

func ProductPostView(product *Product, message int) (productPost *ProductPost) {
	productPost = &ProductPost{}
	productPost.ID = product.Slug
	productPost.ProductName = product.Name
	productPost.Quantity = product.Stock
	productPost.Price = product.Price
	switch message {
	case 0: productPost.Message = "product successfully added"
	case 1: productPost.Message = "adding product unsuccessful"
	case 2:
		productPost.Message = ""
		wg.Done()
	}
	return
}

func GetProductsView(products *[]Product) (productPosts []ProductPost){
	for i := 0 ; i<len(*products) ; i++ {
		wg.Add(1)
		go func() {
			productPosts = append(productPosts, *ProductPostView( &((*products)[i]), 2))
		} ()
	}
	return productPosts
}
