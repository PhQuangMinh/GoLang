package model

import "time"

type User struct {
	Id          int64     `json:"id" gorm:"column:id;type:id;primaryKey;autoIncrement"`
	FirstName   *string   `json:"first_name" gorm:"column:first_name;type:varchar(255);not null" validate:"required"`
	LastName    *string   `json:"last_name" gorm:"column:last_name;type:varchar(255);not null" validate:"required"`
	UserName    string    `json:"user_name" gorm:"column:user_name;type:varchar(255);not null" validate:"required"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(255);not null" validate:"required,min=8,max=32"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(30);not null" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100);not null" validate:"required"`
	UserType    string    `json:"user_type" gorm:"column:user_type;type:varchar(50);not null" validate:"required"`
	CreateAt    time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateAt    time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

type UserInfo struct {
	FirstName   *string `json:"first_name" gorm:"column:first_name;type:varchar(255);not null" validate:"required"`
	LastName    *string `json:"last_name" gorm:"column:last_name;type:varchar(255);not null" validate:"required"`
	UserName    string  `json:"user_name" gorm:"column:user_name;type:varchar(255);not null" validate:"required"`
	PhoneNumber string  `json:"phone_number" gorm:"column:phone_number;type:varchar(30);not null" validate:"required"`
	Email       string  `json:"email" gorm:"column:email;type:varchar(100);not null" validate:"required"`
	UserType    string  `json:"user_type" gorm:"column:user_type;type:varchar(50);not null" validate:"required"`
}

type UserUpdate struct {
	FirstName   *string `json:"first_name" gorm:"column:first_name;type:varchar(255);not null" validate:"required"`
	LastName    *string `json:"last_name" gorm:"column:last_name;type:varchar(255);not null" validate:"required"`
	PhoneNumber string  `json:"phone_number" gorm:"column:phone_number;type:varchar(30);not null" validate:"required"`
	Email       string  `json:"email" gorm:"column:email;type:varchar(100);not null" validate:"required"`
}
