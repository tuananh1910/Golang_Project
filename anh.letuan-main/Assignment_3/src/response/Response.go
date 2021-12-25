package response

import "strings"

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Error interface{} `json:"error"`
	Data interface{} `json:"data"`
}

type EmptyObj struct {

}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildErrResponse(message string, err string, data interface{}) Response {
	error := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   error,
		Data:    data,
	}
	return res
}