package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	OauthId  int
	OauthKey string
	UserKey  string
}
