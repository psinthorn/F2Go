package services

import (
	utils "github.com/psinthorn/F2Go/utils/errors"
)

type aboutService struct{}

var AboutService aboutService

func (a *aboutService) GetAbout(id int64) (*abouts.About *utils.RestErr) {
	return abouts.AboutDao.GetAbout(id)
}
