package services

import (
	"github.com/sswapnil2/golang-microservices/mvc/domain"
	"github.com/sswapnil2/golang-microservices/mvc/utils"
)

type usersService struct {}

var (
	UserService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {

	return domain.UserDao.GetUser(userId)

}
