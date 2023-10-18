package helper

import (
	"fmt"
	"funding/src/config"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseHelper(message string, code int, status string, data interface{}) Response {
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

func FormatValidationError(err error) gin.H {
	var errors []string
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
	}
	errorMessage := gin.H{"errors": errors}

	return errorMessage
}

func Error(err error) gin.H {

	return gin.H{"errors": err}
}
func PathUpload(dst ...string) (string, string) {
	fmt.Println(len(dst))
	path := "public/images"
	file := ""
	for idx, str := range dst {
		if str == "" {
			continue
		}
		if idx == len(dst) {
			lastIndex := strings.Split(str, ".")
			idFile := config.Uuid()
			rename := strconv.Itoa(idFile) + "." + lastIndex[len(lastIndex)-1]
			path += "/" + str
			file += "/" + rename
			continue
		}
		path += "/" + str
		file += "/" + str
	}
	return path, file
}
