# errorsx

Golang error extra. Wrap error, set error message, set error code.

This project is based on my problem where I want to wrap error to add stack trace automatically [see feature](#3-wrap-error-add-func-name).

Also I want to set error message so my REST API return only error message and I still have full error stack trace [see feature](#4-set-error-message).

Also I want to set error code so my REST API return response code based on what I set [see feature](#6-set-error-code).

## Get Started

```go
go get github.com/Hidayathamir/errorsx
```

```go
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Hidayathamir/errorsx"
)

var ErrForbidden = errors.New("this route is forbidden")

func main() {
	err := errors.New("err1")
	err = errorsx.Wrap(err, "err2")
	err = errorsx.SetMessageE(err, ErrForbidden)
	err = errorsx.SetCode(err, http.StatusForbidden)

	fmt.Println(err.Error())             // --403-- ~~this route is forbidden~~:: err2:: err1
	fmt.Println(errorsx.GetMessage(err)) // this route is forbidden
	fmt.Println(errorsx.GetCode(err))    // 403

	fmt.Println(errors.Is(err, ErrForbidden)) // true, able to use errors.Is
}
```

## Feature

#### 1. Wrap error

[see example](./example/wrap_test.go).

```go
func Test_Wrap(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.Wrap(err, "err2")

	assert.Equal(t, "err2:: err1", err.Error()) // <---
}
```

#### 2. Wrap error, then check error is.

[see example](./example/wrap_test.go).

```go
var ErrRecordNotFound = errors.New("record not found")
var ErrBadRequest = errors.New("bad request")

func Test_WrapE(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.WrapE(err, ErrRecordNotFound)
	err = errorsx.WrapE(err, ErrBadRequest)

	assert.Equal(t, "bad request:: record not found:: err1", err.Error()) // <---

	assert.True(t, errors.Is(err, ErrRecordNotFound)) // <--- able to use errors.Is
	assert.True(t, errors.Is(err, ErrBadRequest))     // <--- able to use errors.Is
}
```

#### 3. Wrap error add func name.

Usefull to adding stack trace. [see example](./example/wrap_test.go).

```go
package example

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
```

#### 4. Set error message.

[see example](./example/message_test.go).

```go
func Test_SetMessage(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.SetMessage(err, "err2")

	assert.Equal(t, "~~err2~~:: err1", err.Error()) // <--- still have full error
	assert.Equal(t, "err2", errorsx.GetMessage(err)) // <---
}
```

#### 5. Set error message, then check error is.

[see example](./example/message_test.go).

```go
var ErrInsufficientStorage = errors.New("insufficient storage")

func Test_SetMessageE(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.SetMessageE(err, ErrInsufficientStorage)

	assert.Equal(t, "~~insufficient storage~~:: err1", err.Error()) // <--- still have full error
	assert.Equal(t, "insufficient storage", errorsx.GetMessage(err)) // <---

	assert.True(t, errors.Is(err, ErrInsufficientStorage)) // <--- able to use errors.Is
}
```

#### 6. Set error code.

[see example](./example/code_test.go).

```go
func Test_SetCode(t *testing.T) {
	err := errors.New("err1")
	err = errorsx.SetCode(err, http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, errorsx.GetCode(err)) // <---
}
```
