package service

import (
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"
)

// ArticleStore holds db connection and implement db functions
type ArticleStore struct {
	db *sql.DB
}

// GetInstance accepts db connection return ArticleStore instance
func GetInstance(db *sql.DB) *ArticleStore {
	return &ArticleStore{db: db}
}

// SaveArticle insert the record and return error
func (a *ArticleStore) SaveArticle(article *Article) (err error) {
	query := `INSERT INTO articles(id,title,author,content) VALUES (?,?,?,?)`

	_, err = a.db.Exec(query, article.ID, article.Title, article.Author, article.Content)

	return err
}

// GetArticleByID returns article by id or error if any
func (a *ArticleStore) GetArticleByID(id string) (*Article, error) {
	var v Article

	query := `SELECT id,title,COALESCE(author,'') as author, COALESCE(content,'') as content
      		  FROM articles 
      		  WHERE id=?`

	row := a.db.QueryRow(query, id)

	err := row.Scan(&v.ID, &v.Title, &v.Author, &v.Content)
	if err != nil {
		log.Error("error fetching article by id", "err", err, "id", id)

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}

		return nil, err
	}

	return &v, nil
}

// GetAllArticles fetches all articles
func (a *ArticleStore) GetAllArticles() ([]*Article, error) {
	var data []*Article

	query := `SELECT id,title,COALESCE(author,'') as author, COALESCE(content,'') as content
      		  FROM articles`

	rows, err := a.db.Query(query)
	if err != nil {
		log.Error("error GetAllArticles ", "err: ", err)
		return nil, err
	}

	for rows.Next() {
		var v Article

		err = rows.Scan(&v.ID, &v.Title, &v.Author, &v.Content)
		if err != nil {
			log.Error("error in GetAllArticles row scan", "err: ", err)
			return nil, err
		}

		data = append(data, &v)
	}

	return data, nil
}
