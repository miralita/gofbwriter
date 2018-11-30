package go_fbwriter

import (
	"fmt"
	"runtime"
)

type errorCode int
//go:generate stringer -type=errorCode

const (
	ERR_EMPTY_FIRST_NAME errorCode = iota
	ERR_EMPTY_FIELD
	ERR_SYSTEM
)

type errorf struct {
	etype errorCode
	message string
	funcName string
	fileLine int
	next *errorf
}

func makeError(code errorCode, message string) error {
	err := &errorf{etype: code, message: message}
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		err.funcName = details.Name()
		_, err.fileLine = details.FileLine(pc)
	}
	return err
}

func wrapError(err error, code errorCode, message string) error {
	e, ok := err.(*errorf)
	if !ok {
		e = &errorf{etype: ERR_SYSTEM, message: err.Error()}
	}
	new_err := &errorf{etype: code, message: message, next: e}
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		new_err.funcName = details.Name()
		_, new_err.fileLine = details.FileLine(pc)
	}
	return new_err
}

func (s *errorf) Error() string {
	err := fmt.Sprintf("%s: %s at %s(%d)", s.etype.String(), s.message, s.funcName, s.fileLine)
	if (s.next != nil) {
		err += "\n" + s.next.Error()
	}
	return err
}
