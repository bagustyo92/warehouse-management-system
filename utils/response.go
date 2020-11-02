package utils

type Body struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Payload    interface{} `json:"payload"`
}

// Responsse construct ManyOption's response payload format
func Response(statusCode int, message interface{}, payload interface{}) (int, *Body) {
	err, ok := message.(error)
	if ok {
		message = err.Error()
	}

	return statusCode, &Body{
		statusCode,
		message,
		payload,
	}
}
