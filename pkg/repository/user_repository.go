package repository

import (
	"context"

	"github.com/KoteiIto/wire-sample/pkg/model"
)

type IUserRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type UserRepository struct {
}

var _ IUserRepository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	// Do something ...
	return &model.User{UserId: 1, Email: email}, nil
}
