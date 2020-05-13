package services

import (
	domain "github.com/psinthorn/F2Go/domain/about"
	errors "github.com/psinthorn/F2Go/utils"
)

type aboutService struct{}

var AboutService aboutService

func (a *aboutService) GetAbout(id int64) (*domain.About, *errors.RestErr) {
	return domain.AboutDao.GetAbout(id)
}
