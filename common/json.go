package common

import (
	"encoding/json"
	"net/http"
)

// to marshall/unmarshall json from gateway
func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ReadJSON(r *http.Request, data any) error {
	// Decoder reads from request body (raw) and map to `data`
	//(`data` --> usually in struct or map that matches the JSON structure)
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}
