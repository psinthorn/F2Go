package services

import (
	domain "github.com/psinthorn/F2Go/domain/welcome"
	"github.com/psinthorn/F2Go/utils"
)

type welcomeService struct{}

var WelcomeService welcomeService

func (ws *welcomeService) GetWelcome(id int64) (*domain.Welcome, *utils.ApplicationError) {
	return domain.WelcomeDao.GetWelcome(1)
}
