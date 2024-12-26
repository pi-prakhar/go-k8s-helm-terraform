package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler interface {
	HandleTest(w http.ResponseWriter, r *http.Request)
	HandleBase(w http.ResponseWriter, r *http.Request)
	HandleBaseTest(w http.ResponseWriter, r *http.Request)
	HandleBaseServer(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) HandleTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	time.Sleep(50 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "text/plain")
	json.NewEncoder(w).Encode("Hello World Test 2")
}

func (h *handler) HandleBase(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	time.Sleep(50 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "text/plain")
	json.NewEncoder(w).Encode("Hello World Base 2")
}

func (h *handler) HandleBaseServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	time.Sleep(50 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "text/plain")
	json.NewEncoder(w).Encode("Hello World Base Server 2")
}

func (h *handler) HandleBaseTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	time.Sleep(50 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "text/plain")
	json.NewEncoder(w).Encode("Hello World Base Test 2")
}
