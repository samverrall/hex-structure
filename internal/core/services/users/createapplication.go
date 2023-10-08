package users

import (
	"context"
	"fmt"

	"github.com/samverrall/hex-structure/internal/core/domain/user"
)

type CreateAccountReq struct {
	Username string
}

type CreateAccountResp struct {
	UserID string
}

func (s *Service) CreateAccount(ctx context.Context, req CreateAccountReq) (*CreateAccountResp, error) {
	userName, err := user.NewUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid username supplied: %w", err)
	}

	user := user.New(userName)

	if err := s.userRepo.Add(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to add a user: %w", err)
	}

	return &CreateAccountResp{UserID: user.ID.String()}, nil
}
