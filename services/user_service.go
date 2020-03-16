package services

import "github.com/psinthorn/F2Go/models"

func GetUser(userId int64) (*models.User, error) {
	return models.GetUser(userId)
}
