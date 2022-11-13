package errcode

import (
	"fmt"
	"log"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string
}

var codes = map[int]string{}

func InitError(code int, msg string) Error {
	if _, ok := codes[code]; ok {
		log.Panicf("错误码 %d 已经存在，请更换一个", code)
	} else {
		codes[code] = msg
	}
	return Error{
		code: code,
		msg:  msg,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code(), e.Msg())
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Msg() string {
	return e.msg
}

func (e Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e Error) Details() []string {
	return e.details
}

func (e Error) WithDetails(details ...string) Error {
	e.details = details
	return e
}

func (e Error) ToStatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	case NotFound.Code():
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
