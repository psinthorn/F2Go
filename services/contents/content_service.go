package contents

import (
	"fmt"

	utils "github.com/psinthorn/F2Go/utils/errors"
)

const (
	errorNoRows            = ""
	queryInsertContent     = ""
	querySelectContentById = ""
)

func CreateContent() *utils.RestErr {
	fmt.Println("Save content to database")
	return nil
}

func GetContent() *utils.RestErr {
	fmt.Println("Get content from database")
	return nil
}
