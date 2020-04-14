package services

import (
	domain "github.com/psinthorn/F2Go/domain/about"
	"github.com/psinthorn/F2Go/utils"
)

type aboutService struct{}

var AboutService aboutService

func (a *aboutService) GetAbout(id int64) (*domain.About, *utils.ApplicationError) {
	return domain.AboutDao.GetAbout(id)
}
