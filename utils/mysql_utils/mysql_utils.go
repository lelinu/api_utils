package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/lelinu/api_utils/utils/error_utils"
	"strings"
)

const(
	ErrCodeUniqueKey = 1062
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *error_utils.ApiError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok{
		if strings.Contains(err.Error(), ErrorNoRows){
			return error_utils.NewInternalServerError("no records found")
		}
		return error_utils.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case ErrCodeUniqueKey:
		return error_utils.NewBadRequestError("invalid data")
	}
	return error_utils.NewInternalServerError("error processing request")
}
