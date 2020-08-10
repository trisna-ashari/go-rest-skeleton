package repository

import (
	"fmt"
	"go-rest-skeleton/infrastructure/exception"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultPage    = 1
	defaultPerPage = 5
	maxPerPage     = 25
)

// Parameters represent it self.
type Parameters struct {
	Offset  int
	Limit   int
	PerPage int
	Page    int
	Order   string
}

// Meta represent it self.
type Meta struct {
	PerPage int         `json:"per_page"`
	Page    int         `json:"page"`
	Total   interface{} `json:"total"`
}

// NewParameters construct Parameters from request.
func NewParameters(c *gin.Context) *Parameters {
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", strconv.Itoa(defaultPerPage)))
	order := c.DefaultQuery("order", "desc")

	if perPage > maxPerPage {
		c.Set("args", fmt.Sprintf("Max:%d", maxPerPage))
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextPerPage)
	}

	var offset int
	var limit int
	if page <= 1 {
		offset = 0
		limit = perPage
	}
	if page > 1 {
		offset = (page - 1) * perPage
		limit = perPage
	}

	return &Parameters{
		Offset:  offset,
		Limit:   limit,
		PerPage: perPage,
		Page:    page,
		Order:   order,
	}
}

// NewMeta construct of metadata for response.
func NewMeta(p *Parameters, total int) *Meta {
	return &Meta{
		Page:    p.Page,
		PerPage: p.PerPage,
		Total:   total,
	}
}
