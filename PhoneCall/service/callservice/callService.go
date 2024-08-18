package callservice

import (
	"PhoneCall/repository"
)

type CallService struct {
	CallRepo repository.CallRepo
}

func NewCallService(CallRepo repository.CallRepo) *CallService {
	return &CallService{CallRepo: CallRepo}
}
