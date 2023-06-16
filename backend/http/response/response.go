package response

type ResponseBody struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

var (
	MessageSuccess = "success"
	MessageError   = "error"
)
