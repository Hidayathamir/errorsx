package errorsx

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

// SetCode set error code based on net/http status code.
func SetCode(err error, code int) error {
	if err == nil {
		return fmt.Errorf("--%d--", code)
	}
	return fmt.Errorf("--%d-- %w", code, err)
}

// GetCode return http status code from error. If code not found will return http.StatusInternalServerError.
func GetCode(err error) int {
	if err == nil {
		return http.StatusInternalServerError
	}

	re := regexp.MustCompile("--(.*?)--")
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 1 {
		code, err := strconv.Atoi(matches[1])
		if err != nil {
			return http.StatusInternalServerError
		}
		return code
	}

	return http.StatusInternalServerError
}
