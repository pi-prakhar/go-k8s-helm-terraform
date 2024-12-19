package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler interface {
	Test(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Test(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	time.Sleep(500 * time.Millisecond)

	json.NewEncoder(w).Encode("Hello World")
}
