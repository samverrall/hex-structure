package postgres

import (
	"context"
	"database/sql"

	"github.com/samverrall/hex-structure/internal/core/domain/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo() (*UserRepo, error) {
	return &UserRepo{}, nil
}

func (ur *UserRepo) Add(ctx context.Context, u user.User) error {
	return nil
}
