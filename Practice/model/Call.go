package model

import "time"

type Call struct {
	Id                int       `json:"id" gorm:"column:id"`
	PhoneNumber       string    `json:"phone_number" gorm:"column:phone_number"`
	Metadata          string    `json:"metadata" gorm:"column:metadata"`
	CallResult        string    `json:"call_result" gorm:"column:call_result"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"column:updated_at"`
	CallTime          time.Time `json:"call_time,omitempty" gorm:"column:call_time"`
	ReceiveResultTime time.Time `json:"receive_result_time,omitempty" gorm:"column:receive_result_time"`
	CallAnsweredTime  time.Time `json:"call_answered_time,omitempty" gorm:"column:call_answered_time"`
	CallEndedTime     time.Time `json:"call_ended_time,omitempty" gorm:"column:call_ended_time"`
}
