package main

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Article struct {
	ID       int
	AuthorID int
	Title    string
	Content  string
	Created  time.Time
	Expires  time.Time
}
