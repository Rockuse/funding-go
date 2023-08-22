package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    uint16 `json:"code"`
	Status  string `json:"status"`
}

func ResponseHelper(message string, code uint16, status string, data interface{}) Response {
	Meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	ResponseJSON := Response{
		Meta: Meta,
		Data: data,
	}
	return ResponseJSON
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
