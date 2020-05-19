package mysql_errors

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *utils.RestErr {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return utils.NewNotFoundError("No records found")
		}
		return utils.NewInternalServerError(fmt.Sprintf("Error parsing database response"))
	}

	switch mysqlErr.Number {
	case 1062:
		return utils.NewInternalServerError("Invalid data or email already exist")
	}
	return utils.NewInternalServerError("Error processing request")
}
