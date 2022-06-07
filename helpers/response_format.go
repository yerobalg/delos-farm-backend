package helpers

type Response struct {
	Message string        `json:"message"`
	Success bool          `json:"success"`
	Data    interface{}   `json:"data"`
}

func ResponseFormat(
	msg string,
	isSuccess bool,
	data interface{},
) Response {
	return Response{
		Message: msg,
		Success: isSuccess,
		Data:    data,
	}
}
