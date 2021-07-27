package Models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type Customer struct {
	gorm.Model
	Slug string `gorm:"unique_index" json:"slug_id"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName string `json:"lastName"`
	Email string `gorm:"unique_index" json:"email"`
}

func (customer *Customer) BeforeCreate() (err error) {
	idString := strconv.Itoa(int(customer.ID))
	idString = idString + strings.Repeat("0", 6- len(idString))
	customer.Slug = "PROD" + idString
	return
}