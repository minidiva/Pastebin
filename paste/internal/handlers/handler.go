package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"pastebin/internal/service"
	"time"
)

type PasteService interface {
	CreatePaste(ctx context.Context, text string, ttl time.Duration) (string, error)
	GetPaste(ctx context.Context, key string) (string, error)
}

type PasteHandler struct {
	service PasteService
}

func NewPasteHandler(s PasteService) *PasteHandler {
	return &PasteHandler{
		service: s,
	}
}

func (h *PasteHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"healthy"}`))
}

type CreatePasteRequest struct {
	Text string `json:"text"`
	TTL  string `json:"ttl"`
}

type CreatePasteResponse struct {
	Key string `json:"key"`
}

func (h *PasteHandler) CreatePaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Wrong Method")
		return
	}

	var req CreatePasteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad request")
		return
	}

	duration, err := time.ParseDuration(req.TTL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing TTL")
	}

	if req.Text == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing text")
		return
	}

	// TODO: работа с контекстом
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	key, err := h.service.CreatePaste(ctx, req.Text, duration)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := CreatePasteResponse{
		Key: key,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *PasteHandler) GetPaste(w http.ResponseWriter, r *http.Request) {
	keyStr := r.URL.Query().Get("key")
	if keyStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing key")
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Wrong Method")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	text, err := h.service.GetPaste(ctx, keyStr)
	if err != nil {

		switch {
		case errors.Is(err, service.ErrPasteNotFound):
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Paste Not Found!")

		case errors.Is(err, service.ErrPasteExpired):
			w.WriteHeader(http.StatusGone)
			fmt.Fprint(w, "Paste Expired!")

		default:
			log.Printf("internal error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(text))
}
