package handlers

import (
	"pastebin/internal/entity"
)

type PasteService interface {
	CreatePaste(user entity.Paste) error
}

type PasteHandler struct {
	service PasteService
}

func NewPasteHandler(s PasteService) *PasteHandler {
	return &PasteHandler{
		service: s,
	}
}
