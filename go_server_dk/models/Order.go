package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	OrderProducts []Product
	Price         int
	UserId        string
	Status        string
}
