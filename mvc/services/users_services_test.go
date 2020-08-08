package services

import (
	"github.com/sswapnil2/golang-microservices/mvc/domain"
	"github.com/sswapnil2/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	//userDaoMock usersDaoMock
	getUserFunction func(int64) (*domain.User, *utils.ApplicationError)

)

func init(){
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {}

func (u *usersDaoMock) GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(id)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exists",
		}
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}

	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}