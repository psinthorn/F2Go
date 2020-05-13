package services

import (
	domain "github.com/psinthorn/F2Go/domain/welcome"
	errors "github.com/psinthorn/F2Go/utils"
)

type welcomeService struct{}

var WelcomeService welcomeService

func (ws *welcomeService) GetWelcome(id int64) (*domain.Welcome, *errors.RestErr) {
	return domain.WelcomeDao.GetWelcome(1)
}
