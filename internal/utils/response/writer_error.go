package response

import "net/http"

func WriteError(w http.ResponseWriter, err error) error {
	httpErr, ok := err.(HttpError)
	if !ok {
		httpErr = HttpError{
			Code:    500,
			Message: "internal server error",
		}
	}

	return WriteJson(w, httpErr.Code, Response{
		Status: StatusError,
		Error:  httpErr.Message,
	})
}
