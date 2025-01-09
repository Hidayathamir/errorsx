package example

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Hidayathamir/errorsx"
	"github.com/stretchr/testify/assert"
)

func Test_SetMessage(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.SetMessage(err, "err2")

	assert.Equal(t, "~~err2~~:: err1", err.Error())  // <--- still have full error
	assert.Equal(t, "err2", errorsx.GetMessage(err)) // <---
	assert.Equal(t, http.StatusInternalServerError, errorsx.GetCode(err))
}

var ErrInsufficientStorage = errors.New("insufficient storage")

func Test_SetMessageE(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.SetMessageE(err, ErrInsufficientStorage)

	assert.Equal(t, "~~insufficient storage~~:: err1", err.Error())  // <--- still have full error
	assert.Equal(t, "insufficient storage", errorsx.GetMessage(err)) // <---
	assert.Equal(t, http.StatusInternalServerError, errorsx.GetCode(err))

	assert.True(t, errors.Is(err, ErrInsufficientStorage)) // <--- able to use errors.Is
}
