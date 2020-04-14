package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)

}
