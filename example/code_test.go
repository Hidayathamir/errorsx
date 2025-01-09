package example

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Hidayathamir/errorsx"
	"github.com/stretchr/testify/assert"
)

func Test_SetCode(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.SetCode(err, http.StatusBadRequest)

	assert.Equal(t, "--400-- err1", err.Error())
	assert.Equal(t, "err1", errorsx.GetMessage(err))
	assert.Equal(t, http.StatusBadRequest, errorsx.GetCode(err)) // <---
}
