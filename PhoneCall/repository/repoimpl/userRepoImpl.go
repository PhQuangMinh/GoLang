package repoimpl

import (
	"PhoneCall/driver"
	models "PhoneCall/models"
)

type UserRepoImpl struct {
	MySQL *driver.MySQL
}

func NewUserRepoImpl(MySQL *driver.MySQL) *UserRepoImpl {
	return &UserRepoImpl{MySQL: MySQL}
}

func (UserRepo *UserRepoImpl) PostUser(user *models.User) (*models.User, error) {
	if err := UserRepo.MySQL.SQL.Table("users").Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (UserRepo *UserRepoImpl) GetUserById(id int64) (*models.User, error) {
	user := &models.User{}
	err := UserRepo.MySQL.SQL.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (UserRepo *UserRepoImpl) GetUsers() ([]*models.User, error) {
	var users []*models.User

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
