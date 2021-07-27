package Models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Slug  string  `gorm:"unique_index;not null" json:"slug"`
	Price int     `gorm:"not null" json:"price"`
	Stock int     `gorm:"not null" json:"stock"`
	Image string  `json:"image_url"`
}

func (product *Product) BeforeCreate() (err error) {
	idString := strconv.Itoa(int(product.ID))
	idString = idString + strings.Repeat("0", 6- len(idString))
	product.Slug = "ORD" + idString
	return
}