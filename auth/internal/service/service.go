package service

import (
	"auth/internal/entity"
)

type UserRepo interface {
	CreateUser(user entity.User) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateUser(user entity.User) error {
	err := s.repo.CreateUser(user)

	if err != nil {

	}
	return nil
}
