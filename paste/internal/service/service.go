package service

import (
	"pastebin/internal/entity"
)

type PasteRepo interface {
	CreatePaste(paste entity.Paste) error
}

type PasteService struct {
	repo PasteRepo
}

func NewPasteService(r PasteRepo) *PasteService {
	return &PasteService{
		repo: r,
	}
}

func (p PasteService) CreatePaste(paste entity.Paste) error {

	//

	// err := p.repo.CreatePaste()

	return nil
}
