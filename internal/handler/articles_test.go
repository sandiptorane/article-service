package handler

import (
	"article-service/internal/service"
	mocks "article-service/mocks/internal_/service"
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testID = "042549e2-d4de-4e9d-8452-d6030284421c"

func TestAddArticle_Success(t *testing.T) {
	type fields struct {
		DB service.IArticle
	}

	tests := []struct {
		name       string
		fields     fields
		body       []byte
		wantStatus int
	}{
		{
			name: "happy flow",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					mockDB.EXPECT().SaveArticle(mock.Anything).Return(nil)

					return mockDB
				}(),
			},
			body: []byte(`{
               "title": "sample title",
               "content" : "this is a content",
               "author": "test author"
              }`),
			wantStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				DB: tt.fields.DB,
			}

			r := gin.Default()
			r.POST("/articles", app.AddArticle)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/articles", bytes.NewReader(tt.body))
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())
		})
	}
}

func TestAddArticle_Failure(t *testing.T) {
	type fields struct {
		DB service.IArticle
	}

	tests := []struct {
		name       string
		fields     fields
		body       []byte
		wantStatus int
	}{
		{
			name: "validation error",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					return mockDB
				}(),
			},
			body: []byte(`{
               "title": "",
               "content" : "this is a content",
               "author": "test author"
              }`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "unmarshall error",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					return mockDB
				}(),
			},
			body: []byte(`{
               "title": "test title",
               "content" : 101,
               "author": "test author"
              }`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "failed to save article",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					mockDB.EXPECT().SaveArticle(mock.Anything).Return(errors.New("db error"))
					return mockDB
				}(),
			},
			body: []byte(`{
               "title": "test title",
               "content" : "test content",
               "author": "test author"
              }`),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				DB: tt.fields.DB,
			}

			r := gin.Default()
			r.POST("/articles", app.AddArticle)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/articles", bytes.NewReader(tt.body))
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())
		})
	}
}

func TestGetArticle_Success(t *testing.T) {
	type fields struct {
		DB service.IArticle
	}

	type args struct {
		articleId string
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name: "happy flow",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					mockDB.EXPECT().GetArticleByID(mock.Anything).Return(&service.Article{
						ID:      testID,
						Title:   "sample title",
						Author:  "this is a content",
						Content: "test author",
					}, nil)

					return mockDB
				}(),
			},
			args:       args{articleId: testID},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				DB: tt.fields.DB,
			}

			// register route
			r := gin.Default()
			r.GET("/articles/:article_id", app.GetArticle)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/articles/"+tt.args.articleId, nil)
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())
		})
	}
}

func TestGetArticle_Failure(t *testing.T) {
	type fields struct {
		DB service.IArticle
	}

	type args struct {
		articleId string
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name: "article not found",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					mockDB.EXPECT().GetArticleByID(mock.Anything).Return(nil, service.ErrRecordNotFound)

					return mockDB
				}(),
			},
			args:       args{articleId: testID},
			wantStatus: http.StatusNotFound,
		},
		{
			name: "error fetching article",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					mockDB.EXPECT().GetArticleByID(mock.Anything).Return(nil, errors.New("db error"))

					return mockDB
				}(),
			},
			args:       args{articleId: testID},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "empty id",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					return mockDB
				}(),
			},
			args:       args{articleId: " "},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				DB: tt.fields.DB,
			}

			// register route
			r := gin.Default()
			r.GET("/articles/:article_id", app.GetArticle)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/articles/"+tt.args.articleId, nil)
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())
		})
	}
}

func TestGetAllArticles_Success(t *testing.T) {
	type fields struct {
		DB service.IArticle
	}

	tests := []struct {
		name       string
		fields     fields
		wantStatus int
	}{
		{
			name: "happy flow",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					mockDB.EXPECT().GetAllArticles().Return([]*service.Article{{
						ID:      testID,
						Title:   "sample title",
						Author:  "this is a content",
						Content: "test author",
					},
					}, nil)

					return mockDB
				}(),
			},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				DB: tt.fields.DB,
			}

			// register route
			r := gin.Default()
			r.GET("/articles", app.GetAllArticles)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/articles", nil)
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())
		})
	}
}

func TestGetAllArticles_Failure(t *testing.T) {
	type fields struct {
		DB service.IArticle
	}

	tests := []struct {
		name       string
		fields     fields
		wantStatus int
	}{
		{
			name: "failed to fetch articles",
			fields: fields{
				DB: func() service.IArticle {
					mockDB := mocks.NewIArticle(t)
					mockDB.EXPECT().GetAllArticles().Return(nil, errors.New("db error"))

					return mockDB
				}(),
			},

			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				DB: tt.fields.DB,
			}

			// register route
			r := gin.Default()
			r.GET("/articles", app.GetAllArticles)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/articles", nil)
			if err != nil {
				t.Log("new request error ", err)
				return
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			t.Log(w.Body.String())
		})
	}
}
