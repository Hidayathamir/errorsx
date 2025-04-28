package errorsx

import (
	"errors"
	"fmt"
	"strings"
)

// SetMessage set error message. Can be extract using GetMessage func.
func SetMessage(err error, format string, a ...any) error {
	msg := fmt.Sprintf(format, a...)
	if err == nil {
		return errors.New(msg)
	}
	return fmt.Errorf("~~%s~~:: %w", msg, err)
}

// SetMessageE set error message. Can be extract using GetMessage func.
func SetMessageE(err1, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	}
	if err1 == nil {
		return err2
	}
	if err2 == nil {
		return err1
	}
	return fmt.Errorf("~~%w~~:: %w", err2, err1)
}

// GetMessage get message detail from error.
func GetMessage(err error) string {
	if err == nil {
		return ""
	}

	errStrList := strings.Split(err.Error(), ":: ")
	for _, errStr := range errStrList {
		if len(errStr) < 4 {
			continue
		}
		prefix := errStr[:2]
		middle := errStr[2 : len(errStr)-2]
		suffix := errStr[len(errStr)-2:]
		isValidFormat := prefix == "~~" && suffix == "~~"
		if !isValidFormat {
			continue
		}
		return middle
	}

	lastErrStr := errStrList[len(errStrList)-1]
	return lastErrStr
}
