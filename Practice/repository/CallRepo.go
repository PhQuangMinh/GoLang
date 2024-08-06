package repository

import (
	"Practice/model"
	"Practice/repository/repoimpl"
	"github.com/gin-gonic/gin"
	"time"
)

type CallRepo interface {
	GetCalls(startAt, endAt time.Time) ([]*model.Call, error)
	GetCallByID(callID int64) (*model.Call, error)
	GetValueField(callID int64, displayField string) (model.Call, error)
	Post(rabbit *repoimpl.RabbitMQ, nameQueue string) func(context *gin.Context)
	Update(data model.Call) func(context *gin.Context)
}
