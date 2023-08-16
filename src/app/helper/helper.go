package helper

type Response struct {
	meta Meta
	data interface{}
}

type Meta struct {
	message string
	code    uint16
	status  string
}

func ResponseHelper(message string, code uint16, status string, data interface{}) Response {
	Meta := Meta{
		message: message,
		code:    code,
		status:  status,
	}
	ResponseJSON := Response{
		meta: Meta,
		data: data,
	}
	return ResponseJSON
}
