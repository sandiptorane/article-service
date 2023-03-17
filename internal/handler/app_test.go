package handler

import (
	"article-service/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		db service.IArticle
	}
	var tests = []struct {
		name string
		args args
	}{
		{
			name: "Happy flow",
			args: args{
				db: &service.ArticleStore{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := New(tt.args.db)

			assert.NotNil(t, app)
			assert.NotNil(t, app.DB)
		})
	}
}
