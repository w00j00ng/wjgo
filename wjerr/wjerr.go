package wjerr

import (
	"time"
)

type Error struct {
	mCode    EError
	mMessage string
	mTime    string
}

func NewError(code EError, message string) (newError Error) {
	return Error{
		mCode:    code,
		mMessage: message,
		mTime:    time.Now().Format("2006-01-02 15:04:05.000Z"),
	}
}

func InheritError(code EError, err error) (newError Error) {
	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = "no error"
	}
	return Error{
		mCode:    code,
		mMessage: message,
		mTime:    time.Now().Format("2006-01-02 15:04:05.000Z"),
	}
}

func NoError() (newError Error) {
	return Error{
		mCode:    E_ERROR_OK,
		mMessage: "no error",
		mTime:    time.Now().Format("2006-01-02 15:04:05.000Z"),
	}
}

func (e Error) Error() (errLine string) {
	return e.mMessage
}

func (e *Error) GetTime() (mTime string) {
	return e.mTime
}

func (e *Error) GetCode() (mCode EError) {
	return e.mCode
}

func (e *Error) UpdateTime() {
	e.mTime = time.Now().Format("2006-01-02 15:04:05.000Z")
}
