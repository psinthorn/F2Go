package services

import (
	domain "github.com/psinthorn/F2Go/domain/welcome"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

type welcomeService struct{}

var WelcomeService welcomeService

func (ws *welcomeService) GetWelcome(id int64) (*domain.Welcome, *utils.RestErr) {
	return domain.WelcomeDao.GetWelcome(1)
}
