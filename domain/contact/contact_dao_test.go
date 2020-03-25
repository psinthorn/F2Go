package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContactContactNotFound(t *testing.T) {
	contact, err := GetContact(0)

	// use streche/testify for testing
	assert.Nil(t, contact, "We not expecting to get any contact from id 0")
	assert.NotNil(t, err, "Expected error on contact id 0")
	assert.EqualValues(t, "contact_id 0 not exist", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)

	// // write own code to test same as testify above
	// if user != nil {
	// 	t.Error("We not expecting to get any contact from id 0")
	// }

}

func TestGetContact(t *testing.T) {
	contact, err := GetContact(1)

	assert.Nil(t, err, "expect to get nil error")
	assert.NotNil(t, contact, "expect to get contact with id 1")
	assert.EqualValues(t, http.StatusOK, 200)
	assert.EqualValues(t, contact.SubTitle, "F2 Co.,Ltd.")

}
