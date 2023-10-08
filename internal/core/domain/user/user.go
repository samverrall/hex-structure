package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username Username
}

func New(un Username) User {
	return User{
		ID:       uuid.New(),
		Username: un,
	}
}
