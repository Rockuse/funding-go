package helper

import (
	"fmt"
	util "funding/src/app/common/utilities"
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
		fmt.Println(err)
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
	}
	errorMessage := gin.H{"errors": errors}
	return errorMessage
}

func PathUpload(dst ...string) (string, string) {
	rename := strconv.Itoa(util.Uuid())
	newPath := "public/images"
	fileName := ""
	for idx, str := range dst {
		if str == "" {
			continue
		}
		if len(dst)-1 == idx {
			ext := strings.Split(str, ".")
			// fmt.Println(ext[len(ext)-1])
			newPath += "/" + rename + "." + ext[len(ext)-1]
			fileName += "/" + rename + "." + ext[len(ext)-1]
			continue
		}
		newPath += "/" + str
		fileName += "/" + str
	}
	return newPath, fileName
}
