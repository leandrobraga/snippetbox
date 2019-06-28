package models

import (
	"errors"
	"time"
)

// ErrNoRecord is a default message error
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet struct represent a snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
