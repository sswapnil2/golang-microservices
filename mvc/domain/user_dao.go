package domain

import (
	"fmt"
	"github.com/sswapnil2/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User {
		123:  {
			Id: 123,
			FirstName: "Swapnil",
			LastName: "Shindemeshram",
			Email: "smeshram258@gmail.com",
		},
	}

	UserDao DoaUserInterface
)

type DoaUserInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {}

func init() {
	UserDao = &userDao{}
}


func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError){

	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil,  &utils.ApplicationError{
		Message: fmt.Sprintf("User %v not found", userId),
		Status: "not_found",
		StatusCode: http.StatusNotFound,
	}

}
