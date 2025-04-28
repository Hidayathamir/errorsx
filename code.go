package errorsx

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// SetCode set error code based on net/http status code.
func SetCode(err error, code int) error {
	if err == nil {
		return fmt.Errorf("--%d--", code)
	}
	return fmt.Errorf("--%d--:: %w", code, err)
}

// GetCode return http status code from error. If code not found will return http.StatusInternalServerError.
func GetCode(err error) int {
	if err == nil {
		return http.StatusInternalServerError
	}

	errStrList := strings.Split(err.Error(), ":: ")
	for _, errStr := range errStrList {
		if len(errStr) < 4 {
			continue
		}
		prefix := errStr[:2]
		middle := errStr[2 : len(errStr)-2]
		suffix := errStr[len(errStr)-2:]
		isValidFormat := prefix == "--" && suffix == "--"
		if !isValidFormat {
			continue
		}
		code, err := strconv.Atoi(middle)
		if err != nil {
			return http.StatusInternalServerError
		}
		return code
	}

	return http.StatusInternalServerError
}
