package model

// POSTArticleReq holds the request details for AddArticles
type POSTArticleReq struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author,omitempty" validate:"required"`
}

// POSTArticleResp holds response details of AddArticles
type POSTArticleResp struct {
	ID string `json:"id"`
}

// Article holds article details
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
