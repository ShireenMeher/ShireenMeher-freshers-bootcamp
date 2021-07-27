package main

import (
	"Day5/Config"
	"Day5/Models"
	"Day5/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql",
		Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		panic(err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{})
	Config.DB.AutoMigrate(&Models.Order{})
	r := Routes.SetUpRouter()
	r.Run()
}
