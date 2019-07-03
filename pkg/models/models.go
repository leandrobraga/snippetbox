package models

import (
	"errors"
	"time"
)

// ErrNoRecord is a default message error
var ErrNoRecord = errors.New("sql: no rows in result set")

// Snippet struct represent a snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
