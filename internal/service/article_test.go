package service

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var testID = "042549e2-d4de-4e9d-8452-d6030284421c"

func TestSaveArticle(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		article *Article
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		mockQuery func(mock sqlmock.Sqlmock)
	}{
		{
			name:   "Happy flow",
			fields: fields{},
			args: args{
				article: &Article{
					ID:      testID,
					Title:   "test title",
					Author:  "test author",
					Content: "test content",
				},
			},
			wantErr: false,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO articles").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.mockQuery(mock)

			a := &ArticleStore{
				db: db,
			}

			err = a.SaveArticle(tt.args.article)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestGetInstance(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *ArticleStore
	}{
		{
			name: "Happy flow",
			args: args{
				db: &sql.DB{},
			},
			want: &ArticleStore{db: &sql.DB{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetInstance(tt.args.db), "GetInstance(%v)", tt.args.db)
		})
	}
}

func TestGetArticle(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id string
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		mockQuery func(mock sqlmock.Sqlmock)
	}{
		{
			name:   "Happy flow",
			fields: fields{},
			args: args{
				id: testID,
			},
			wantErr: false,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id,title,COALESCE(author,'') as author, COALESCE(content,'') as content FROM articles WHERE id=?")).WithArgs(testID).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "content"}).AddRow(testID, "test title", "", ""))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.mockQuery(mock)

			a := &ArticleStore{
				db: db,
			}

			d, err := a.GetArticleByID(tt.args.id)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, d)
		})
	}
}

func TestGetArticleFailure(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id string
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		mockQuery func(mock sqlmock.Sqlmock)
	}{
		{
			name:   "not found",
			fields: fields{},
			args: args{
				id: testID,
			},
			wantErr: true,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id,title,COALESCE(author,'') as author, COALESCE(content,'') as content FROM articles WHERE id=?")).WithArgs(testID).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "content"}))
			},
		},
		{
			name:   "other error",
			fields: fields{},
			args: args{
				id: testID,
			},
			wantErr: true,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id,title,COALESCE(author,'') as author, COALESCE(content,'') as content FROM articles WHERE id=?")).WithArgs(testID).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "content"}).AddRow(testID, "test title", "", "").RowError(0, errors.New("db error")))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.mockQuery(mock)

			a := &ArticleStore{
				db: db,
			}

			d, err := a.GetArticleByID(tt.args.id)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, d)
		})
	}
}

func TestGetAllArticles(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	tests := []struct {
		name      string
		fields    fields
		wantErr   bool
		mockQuery func(mock sqlmock.Sqlmock)
	}{
		{
			name:    "Happy flow",
			fields:  fields{},
			wantErr: false,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id,title,COALESCE(author,'') as author, COALESCE(content,'') as content FROM articles")).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "content"}).AddRow(testID, "test title", "", ""))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.mockQuery(mock)

			a := &ArticleStore{
				db: db,
			}

			d, err := a.GetAllArticles()
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, d)
		})
	}
}

func TestGetAllArticlesFailure(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	tests := []struct {
		name      string
		fields    fields
		wantErr   bool
		mockQuery func(mock sqlmock.Sqlmock)
	}{
		{
			name:    "query error",
			fields:  fields{},
			wantErr: true,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id,title,COALESCE(author,'') as author, COALESCE(content,'') as content FROM articles")).WillReturnError(errors.New("invalid syntax"))
			},
		},
		{
			name:    "row scan error",
			fields:  fields{},
			wantErr: true,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id,title,COALESCE(author,'') as author, COALESCE(content,'') as content FROM articles")).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "content"}).AddRow("test", "test", 1001, nil))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.mockQuery(mock)

			a := &ArticleStore{
				db: db,
			}

			d, err := a.GetAllArticles()
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, d)
		})
	}
}
