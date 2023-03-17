package router

import (
	"article-service/internal/handler"
	"article-service/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	type args struct {
		app *handler.Application
	}
	tests := []struct {
		name string
		args args
	}{{
		name: "Happy path",
		args: args{
			app: handler.New(&service.ArticleStore{}),
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RegisterRoutes(tt.args.app)
			assert.NotNil(t, got)
		})
	}
}
