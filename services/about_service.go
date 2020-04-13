package service

import (
	"github.com/psinthorn/F2Go/domain"
	"github.com/psinthorn/F2Go/utils"
)

type aboutService struct{}

var AboutService aboutService

func (a *AboutService) GetAbout(id int64) (*domain.About, *utils.ApplicationError) {
	return *domain(id)
}
