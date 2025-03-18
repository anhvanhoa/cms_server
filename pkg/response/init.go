package pkgres

import (
	"net/http"
)

type ErrorApp struct {
	Message string      `json:",omitempty"`
	Data    interface{} `json:",omitempty"`
	code    int
}

func NewErr(msg string) *ErrorApp {
	return &ErrorApp{
		Message: msg,
		code:    http.StatusInternalServerError,
	}
}

func Err(err error) *ErrorApp {
	return &ErrorApp{
		Message: err.Error(),
		code:    http.StatusInternalServerError,
	}
}

func (r *ErrorApp) Error() string {
	return r.Message
}

func (r *ErrorApp) SetMessage(message string) *ErrorApp {
	r.Message = message
	return r
}

func (r *ErrorApp) SetData(data interface{}) *ErrorApp {
	r.Data = data
	return r
}

func (r *ErrorApp) Code(code int) *ErrorApp {
	r.code = code
	return r
}

func (r *ErrorApp) GetCode() int {
	return r.code
}

func (r *ErrorApp) BadReq() *ErrorApp {
	return r.Code(http.StatusBadRequest)
}

func (r *ErrorApp) UnprocessableEntity() *ErrorApp {
	return r.Code(http.StatusUnprocessableEntity)
}

func (r *ErrorApp) InternalServerError() *ErrorApp {
	return r.Code(http.StatusInternalServerError)
}

func (r *ErrorApp) NotFound() *ErrorApp {
	return r.Code(http.StatusNotFound)
}

func (r *ErrorApp) Unauthorized() *ErrorApp {
	return r.Code(http.StatusUnauthorized)
}

func (r *ErrorApp) Forbidden() *ErrorApp {
	return r.Code(http.StatusForbidden)
}

func (r *ErrorApp) Conflict() *ErrorApp {
	return r.Code(http.StatusConflict)
}

type Response interface {
	SetMessage(message string) *response
	SetData(data interface{}) *response
	Code(code int) *response
	GetCode() int
}

type response struct {
	Message string      `json:",omitempty"`
	Data    interface{} `json:",omitempty"`
	code    int
}

func NewRes(msg string) *response {
	return &response{
		Message: msg,
	}
}

func ResData(data interface{}) *response {
	return &response{
		Data: data,
	}
}

func (r *response) SetMessage(message string) *response {
	r.Message = message
	return r
}

func (r *response) SetData(data interface{}) *response {
	r.Data = data
	return r
}

func (r *response) Code(code int) *response {
	r.code = code
	return r
}

func (r *response) GetCode() int {
	return r.code
}
