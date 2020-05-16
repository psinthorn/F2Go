package services

import (
	contacts "github.com/psinthorn/F2Go/domain/contact"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

type contactService struct{}

var ContactService contactService

func (c *contactService) GetContact(id int64) (*contacts.Contact, *utils.RestErr) {
	return contacts.ContactDao.GetContact(id)
}
