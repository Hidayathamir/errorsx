package example

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Hidayathamir/errorsx"
	"github.com/stretchr/testify/assert"
)

func Test_Wrap(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.Wrap(err, "err2")

	assert.Equal(t, "err2:: err1", err.Error()) // <---
	assert.Equal(t, "err1", errorsx.GetMessage(err))
	assert.Equal(t, http.StatusInternalServerError, errorsx.GetCode(err))
}

var ErrRecordNotFound = errors.New("record not found")
var ErrBadRequest = errors.New("bad request")

func Test_WrapE(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.WrapE(err, ErrRecordNotFound)
	err = errorsx.WrapE(err, ErrBadRequest)

	assert.Equal(t, "bad request:: record not found:: err1", err.Error()) // <---
	assert.Equal(t, "err1", errorsx.GetMessage(err))
	assert.Equal(t, http.StatusInternalServerError, errorsx.GetCode(err))

	assert.True(t, errors.Is(err, ErrRecordNotFound)) // <--- able to use errors.Is
	assert.True(t, errors.Is(err, ErrBadRequest))     // <--- able to use errors.Is
}
