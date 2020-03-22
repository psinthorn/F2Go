package services

import (
	"github.com/psinthorn/F2Go/domain"
	"github.com/psinthorn/F2Go/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
