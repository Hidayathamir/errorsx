package example

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Hidayathamir/errorsx"
	"github.com/stretchr/testify/assert"
)

func Test_ErrorIsNil(t *testing.T) {
	var err error

	assert.Equal(t, nil, err)
	assert.Equal(t, "", errorsx.GetMessage(err))
	assert.Equal(t, http.StatusInternalServerError, errorsx.GetCode(err))
}

func Test_ErrorIsNotNil(t *testing.T) {
	err := errors.New("err1")

	assert.Equal(t, "err1", err.Error())
	assert.Equal(t, "err1", errorsx.GetMessage(err))
	assert.Equal(t, http.StatusInternalServerError, errorsx.GetCode(err))
}
