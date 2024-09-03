package helper

type ResponseError []struct {
	Message    string `json:"message"`
	Extensions struct {
		Error      string `json:"error"`
		StatusCode int    `json:"status_code"`
	} `json:"extensions"`
}
