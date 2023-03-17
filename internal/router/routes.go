package router

import (
	"article-service/internal/handler"
	"article-service/pkg/response"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers routes and handlers and return router
func RegisterRoutes(app *handler.Application) *gin.Engine {
	r := gin.Default()

	r.NoRoute(NoRoute)

	r.POST("/articles", app.AddArticle)
	r.GET("/articles/:article_id", app.GetArticle)
	r.GET("/articles", app.GetAllArticles)

	r.GET("/health", HealthCheck)

	return r
}

func NoRoute(c *gin.Context) {
	response.NotFound(c, "route not found", nil)
}

func HealthCheck(c *gin.Context) {
	response.Success(c, "ok", nil)
}
