package response

import "net/http"

func newHttpError(code int, defaultMsg string, msg ...string) HttpError {
	message := defaultMsg
	if len(message) > 0 && msg[0] != "" {
		message = msg[0]
	}
	return HttpError{
		Code:    code,
		Message: message,
	}
}

func NewBadRequest(msg ...string) HttpError {
	return newHttpError(http.StatusBadRequest, "bad request", msg...)
}

func NewNotFound(msg ...string) HttpError {
	return newHttpError(http.StatusNotFound, "resource not found", msg...)
}

func NewUnauthorized(msg ...string) HttpError {
	return newHttpError(http.StatusUnauthorized, "unauthorized", msg...)
}
