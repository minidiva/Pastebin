package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type PasteService interface {
	CreatePaste(ctx context.Context, text string, ttl time.Duration) error
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

// DONE:
// Принимает POST-запрос!! от клиента с текстом пасты
// Валидирует (не пустой)
// Передаёт текст дальше на слой сервиса

type CreatePasteRequest struct {
	Text string        `json:"text"`
	TTL  time.Duration `json:"ttl"`
}

func (h *PasteHandler) CreatePaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req CreatePasteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: работа с контекстом
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.CreatePaste(ctx, req.Text, req.TTL); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
