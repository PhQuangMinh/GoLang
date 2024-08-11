package repository

import models "PhoneCall/models"

type UserRepo interface {
	PostUser(user *models.User) (*models.User, error)
	GetUserById(id int64) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(user *models.User, id int64) (*models.User, error)
	DeleteUser(id int64) error
	VerifyValueField(fieldName string, valueField string) (*models.User, error)
	UpdateValueFields(id int64, updates map[string]interface{}) error
}
