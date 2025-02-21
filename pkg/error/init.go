package pkgerror

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Msg    string
	code   int
	stack  *stack
	global bool
}

func New(msg string) *AppError {
	stack := callers(3)
	return &AppError{Msg: msg, stack: stack, global: stack.isGlobal()}
}

func Err(err error) *AppError {
	stack := callers(3)
	return &AppError{Msg: err.Error(), stack: stack, global: stack.isGlobal()}
}

func Errf(format string, args ...interface{}) error {
	stack := callers(3)
	return &AppError{
		global: stack.isGlobal(),
		Msg:    fmt.Sprintf(format, args...),
		stack:  stack,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprint(e.Msg)
}

func (e *AppError) Code(code int) *AppError {
	e.code = code
	return e
}

func (e *AppError) GetCode() int {
	return e.code
}

func (e *AppError) BadRequest() *AppError {
	return e.Code(http.StatusBadRequest)
}

func (e *AppError) InternalServerError() *AppError {
	return e.Code(http.StatusInternalServerError)
}

func (e *AppError) Unauthorized() *AppError {
	return e.Code(http.StatusUnauthorized)
}

func (e *AppError) Forbidden() *AppError {
	return e.Code(http.StatusForbidden)
}

func (e *AppError) NotFound() *AppError {
	return e.Code(http.StatusNotFound)
}

func (e *AppError) Conflict() *AppError {
	return e.Code(http.StatusConflict)
}

func (e *AppError) Gone() *AppError {
	return e.Code(http.StatusGone)
}

func (e *AppError) UnprocessableEntity() *AppError {
	return e.Code(http.StatusUnprocessableEntity)
}
