package errorsx

import (
	"runtime"
	"strings"
)

const defaultSkip = 1

// WrapAddFuncName return wrapped error with added func name.
func WrapAddFuncName(err error, options ...OptionWrapAddFuncName) error {
	option := &OptWrapAddFuncName{Skip: defaultSkip}
	for _, opt := range options {
		opt(option)
	}

	funcName := "?"

	pc, _, _, ok := runtime.Caller(option.Skip)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			funcNameWithModule := fn.Name()
			funcNameWithModuleSplit := strings.Split(funcNameWithModule, "/")
			funcName = funcNameWithModuleSplit[len(funcNameWithModuleSplit)-1]
		}
	}

	return Wrap(err, funcName)
}

type OptWrapAddFuncName struct{ Skip int }

type OptionWrapAddFuncName func(*OptWrapAddFuncName)

// WithSkip sets the number of stack frames to skip.
func WithSkip(skip int) OptionWrapAddFuncName {
	return func(o *OptWrapAddFuncName) {
		o.Skip = skip + defaultSkip
	}
}
