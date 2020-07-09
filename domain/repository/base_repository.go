package repository

import (
	"strconv"

	"github.com/gin-gonic/gin"
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
	PerPage int
	Page    int
	Total   interface{}
}

// NewParameters construct Parameters from request.
func NewParameters(c *gin.Context) *Parameters {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	order := c.DefaultQuery("order", "desc")
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
