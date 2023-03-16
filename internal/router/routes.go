package router

import (
	"article-service/internal/handler"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// RegisterRoutes registers routes and handlers and return router
func RegisterRoutes(app *handler.Application) *gin.Engine {
	r := gin.Default()
	r.Use(Logger())

	r.HandleMethodNotAllowed = true

	v1 := r.Group("/")

	v1.POST("articles", app.AddArticle)
	v1.GET("articles/:article_id", app.GetArticle)
	v1.GET("articles", app.GetAllArticles)

	return r
}

// Logger is middleware which
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		// after request
		latency := time.Since(t)
		log.Info("PATH :", c.FullPath(), "API latency:", latency)
	}
}
