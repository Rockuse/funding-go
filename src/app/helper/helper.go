package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string
	Code    uint16
	Status  string
}

func ResponseHelper(message string, code uint16, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	responseJSON := Response{
		Meta: meta,
		Data: data,
	}
	return responseJSON
}
