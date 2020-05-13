package abouts

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAboutNotFound(t *testing.T) {
	about, err := GetAbout(0)

	assert.Nil(t, about, "Expect to get nothings from id 0")
	assert.NotNil(t, err, "Expect to get error from get about by id 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, err.Message, "About with id 0 not exist")
}

func TestGetAbout(t *testing.T) {
	about, err := GetAbout(1)

	assert.Nil(t, err, "Not expect to get any error")
	assert.NotNil(t, about, "Expect to get about from id 1")
	assert.EqualValues(t, http.StatusOK, 200)
	assert.EqualValues(t, about.Title, "About Us")
}
