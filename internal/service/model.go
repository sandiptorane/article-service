package service

import (
	"errors"
)

// Article holds article details
type Article struct {
	ID      string
	Title   string
	Author  string
	Content string
}

// ErrRecordNotFound represent record not found in db
var ErrRecordNotFound = errors.New("record not found")
