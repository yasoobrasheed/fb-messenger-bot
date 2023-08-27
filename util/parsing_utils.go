package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseAndUnmarshallRequestBody(r *http.Request, data interface{}) error {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(requestBody, data)
	if err != nil {
		log.Println("Error unmarshalling request body:", err)
		return err
	}

	return nil
}
