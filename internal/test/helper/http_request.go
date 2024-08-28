package helper

import (
	"bytes"
	"net/http"
)

func HttpRequest(query string, url string, method string) (resp *http.Response, close func(), err error) {
	var req *http.Request
	req, err = http.NewRequest(method, url+"/query", bytes.NewBuffer([]byte(query)))

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	close = func() {
		req.Body.Close()
		resp.Body.Close()
	}

	return
}
