package domain

import (
	"fmt"
	"log"
	"net/http"

	utils "github.com/psinthorn/F2Go/utils/errors"
)

type welcomeDao struct{}

var (
	welcomes = map[int64]*Welcome{
		1: {Id: 1,
			Title:  "Welcome to F2 Co.,Ltd. Official Website",
			Body:   "Our Official website is on development if you can visit us daily to see our website progress is will be great as we will continue development and update as non-stop.",
			Status: true,
		},
	}

	WelcomeDao welcomeDao
)

func (wc *welcomeDao) GetWelcome(id int64) (*Welcome, *utils.RestErr) {
	log.Println("we're access welcome database")
	if welcome := welcomes[id]; welcome != nil {
		return welcome, nil
	}

	return nil, &utils.RestErr{
		Message:    fmt.Sprintf("welcome id %v not exist", id),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
