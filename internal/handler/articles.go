package handler

import (
	"article-service/internal/model"
	"article-service/internal/service"
	"article-service/pkg/response"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// AddArticle create article and return the id in response
func (app *Application) AddArticle(c *gin.Context) {
	var req model.POSTArticleReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Warn("unmarshall error : ", err)

		response.BadRequest(c, err.Error(), nil)

		return
	}

	s := validator.New()
	err = s.Struct(removeWhiteSpaces(&req))
	if err != nil {
		log.Warn("unmarshall or validation error : ", err)

		response.BadRequest(c, err.Error(), nil)

		return
	}

	data := &service.Article{
		ID:      uuid.New().String(),
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
	}

	err = app.DB.SaveArticle(data)
	if err != nil {
		log.Error("SaveArticle db error : ", err)

		response.InternalServerError(c, "failed to add article", nil)

		return
	}

	response.Created(c, "Success", &model.POSTArticleResp{ID: data.ID})
}

func removeWhiteSpaces(req *model.POSTArticleReq) *model.POSTArticleReq {
	req.Title = strings.TrimSpace(req.Title)
	req.Content = strings.TrimSpace(req.Content)
	req.Author = strings.TrimSpace(req.Author)

	return req
}

// GetArticle to fetch single article details using id
func (app *Application) GetArticle(c *gin.Context) {
	// get id  from params
	id := strings.TrimSpace(c.Param("article_id"))

	if id == "" {
		log.Warn("id is empty")
		response.BadRequest(c, "article_id should not be empty", nil)

		return
	}

	data, err := app.DB.GetArticleByID(id)
	if err != nil {
		log.Error("error fetching article by id err: ", err, "id: ", id)

		if errors.Is(err, service.ErrRecordNotFound) {
			response.NotFound(c, "article not found", nil)
			return
		}

		response.InternalServerError(c, "failed to fetch article", nil)

		return
	}

	resp := &model.Article{
		ID:      data.ID,
		Title:   data.Title,
		Content: data.Content,
		Author:  data.Author,
	}

	response.Success(c, "Success", resp)
}

// GetAllArticles fetches all articles
func (app *Application) GetAllArticles(c *gin.Context) {
	data, err := app.DB.GetAllArticles()
	if err != nil {
		log.Error("GetAllArticles db error: ", err)

		response.InternalServerError(c, "failed to fetch all articles", nil)

		return
	}

	// prepare resp
	var resp []*model.Article

	for _, d := range data {
		resp = append(resp, &model.Article{
			ID:      d.ID,
			Title:   d.Title,
			Content: d.Content,
			Author:  d.Author,
		})
	}

	response.Success(c, "Success", resp)
}
