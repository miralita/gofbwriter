package gofbwriter

import (
	"fmt"
	"runtime"
)

type errorCode int

//go:generate stringer -type=errorCode

const (
	//ErrEmptyFirstName - empty required first name
	ErrEmptyFirstName errorCode = iota
	//ErrEmptyField - empty required field
	ErrEmptyField
	//ErrNestedEntity - empty required nested entity
	ErrNestedEntity
	//ErrUnsupportedNestedItem - unsupported nested item
	ErrUnsupportedNestedItem
	//ErrSystem - system error
	ErrSystem
)

type errorf struct {
	etype    errorCode
	message  string
	funcName string
	fileLine int
	next     *errorf
}

func makeError(code errorCode, message string, args ...interface{}) error {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	err := &errorf{etype: code, message: message}
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		err.funcName = details.Name()
		_, err.fileLine = details.FileLine(pc)
	}
	return err
}

func wrapError(err error, code errorCode, message string, args ...interface{}) error {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	e, ok := err.(*errorf)
	if !ok {
		e = &errorf{etype: ErrSystem, message: err.Error()}
	}
	newErr := &errorf{etype: code, message: message, next: e}
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		newErr.funcName = details.Name()
		_, newErr.fileLine = details.FileLine(pc)
	}
	return newErr
}

func (s *errorf) Error() string {
	err := fmt.Sprintf("%s: %s at %s(%d)", s.etype.String(), s.message, s.funcName, s.fileLine)
	if s.next != nil {
		err += "\n" + s.next.Error()
	}
	return err
}
