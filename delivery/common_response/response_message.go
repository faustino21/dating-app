package common_response

type ResponseMessage struct {
	Status      string      `json:"status"`
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

type ErrorMessage struct {
	HttpCode         int
	ErrorDescription ErrorDescription
}

type ErrorDescription struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func NewResponseMessage(code, description string, data interface{}) *ResponseMessage {
	return &ResponseMessage{"success", code, description, data}
}

func NewErrorMessage(httpCode int, errCode string, message string) *ErrorMessage {
	return &ErrorMessage{
		HttpCode: httpCode,
		ErrorDescription: ErrorDescription{
			Code:        errCode,
			Description: message,
		},
	}
}
