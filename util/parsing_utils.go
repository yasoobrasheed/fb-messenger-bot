package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseAndUnmarshallRequestBody(r *http.Request, data interface{}) error {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(requestBody, data)
	if err != nil {
		return err
	}

	return nil
}
