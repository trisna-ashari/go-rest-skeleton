package repository

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Parameters struct {
	Offset  int
	Limit   int
	PerPage int
	Page    int
	Order   string
}

type Meta struct {
	PerPage int
	Page	int
	Total	interface{}
}

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

func NewMeta(p *Parameters, total int) *Meta {
	return &Meta{
		Page: p.Page,
		PerPage: p.PerPage,
		Total: total,
	}
}
