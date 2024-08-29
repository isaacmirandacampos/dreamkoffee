package helper

type ErrorResponse struct {
	Errors []struct {
		Message    string `json:"message"`
		Extensions struct {
			Error      string `json:"error"`
			StatusCode int    `json:"status_code"`
		} `json:"extensions"`
	} `json:"errors"`
	Data interface{} `json:"data"`
}
