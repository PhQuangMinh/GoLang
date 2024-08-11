package repository

import (
	models "PhoneCall/models"
	"PhoneCall/repository/repoimpl"
	"github.com/gin-gonic/gin"
	"time"
)

type CallRepo interface {
	GetCalls(startAt, endAt time.Time) ([]*models.Call, error)
	GetCallByID(callID int64) (*models.Call, error)
	GetValueField(callID int64, displayField string) (models.Call, error)
	Post(rabbit *repoimpl.RabbitMQ, nameQueue string) func(context *gin.Context)
	Update(data models.Call) func(context *gin.Context)
}
