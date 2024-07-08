package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type Topic struct {
	gorm.Model
	Title    string
	Category string
	UserID   uint
	User     User
}

type Comment struct {
	gorm.Model
	Content string
	TopicID uint
	Topic   Topic
	UserID  uint
	User    User
}

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);unique"`
}
