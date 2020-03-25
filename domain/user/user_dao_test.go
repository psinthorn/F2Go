package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotExist(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "Not expect to get any user on id 0")
	assert.NotNil(t, err, "Expect to get error from get user id 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, err.Message, "user id 0 not exist")

}

func TestGetUser(t *testing.T) {
	user, err := GetUser(1)

	assert.Nil(t, err, "Expect to get nil on error")
	assert.NotNil(t, user, "Expect to get user id 1")
	assert.EqualValues(t, http.StatusOK, 200)
	assert.EqualValues(t, user.FirstName, "Sinthorn")
}
