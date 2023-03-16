package service

// IArticle used by application code
type IArticle interface {
	SaveArticle(article *Article) (err error)
	GetArticleByID(id string) (*Article, error)
	GetAllArticles() ([]*Article, error)
}
