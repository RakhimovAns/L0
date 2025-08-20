// Package exerr exerr - explained error)
package exerr

import (
	"errors"
	"fmt"
	"net/http"
)

type Option func(e *Error)

type Error struct {
	message string
	op      string
	code    int
	details string
}

// New creates new exerr instance
//
// default code is http.StatusInternalServerError
func New(message string, options ...Option) *Error {
	err := &Error{
		message: message,
		code:    http.StatusInternalServerError,
	}
	err = err.With(options...)

	return err
}

func Wrap(err error, options ...Option) *Error {
	newErr := &Error{
		message: err.Error(),
	}

	newErr = newErr.With(options...)

	return newErr
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s %s", e.op, e.message, e.details)
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) With(options ...Option) *Error {
	ne := e.Clone()

	for _, option := range options {
		option(ne)
	}

	return ne
}

func (e *Error) Clone() *Error {
	if e == nil {
		return nil
	}
	copy := *e
	return &copy
}

func WithOp(op string) Option {
	return func(e *Error) {
		e.op = op
	}
}

// WithErr just wraps error into already existing error message, which will look like
// some message: some error
func WithErr(err error) Option {
	return func(e *Error) {
		e.message = fmt.Sprintf("%s: %s", e.message, err.Error())
	}
}

func WithCode(code int) Option {
	return func(e *Error) {
		e.code = code
	}
}

func WithDetails(details string, args ...any) Option {
	return func(e *Error) {
		e.details = fmt.Sprintf(details, args...)
	}
}

func Is(err error, target *Error) bool {
	_, ok := to(err)
	if !ok {
		return false
	}

	return errors.Is(err, target)
}

func to(err error) (*Error, bool) {
	var exerr *Error

	if errors.As(err, &exerr) {
		return exerr, true
	}

	return nil, false
}

var (
	ErrInvalidParam = New("invalid param", WithCode(http.StatusUnprocessableEntity))
)
