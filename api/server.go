package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Print(err)
	}
}
