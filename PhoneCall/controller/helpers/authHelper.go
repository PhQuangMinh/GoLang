package helpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// MatchUserTypeId
// Kiểm tra xem có được truy cập thông tin hay không, user, khác id -> không được/*
func MatchUserTypeId(c *gin.Context, userId int64) error {
	userType := c.GetString("user_type")
	id, err := strconv.ParseInt(c.GetString("id"), 10, 64)
	fmt.Println(id, userId, userType)
	if err != nil {
		return err
	}
	if userType == "USER" && id != userId {
		return errors.New("unauthorized to access this resource")
	}
	return nil
}

func CheckUserType(c *gin.Context, userType string) error {
	userTypeUser := c.GetString("user_type")
	if userType != userTypeUser {
		return errors.New("unauthorized to access this resource")
	}
	return nil
}
