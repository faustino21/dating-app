package commonResponse

var (
	responseSuccess = "00"
	responseFailed  = "XX"
	statusSuccess   = "Success"
	statusFailed    = "Failed"
)

type ResponseMessage struct {
	HttpResponse int         `json:"http_response"`
	RespCode     string      `json:"resp_code"`
	Status       string      `json:"status"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

type FailedMessage struct {
	HttpResponse int    `json:"http_response"`
	RespCode     string `json:"resp_code"`
	Status       string `json:"status"`
	Message      string `json:"message"`
}

func NewSuccessMessage(message string, data interface{}) *ResponseMessage {
	return &ResponseMessage{
		RespCode: responseSuccess,
		Status:   statusSuccess,
		Message:  message,
		Data:     data,
	}
}

func NewFailedMessage(message string) *FailedMessage {
	return &FailedMessage{
		RespCode: responseFailed,
		Status:   statusFailed,
		Message:  message,
	}
}
