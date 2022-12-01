package helper

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors"`
}

func NewResponse(statusCode int, message string, data interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func NewErrorResponse(statusCode int, message string, errors interface{}) ErrorResponse {
	return ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	}
}
