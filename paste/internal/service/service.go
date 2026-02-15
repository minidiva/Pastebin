package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"pastebin/internal/entity"
	"time"
)

type PasteRepo interface {
	CreatePaste(ctx context.Context, paste entity.Paste) error
}

type PasteStorage interface {
	Upload(ctx context.Context, key string, data []byte) error
}

type PasteService struct {
	repo    PasteRepo
	storage PasteStorage
}

func NewPasteService(r PasteRepo, storage PasteStorage) *PasteService {
	return &PasteService{
		repo:    r,
		storage: storage,
	}
}

// TODO:
// Создаёт уникальный ключ для S3 (uuid или хэш)
// Отправляет текст в S3
// Сохраняет метаданные в базу (ID, S3 key aka Link, created_at, ownerID)

func generateKey() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b) + ".txt", nil
}
func (p PasteService) CreatePaste(
	ctx context.Context,
	text string,
	ttl time.Duration,
) error {
	fmt.Println("Service has started to work!")

	key, err := generateKey()
	if err != nil {
		fmt.Printf("failed to generate key: %v\n", err)
		return fmt.Errorf("failed to generate key: %w", err)
	}

	// 1. загрузка в storage
	if err := p.storage.Upload(ctx, key, []byte(text)); err != nil {
		fmt.Printf("failed to upload to storage: %v\n", err)
		return fmt.Errorf("failed to upload to storage: %w", err)
	}

	// 2. создаём метаданные
	now := time.Now()
	paste := entity.Paste{
		Key:       key,
		ExpiresAt: now.Add(ttl),
	}

	// 3. сохраняем в БД
	if err := p.repo.CreatePaste(ctx, paste); err != nil {
		fmt.Printf("failed to save paste metadata: %v\n", err)
		return fmt.Errorf("failed to save paste metadata: %w", err)
	}

	fmt.Printf("Paste created successfully: key=%s ttl=%v\n", key, ttl)
	return nil
}
