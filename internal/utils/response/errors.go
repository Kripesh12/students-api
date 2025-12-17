package response

import (
	"net/http"
)

var (
	ErrNotFound = HttpError{
		Code:    http.StatusNotFound,
		Message: "resource not found",
	}

	ErrBadRequest = HttpError{
		Code:    http.StatusBadRequest,
		Message: "bad request",
	}

	ErrUnauthorized = HttpError{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
	}

	ErrForbidden = HttpError{
		Code:    http.StatusForbidden,
		Message: "forbidden",
	}
)
