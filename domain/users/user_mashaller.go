package users

import "github.com/psinthorn/F2Go/utils"

type Public struct {
	Id          int64  `json:"user_id"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created`
}

type Private struct {
	Id          int64  `json:"id"`
	First_Name  string `json:"first_name`
	Last_Name   string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password:`
	Status      string `json:"status"`
	DateCreated string `json:"date_created`
}

func Mashal(isPubllic bool) Users {
	utils.ImplemenftMe()
	if isPubllic {
		return nil
	}

	//Return private
	return nil
}
