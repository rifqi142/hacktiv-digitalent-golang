package util

import "hacktiv-digitalent-golang/assignment/assignment-2/model"

func CreateResponse(isSuccess bool, data any, errorMessage string) model.Response {
	return model.Response{
		Success: isSuccess,
		Data:    data,
		Message: errorMessage,
	}
}