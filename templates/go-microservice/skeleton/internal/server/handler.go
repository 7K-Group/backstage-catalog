package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type infoResponse struct {
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHandler(serviceName string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(infoResponse{
			Name:      serviceName,
			Status:    "running",
			Timestamp: time.Now().UTC(),
		})
	})

	return mux
}
