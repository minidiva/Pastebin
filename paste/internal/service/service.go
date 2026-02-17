package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"pastebin/internal/entity"
	"time"
)

var (
	ErrPasteNotFound = errors.New("paste not found")
	ErrPasteExpired  = errors.New("paste expired")
)

type PasteRepo interface {
	CreatePaste(ctx context.Context, paste entity.Paste) error
	GetPaste(ctx context.Context, key string) (*entity.Paste, error)
}

type PasteStorage interface {
	Upload(ctx context.Context, key string, data []byte) error
	Download(ctx context.Context, key string) ([]byte, error)
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
// Создаёт уникальный ключ для S3 (uuid или хэш) done
// Отправляет текст в S3 done
// Сохраняет метаданные в базу (ID, S3 key, created_at, ownerID) done

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
) (s3key string, err error) {

	key, err := generateKey()
	if err != nil {
		fmt.Printf("failed to generate key: %v\n", err)
		return "", fmt.Errorf("failed to generate key: %w", err)
	}

	// загрузка в storage
	if err := p.storage.Upload(ctx, key, []byte(text)); err != nil {
		fmt.Printf("failed to upload to storage: %v\n", err)
		return "", fmt.Errorf("failed to upload to storage: %w", err)
	}

	// создаём метаданные
	now := time.Now()
	paste := entity.Paste{
		Key:       key,
		ExpiresAt: now.Add(ttl),
	}

	// сохраняем в БД
	if err := p.repo.CreatePaste(ctx, paste); err != nil {
		fmt.Printf("failed to save paste metadata: %v\n", err)
		return "", fmt.Errorf("failed to save paste metadata: %w", err)
	}

	fmt.Printf("Paste created successfully: key=%s ttl=%v\n", key, ttl)
	return key, nil
}

func (s *PasteService) GetPaste(ctx context.Context, key string) (string, error) {

	paste, err := s.repo.GetPaste(ctx, key)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrPasteNotFound
		}
		return "", fmt.Errorf("repo get paste: %w", err)
	}

	if time.Now().After(paste.ExpiresAt) {
		return "", ErrPasteExpired
	}

	data, err := s.storage.Download(ctx, paste.Key)
	if err != nil {
		return "", fmt.Errorf("storage download: %w", err)
	}

	return string(data), nil
}
