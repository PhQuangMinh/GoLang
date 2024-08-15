package repository

import (
	models "PhoneCall/model"
	"PhoneCall/service/connection"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CallRepo interface {
	GetCalls(startAt, endAt time.Time) ([]*models.Call, error)
	GetCallByID(callID int64) (*models.Call, error)
	GetValueField(callID int64, displayField string) (*models.Call, error)
	CreateNewCall(call *models.Call) (*models.Call, error)
	UpdateCall(c *gin.Context, data *models.Call) (*models.Call, error)
	DeleteCall(id int64) error
}

type CallRepoImpl struct {
	MySQL *connection.MySQL
}

func NewCallRepoImpl(MySQL *connection.MySQL) *CallRepoImpl {
	return &CallRepoImpl{MySQL: MySQL}
}

func (callRepo *CallRepoImpl) GetCalls(startAt, endAt time.Time) ([]*models.Call, error) {
	var calls []*models.Call
	err := callRepo.MySQL.SQL.Table("calls").
		Order("id, client_name, phone_number").
		Where("created_at between ? and ?", startAt, endAt).
		Find(&calls).Error

	if err != nil {
		return nil, err
	}

	return calls, nil
}

func (callRepo *CallRepoImpl) GetCallByID(id int64) (*models.Call, error) {
	call := &models.Call{}
	err := callRepo.MySQL.SQL.Table("calls").
		Where("id = ?", id).
		First(&call).Error

	if err != nil {
		return nil, err
	}

	return call, nil
}

func (callRepo *CallRepoImpl) GetValueField(id int64, field string) (*models.Call, error) {
	var data *models.Call
	if field == "" {
		err := callRepo.MySQL.SQL.Table("calls").
			Where("id = ?", id).
			First(&data).
			Error
		if err != nil {
			return data, err
		}
		return data, nil
	}
	err := callRepo.MySQL.SQL.Table("calls").
		Select("id, client_name, phone_number, CONCAT_WS(':',?,JSON_EXTRACT((SELECT metadata from calls where id = ?), ?)) as metadata, "+
			"call_result, created_at, updated_at, call_time, receive_result_time, call_answered_time, "+
			"call_ended_time", "\""+field+"\"", id, "$[0]."+field).
		Where("id = ?", id).
		First(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (callRepo *CallRepoImpl) CreateNewCall(call *models.Call) (*models.Call, error) {
	if err := callRepo.MySQL.SQL.Create(&call).Error; err != nil {
		return nil, err
	}
	return call, nil
}

func (callRepo *CallRepoImpl) UpdateCall(c *gin.Context, data *models.Call) (*models.Call, error) {
	if err := callRepo.MySQL.SQL.Table("calls").
		Where("id = ?", data.Id).
		Updates(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return nil, err
	}
	return data, nil
}

func (callRepo *CallRepoImpl) DeleteCall(id int64) error {
	err := callRepo.MySQL.SQL.Table("calls").
		Where("id = ?", id).
		Delete(&models.Call{}).Error
	if err != nil {
		return err
	}
	return nil
}
