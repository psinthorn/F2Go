package services

import (
	domain "github.com/psinthorn/F2Go/domain/user"
	"github.com/psinthorn/F2Go/utils"
)

type userService struct{}

var UserService userService

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
