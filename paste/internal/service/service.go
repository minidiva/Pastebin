package service

import (
	"context"
	"pastebin/internal/entity"
	"time"
)

type PasteRepo interface {
	CreatePaste(ctx context.Context, paste entity.Paste) error
}

type PasteService struct {
	repo PasteRepo
}

func NewPasteService(r PasteRepo) *PasteService {
	return &PasteService{
		repo: r,
	}
}

// TODO:
// Создаёт уникальный ключ для S3 (uuid или хэш)
// Отправляет текст в S3
// Сохраняет метаданные в базу (ID, S3 key aka Link, created_at, ownerID)

func (p PasteService) CreatePaste(
	ctx context.Context,
	text string,
	ttl time.Duration, // ttl - TimeToLive
) error {

	// сгенерировать S3 ключ (возможно это надо вынести)
	// запись в S3 "Вызов стораж метода"
	// запись в БД "Вызов БД метода" (если запись в S3 успешная)

	// err := p.repo.CreatePaste()

	return nil
}
