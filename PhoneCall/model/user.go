package model

import "time"

type User struct {
	Id          int64     `json:"id" gorm:"column:id"`
	FirstName   *string   `json:"first_name" gorm:"column:first_name" validate:"required"`
	LastName    *string   `json:"last_name" gorm:"column:last_name" validate:"required"`
	UserName    string    `json:"user_name" gorm:"column:user_name" validate:"required"`
	Password    string    `json:"password" gorm:"column:password" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number" validate:"required"`
	Email       string    `json:"email" gorm:"column:email" validate:"required"`
	UserType    string    `json:"user_type" gorm:"column:user_type" validate:"required"`
	CreateAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdateAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type UserInfo struct {
	FirstName   *string `json:"first_name" gorm:"column:first_name" validate:"required"`
	LastName    *string `json:"last_name" gorm:"column:last_name" validate:"required"`
	UserName    string  `json:"user_name" gorm:"column:user_name" validate:"required"`
	Password    string  `json:"password" gorm:"column:password" validate:"required"`
	PhoneNumber string  `json:"phone_number" gorm:"column:phone_number" validate:"required"`
	Email       string  `json:"email" gorm:"column:email" validate:"required"`
	UserType    string  `json:"user_type" gorm:"column:user_type" validate:"required"`
}
