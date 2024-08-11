package repoimpl

import (
	"PhoneCall/driver"
	models "PhoneCall/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CallRepoImpl struct {
	MySQL *driver.MySQL
}

func NewCallRepoImpl(MySQL *driver.MySQL) *CallRepoImpl {
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

func (callRepo *CallRepoImpl) GetValueField(id int64, field string) (models.Call, error) {
	var data models.Call
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

func (callRepo *CallRepoImpl) Post(rabbit *RabbitMQ, nameQueue string) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data models.Call
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
			return
		}

		if err := callRepo.MySQL.SQL.Create(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error2": err.Error(),
			})
			return
		}
		rabbit.Push(nameQueue, data)
		context.JSON(http.StatusOK, data)
	}
}

func (callRepo *CallRepoImpl) Update(data models.Call) func(context *gin.Context) {
	return func(c *gin.Context) {
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		if err := callRepo.MySQL.SQL.Table("calls").
			Where("id = ?", data.Id).
			Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
