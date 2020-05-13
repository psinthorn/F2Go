package contacts

import (
	"fmt"
	"log"
	"net/http"

	errors "github.com/psinthorn/F2Go/utils"
)

type contactDao struct{}

var ContactDao contactDao

var (
	contacts = map[int64]*Contact{
		1: {Id: 1,
			Title:     "Contact Us",
			SubTitle:  "F2 Co.,Ltd.",
			Body:      "Contact F2 with below infomation and we will reply you within 24 hrs.",
			SubBody:   "",
			Email:     "psinthorn@gmail.com",
			LineId:    "sinthorn83",
			Instagram: "sinthorn",
			Facebook:  "https://fb.me/f2coltd",
			Mobile:    "0989358228",
			Website:   "https://www.f2.co.th",
			Status:    true},
	}
)

func (c *contactDao) GetContact(id int64) (*Contact, *errors.RestErr) {
	log.Println("we're accessing contact database")
	if contact := contacts[id]; contact != nil {
		return contact, nil
	}

	return nil, &errors.RestErr{
		Message:    fmt.Sprintf("contact_id %v not exist", id),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
