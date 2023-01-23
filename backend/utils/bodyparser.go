package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func Parse[T any](r *http.Request, body T) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(b), &body); err != nil {
		return err
	}
	return nil
}
