package helper

import (
	"encoding/json"
	"fmt"
	"io"
)

func TransformBody(body io.ReadCloser, response interface{}) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("could not read HTTP response: %v", err)
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return fmt.Errorf("could not unmarshal JSON response: %v", err)
	}
	return nil
}
