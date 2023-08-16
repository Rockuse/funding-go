package helper

type response struct {
	meta meta
	data interface{}
}

type meta struct {
	message string
	code    uint8
	status  string
}

func ResponseHelper(message string, code uint8, status string, data interface{}) response {
	meta := meta{
		message: message,
		code:    code,
		status:  status,
	}
	response := response{
		meta: meta,
		data: data,
	}
	return response
}
