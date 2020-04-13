package service

import (
	domain "github.com/psinthorn/F2Go/domain/contact"
	"github.com/psinthorn/F2Go/utils"
)

func GetContact(id int64) (*domain.Contact, *utils.ApplicationError) {
	return domain.GetContact(id)
}
