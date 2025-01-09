package errorsx

import (
	"errors"
	"fmt"
	"regexp"
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

	re := regexp.MustCompile("~~(.*?)~~")
	matches := re.FindStringSubmatch(err.Error())
	var msg string
	if len(matches) > 1 {
		msg = matches[1]
	} else {
		errMsgList := strings.Split(err.Error(), ":: ")
		msg = errMsgList[len(errMsgList)-1]
	}

	// rm code
	re = regexp.MustCompile("--(.*?)--")
	x := re.ReplaceAll([]byte(msg), []byte(""))

	msg = strings.TrimSpace(string(x))

	return msg
}
