package apiutility

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data any, status int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func Wrap(key string, value any) map[string]any {
	return map[string]any{key: value}
}
