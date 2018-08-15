package utils

import (
	"encoding/json"
	"net/http"
)

type Response map[string]interface{}

func JsonResponse(w http.ResponseWriter, code int, r Response) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(r)
}
