package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	storage "github.com/piyushk8/StudentAPI/internal/Storage"
	"github.com/piyushk8/StudentAPI/internal/types"
	response "github.com/piyushk8/StudentAPI/internal/utils"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		// body is empty
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, 400, response.ErrorResponse(err.Error()))
			return
		}

		// error in body parsing
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ErrorResponse(err.Error()))
			return
		}

		// invalid data
		if err := validator.New().Struct(student); err != nil {
			validationError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validationError))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.WriteJson(w, http.StatusCreated, types.Response{Success: true, Message: `lastId:` + string(lastId)})
	}
}
