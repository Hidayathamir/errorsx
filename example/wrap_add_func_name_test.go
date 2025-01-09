package example

import (
	"errors"
	"testing"

	"github.com/Hidayathamir/errorsx"
	"github.com/stretchr/testify/assert"
)

func Test_WrapAddFuncName(t *testing.T) {
	err := Four()
	assert.Equal(t, "example.Four:: example.Three:: example.Two:: example.One:: dummy err", err.Error()) // <---
}

func One() error {
	err := errors.New("dummy err")
	return errorsx.WrapAddFuncName(err)
}

func Two() error {
	err := One()
	return errorsx.WrapAddFuncName(err)
}

func Three() error {
	err := Two()
	return errorsx.WrapAddFuncName(err)
}

func Four() error {
	err := Three()
	return errorsx.WrapAddFuncName(err)
}
