package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/piyushk8/StudentAPI/internal/types"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("content-type", "application/json")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func ErrorResponse(err string) types.Response {
	return types.Response{Success: false, Message: err}
}


func ValidationError(errs validator.ValidationErrors) types.Response{
	
	var errmsgs []string

	for _,err := range errs {
		switch err.ActualTag(){
		case "required":
			errmsgs = append(errmsgs, fmt.Sprintf("field %s is required", err.Field()))
		
		default :
			errmsgs = append(errmsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}

	}
	return   types.Response{
		Success: false,
		Message: strings.Join(errmsgs, ", "),
	}
	
}