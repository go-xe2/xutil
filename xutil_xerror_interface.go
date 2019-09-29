package xutil

type XError interface {
	GetCode() int
	GetMessage() string
	Error() error
}