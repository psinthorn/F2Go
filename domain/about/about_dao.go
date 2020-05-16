package abouts

import (
	"fmt"
	"log"

	utils "github.com/psinthorn/F2Go/utils/errors"
)

const (
	queryInsertAbout = "INSERT INTO abouts() VALUES(?,?,?,?,?);"
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

func (a *aboutDao) GetAbout(id int64) (*About, *utils.RestErr) {
	//Connect database

	//Statement preparing

	//Close connection

	//Return result

	log.Println("we're accessing about database")
	if about := abouts[id]; about != nil {
		return about, nil
	}

	return nil, utils.NewNotFoundError(fmt.Sprintf("About with id %v not exist", id))

}
