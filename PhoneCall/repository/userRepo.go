package repository

import (
	"PhoneCall/model"
	"PhoneCall/service/connection"
)

type UserRepo interface {
	CreateNewUser(user *model.User) (*model.User, error)
	GetUserById(id int64) (*model.UserInfo, error)
	GetUsers(paging *model.Paging) ([]*model.UserInfo, error)
	UpdateUser(user *model.UserUpdate, id int64) (*model.UserUpdate, error)
	DeleteUser(id int64) error
	VerifyValueField(fieldName string, valueField string) (*model.User, error)
	UpdateValueFields(id int64, updates map[string]interface{}) error
	GetNumberOfUsers() (int64, error)
}

type UserRepoImpl struct {
	MySQL *connection.MySQL
}

func NewUserRepoImpl(MySQL *connection.MySQL) *UserRepoImpl {
	return &UserRepoImpl{MySQL: MySQL}
}

func (UserRepo *UserRepoImpl) CreateNewUser(user *model.User) (*model.User, error) {
	if err := UserRepo.MySQL.SQL.Table("users").Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (UserRepo *UserRepoImpl) GetUserById(id int64) (*model.UserInfo, error) {
	user := &model.UserInfo{}
	err := UserRepo.MySQL.SQL.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (UserRepo *UserRepoImpl) GetUsers(paging *model.Paging) ([]*model.UserInfo, error) {
	var users []*model.UserInfo

	err := UserRepo.MySQL.SQL.Table("users").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&users).Error
	if err != nil {

		return nil, err
	}
	return users, nil
}

func (UserRepo *UserRepoImpl) UpdateUser(user *model.UserUpdate, id int64) (*model.UserUpdate, error) {
	err := UserRepo.MySQL.SQL.
		Table("users").
		Where("id = ?", id).
		Updates(&user).Error
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

func (UserRepo *UserRepoImpl) VerifyValueField(fieldName string, valueField string) (*model.User, error) {
	var user *model.User
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

func (UserRepo *UserRepoImpl) GetNumberOfUsers() (int64, error) {
	var count int64
	if err := UserRepo.MySQL.SQL.Table("users").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
