package services

import (
	"github.com/psinthorn/F2Go/models"
	"github.com/psinthorn/F2Go/utils"
)

func GetUser(userId int64) (*models.User, *utils.ApplicationError) {
	return models.GetUser(userId)
}
