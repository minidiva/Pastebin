package repo

import "auth/internal/entity"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) CreateUser(user entity.User) error {
	// query := ``

	return nil
}
