package router

import (
	"article-service/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers routes and handlers and return router
func RegisterRoutes(app *handler.Application) *gin.Engine {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	v1 := r.Group("/")

	v1.POST("articles", app.AddArticle)
	v1.GET("articles/:article_id", app.GetArticle)
	v1.GET("articles", app.GetAllArticles)

	return r
}
