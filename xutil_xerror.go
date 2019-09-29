package xutil

import (
	"errors"
	"runtime"
)

type xError struct {
	code int
	message string
	source [][]interface{}
}

func (e *xError) GetCode() int {
	return e.code
}

func (e *xError) GetMessage() string {
	return e.message
}

func (e *xError) Error() error {
	return errors.New(e.message)
}

// 创建错误类
func NewError(msg string, code int) XError {
	var source [][]interface{}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		source = append(source, []interface{}{line, file})
	}
	return &xError{
		code: code,
		message: msg,
		source: source,
	}
}

// 抛出错误
func Thown(msg string, code int) {
	err := NewError(msg, code)
	panic(err)
}

