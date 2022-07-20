package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Status  string
	Value   int
	OrderId int
	UserId  int
}
