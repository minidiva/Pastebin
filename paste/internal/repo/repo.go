package repo

import "pastebin/internal/entity"

type PasteRepo struct{}

func NewRepo() *PasteRepo {
	return &PasteRepo{}
}

func (r PasteRepo) CreatePaste(paste entity.Paste) error {
	return nil
}
