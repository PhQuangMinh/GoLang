package repository

import (
	models "PhoneCall/model"
	"PhoneCall/service/connection"
)

type UserRepo interface {
	CreateNewUser(user *models.UserInfo) (*models.UserInfo, error)
	GetUserById(id int64) (*models.UserInfo, error)
	GetUsers() ([]*models.UserInfo, error)
	UpdateUser(user *models.User, id int64) (*models.User, error)
	DeleteUser(id int64) error
	VerifyValueField(fieldName string, valueField string) (*models.User, error)
	UpdateValueFields(id int64, updates map[string]interface{}) error
}

type UserRepoImpl struct {
	MySQL *connection.MySQL
}

func NewUserRepoImpl(MySQL *connection.MySQL) *UserRepoImpl {
	return &UserRepoImpl{MySQL: MySQL}
}

func (UserRepo *UserRepoImpl) CreateNewUser(user *models.UserInfo) (*models.UserInfo, error) {
	if err := UserRepo.MySQL.SQL.Table("users").Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (UserRepo *UserRepoImpl) GetUserById(id int64) (*models.UserInfo, error) {
	user := &models.UserInfo{}
	err := UserRepo.MySQL.SQL.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (UserRepo *UserRepoImpl) GetUsers() ([]*models.UserInfo, error) {
	var users []*models.UserInfo

	err := UserRepo.MySQL.SQL.Table("users").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (UserRepo *UserRepoImpl) UpdateUser(user *models.User, id int64) (*models.User, error) {
	err := UserRepo.MySQL.SQL.Table("users").Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (UserRepo *UserRepoImpl) DeleteUser(id int64) error {
	err := UserRepo.MySQL.SQL.Table("users").Where("id = ?", id).Delete(nil).Error
	if err != nil {
		return err
	}
	return nil
}

func (UserRepo *UserRepoImpl) VerifyValueField(fieldName string, valueField string) (*models.User, error) {
	var user *models.User
	err := UserRepo.MySQL.SQL.
		Table("users").
		Where(fieldName+" = ?", valueField).
		First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (UserRepo *UserRepoImpl) UpdateValueFields(id int64, updates map[string]interface{}) error {
	err := UserRepo.MySQL.SQL.Table("users").
		Where("id = ?", id).
		Updates(updates).Error
	if err != nil {
		return err
	}
	return nil
}
