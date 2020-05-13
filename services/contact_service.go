package services

import (
	domain "github.com/psinthorn/F2Go/domain/contact"
	errors "github.com/psinthorn/F2Go/utils"
)

type contactService struct{}

var ContactService contactService

func (c *contactService) GetContact(id int64) (*domain.Contact, *errors.RestErr) {
	return domain.ContactDao.GetContact(id)
}
