package handler

import (
	"article-service/internal/service"
)

// Application holds db and handler methods
type Application struct {
	DB service.IArticle
}

// New returns the new instance of Application
func New(db service.IArticle) *Application {
	return &Application{DB: db}
}
