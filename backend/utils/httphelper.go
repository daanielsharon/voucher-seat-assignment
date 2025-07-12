package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func getResponseStatus(statusCode int) string {
	if statusCode >= 200 && statusCode < 300 {
		return "success"
	}

	return "failed"
}

func JSONRequest(w http.ResponseWriter, r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

func JSONResponse[T any](w http.ResponseWriter, statusCode int, message string, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(APIResponse[T]{
		Status:  getResponseStatus(statusCode),
		Message: message,
		Data:    data,
	})
}

func JSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(APIResponse[any]{
		Status:  getResponseStatus(statusCode),
		Message: message,
		Data:    nil,
	})
}
