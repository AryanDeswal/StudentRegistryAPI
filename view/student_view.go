package view

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ReaderError(w http.ResponseWriter, statusCode int, message string) {
	RenderJSON(w, statusCode, map[string]string{"error": message})
}
