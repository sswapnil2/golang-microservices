package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNotFound(t *testing.T) {

	usr, err := GetUser(0)

	assert.Nil(t, usr, "we were expecting user to be nil")

	assert.NotNil(t, err, "we were expecting to have an error")

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "Expected 404 not found status")

}

func TestGetUserNoError(t *testing.T) {

	usr, err := GetUser(123)
	assert.Nil(t, err, "Error should be nil when user")
	assert.NotNil(t, usr, "User should not be nil")
	assert.EqualValues(t, 123, usr.Id)
	assert.EqualValues(t, "Swapnil", usr.FirstName)
	assert.EqualValues(t, "Shindemeshram", usr.LastName)
	assert.EqualValues(t, "smeshram258@gmail.com", usr.Email)


}
