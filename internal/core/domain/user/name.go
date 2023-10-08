package user

import (
	"errors"
	"strings"
)

var (
	ErrEmptyUsername = errors.New("empty username supplied")
)

type Username string

func NewUsername(un string) (Username, error) {
	un = strings.TrimSpace(un)
	if un == "" {
		return "", ErrEmptyUsername
	}

	// Must be less than 15 chars

	// Validate it is a valid username.
	return Username(un), nil
}
