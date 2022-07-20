package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Category    string
	Price       int
	ImageUrl    string
}
