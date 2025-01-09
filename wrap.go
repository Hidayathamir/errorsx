package errorsx

import (
	"errors"
	"fmt"
	"strings"
)

// Wrap return wrapped error.
func Wrap(err error, format string, a ...any) error {
	msg := fmt.Sprintf(format, a...)
	if err == nil {
		return errors.New(msg)
	}
	return fmt.Errorf("%s:: %w", msg, err)
}

// WrapE return wrapped error.
func WrapE(err1 error, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	}
	if err1 == nil {
		return err2
	}
	if err2 == nil {
		return err1
	}
	return fmt.Errorf("%w:: %w", err2, err1)
}

// UnwrapToList unwraps error to list string.
func UnwrapToList(err error) []string {
	if err == nil {
		return nil
	}
	msg := err.Error()
	return strings.Split(msg, ":: ")
}
