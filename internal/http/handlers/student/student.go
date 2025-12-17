package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kripesh12/students-api/internal/types"
	"github.com/kripesh12/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteError(w, response.NewBadRequest("request body should not be empty"))
			return
		}

		if err != nil {
			response.WriteError(w, response.NewBadRequest(err.Error()))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validationErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidateError(validationErr))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
		fmt.Println(student)
	}
}
