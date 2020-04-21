package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/psinthorn/F2Go/utils"
)

type aboutDao struct{}

var (
	abouts = map[int64]*About{
		1: {Id: 1,
			Title:    "About Us",
			SubTitle: "About F2",
			Body:     "F2 is modern technology service company  we're focus on web application mobile and iOT development with deployment service.",
			SubBody:  "Web Hosting, Email and DNS Server is also our base service for more than 18 years.",
			Status:   true},
	}
	AboutDao aboutDao
)

func (a *aboutDao) GetAbout(id int64) (*About, *utils.ApplicationError) {
	log.Println("we're accessing about database")
	if about := abouts[id]; about != nil {
		return about, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("About with id %v not exist", id),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
