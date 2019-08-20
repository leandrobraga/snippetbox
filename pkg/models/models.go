package models

import (
	"errors"
	"time"
)

// ErrNoRecord is a default message error
var (
	ErrNoRecord           = errors.New("sql: no rows in result set")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

// Snippet struct represent a snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User struct represent a user in system
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
