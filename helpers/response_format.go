package helpers

import "delos-farm-backend/domains"

type Response struct {
	Message string        `json:"message"`
	Success bool          `json:"success"`
	Data    interface{}   `json:"data"`
	Stats   domains.Stats `json:"statistics"`
}

func ResponseFormat(
	msg string,
	isSuccess bool,
	data interface{},
	stats domains.Stats,
) Response {
	return Response{
		Message: msg,
		Success: isSuccess,
		Data:    data,
		Stats:   stats,
	}
}
