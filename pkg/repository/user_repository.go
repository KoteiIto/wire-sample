package repository

import (
	"context"
	"database/sql"

	"github.com/KoteiIto/wire-sample/pkg/model"
)

type IUserRepository interface {
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int64) (*model.User, error)
}

type UserRepository struct {

}

var _ IUserRepository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByUserId(ctx context.Context, tx *sql.Tx, userId int64) (*model.User, error) {
	// Do something ...
	return &model.User{UserId: userId}, nil
}
