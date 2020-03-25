package services

import (
	domain "github.com/psinthorn/F2Go/domain/user"
	"github.com/psinthorn/F2Go/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
