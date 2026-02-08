package handlers

import (
	"net/http"
	"pastebin/internal/entity"
)

type PasteService interface {
	CreatePaste(paste entity.Paste) error
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

func (h PasteHandler) CreatePaste(w http.ResponseWriter, r *http.Request) error {

	//

	//err := h.service.CreatePaste()

	return nil

}
