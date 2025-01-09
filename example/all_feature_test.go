package example

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Hidayathamir/errorsx"
	"github.com/stretchr/testify/assert"
)

var ErrForbidden = errors.New("this route is forbidden")

func Test_AllFeature(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.Wrap(err, "err2")
	err = errorsx.SetMessageE(err, ErrForbidden)
	err = errorsx.SetCode(err, http.StatusForbidden)

	assert.Equal(t, "--403-- ~~forbidden~~:: err2:: err1", err.Error()) // <---
	assert.Equal(t, "this route is forbidden", errorsx.GetMessage(err)) // <---
	assert.Equal(t, http.StatusForbidden, errorsx.GetCode(err))         // <---
}
