package services

import (
	"github.com/psinthorn/F2Go/models"
	"github.com/psinthorn/F2Go/utils"
)

func GetContact(id int64) (*models.Contact, *utils.ApplicationError) {
	return models.GetContact(id)
}
