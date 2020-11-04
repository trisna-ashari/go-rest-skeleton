package repository

import (
	"fmt"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/pkg/response"
	"go-rest-skeleton/pkg/validator"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	defaultPage    = 1
	defaultPerPage = 5
	maxPerPage     = 25

	exact   = "EXACT"
	similar = "SIMILAR"
	and     = "AND"
	or      = "OR"
	latest  = "LATEST"
	oldest  = "OLDEST"
)

type QueryParameters struct {
	PerPage         int
	Page            int
	OrderBy         string
	OrderMethod     string
	SearchCondition string
	Equals          map[int]map[string]interface{}
	Likes           map[int]map[string]interface{}
	NotEquals       map[int]map[string]interface{}
}

// Parameters represent it self.
type Parameters struct {
	Offset          int
	Limit           int
	PerPage         int
	Page            int
	Order           string
	QueryMethod     string
	QueryCondition  string
	QueryKey        string
	QueryValue      []interface{}
	QueryParameters *QueryParameters
}

// Meta represent it self.
type Meta struct {
	PerPage int   `json:"per_page"`
	Page    int   `json:"page"`
	Total   int64 `json:"total"`
}

// GqlParameters represent it self.
type GqlParameters struct {
	SearchMethod    string
	SearchCondition string
	SearchKeywords  map[string]interface{}
	Order           string
	Page            int
	PerPage         int
}

func (p *Parameters) ValidateParameter(fields ...interface{}) []response.ErrorForm {
	var qp = p.QueryParameters

	validation := validator.New()
	validation.
		Set("per_page", qp.PerPage, validation.AddRule().MaxValue(maxPerPage).Apply()).
		Set("page", qp.Page, validation.AddRule().MinValue(1).Apply()).
		Set("order_by", qp.OrderBy, validation.AddRule().Required().Apply()).
		Set("order_method", qp.OrderMethod, validation.AddRule().Required().In("asc", "desc").Apply()).
		Set("search_condition", qp.SearchCondition, validation.AddRule().Required().In("and", "or").Apply())

	for _, querySlice := range qp.Equals {
		for key, value := range querySlice {
			validation.
				Set("equal", key, validation.AddRule().IsAlpha().In(fields...).Apply()).
				Set(fmt.Sprintf("equal[%s]", key), value, validation.AddRule().IsAlphaNumericSpaceAndSpecialCharacter().Apply())
		}
	}

	for _, querySlice := range qp.Likes {
		for key, value := range querySlice {
			validation.
				Set("like", key, validation.AddRule().IsAlpha().In(fields...).Apply()).
				Set(fmt.Sprintf("like[%s]", key), value, validation.AddRule().IsAlphaNumericSpaceAndSpecialCharacter().Apply())
		}
	}

	for _, querySlice := range qp.NotEquals {
		for key, value := range querySlice {
			validation.
				Set("not", key, validation.AddRule().IsAlpha().In(fields...).Apply()).
				Set(fmt.Sprintf("not[%s]", key), value, validation.AddRule().IsAlphaNumericSpaceAndSpecialCharacter().Apply())
		}
	}

	return validation.Validate()
}

// NewGinParameters construct Parameters from request.
func NewGinParameters(c *gin.Context) *Parameters {
	searchCondition := c.DefaultQuery("search_condition", "and")
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", strconv.Itoa(defaultPerPage)))
	orderBy := c.DefaultQuery("order_by", "created_at")
	orderMethod := c.DefaultQuery("order_method", "desc")

	var queryCondition string
	var queryKey []string
	var queryValue []interface{}
	var queryEqual = make(map[int]map[string]interface{})
	var queryNotEqual = make(map[int]map[string]interface{})
	var queryLike = make(map[int]map[string]interface{})

	if strings.ToLower(searchCondition) == strings.ToLower(and) {
		queryCondition = " AND "
	}

	if strings.ToLower(searchCondition) == strings.ToLower(or) {
		queryCondition = " OR "
	}

	for key, valueSlice := range c.Request.URL.Query() {
		reEqual, _ := regexp.Compile("equal\\[(.*[a-z])\\]")
		reEqualSlice := reEqual.FindStringSubmatch(key)
		if len(reEqualSlice) > 0 {
			if len(valueSlice) > 1 {
				for i, value := range valueSlice {
					queryEqual[i] = map[string]interface{}{reEqualSlice[1]: value}
				}
			} else {
				queryEqual[0] = map[string]interface{}{reEqualSlice[1]: strings.Join(valueSlice, "")}
			}
		}

		reNotEqual, _ := regexp.Compile("not\\[(.*[a-z])\\]")
		reNotEqualSlice := reNotEqual.FindStringSubmatch(key)
		if len(reNotEqualSlice) > 0 {
			if len(valueSlice) > 1 {
				for i, value := range valueSlice {
					queryNotEqual[i] = map[string]interface{}{reNotEqualSlice[1]: value}
				}
			} else {
				queryNotEqual[0] = map[string]interface{}{reNotEqualSlice[1]: strings.Join(valueSlice, "")}
			}
		}

		reLike, _ := regexp.Compile("like\\[(.*[a-z])\\]")
		reLikeSlice := reLike.FindStringSubmatch(key)
		if len(reLikeSlice) > 0 {
			if len(valueSlice) > 1 {
				for i, value := range valueSlice {
					queryLike[i] = map[string]interface{}{reLikeSlice[1]: value}
				}
			} else {
				queryLike[0] = map[string]interface{}{reLikeSlice[1]: strings.Join(valueSlice, "")}
			}
		}
	}

	for _, querySlice := range queryEqual {
		for key, value := range querySlice {
			if value != "" {
				queryKey = append(queryKey, key+" = ?")
				queryValue = append(queryValue, value)
			}
		}
	}

	for _, querySlice := range queryNotEqual {
		for key, value := range querySlice {
			if value != "" {
				queryKey = append(queryKey, key+" != ?")
				queryValue = append(queryValue, value)
			}
		}
	}

	for _, querySlice := range queryLike {
		for key, value := range querySlice {
			if value != "" {
				value = "%" + value.(string) + "%"
				queryKey = append(queryKey, key+" LIKE ?")
				queryValue = append(queryValue, value)
			}
		}
	}

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

	var queryParameters = &QueryParameters{
		PerPage:         perPage,
		Page:            page,
		OrderBy:         orderBy,
		OrderMethod:     orderMethod,
		SearchCondition: searchCondition,
		Equals:          queryEqual,
		Likes:           queryLike,
		NotEquals:       queryNotEqual,
	}

	return &Parameters{
		Offset:          offset,
		Limit:           limit,
		PerPage:         perPage,
		Page:            page,
		Order:           orderBy + " " + orderMethod,
		QueryKey:        strings.Join(queryKey, queryCondition),
		QueryValue:      queryValue,
		QueryParameters: queryParameters,
	}
}

// NewGqlParameters construct Parameters from request.
func NewGqlParameters(c *gin.Context, gqlParameters *GqlParameters) *Parameters {
	var search = gqlParameters.SearchKeywords
	var searchMethod = gqlParameters.SearchMethod
	var searchCondition = gqlParameters.SearchCondition
	var order = gqlParameters.Order
	var orderBy = "created_at"
	var orderMethod = "desc"
	var page = gqlParameters.Page
	var perPage = gqlParameters.PerPage
	var queryMethod string
	var queryCondition string
	var queryKey []string
	var queryValue []interface{}

	if searchMethod == exact {
		queryMethod = " = ?"
	}

	if searchMethod == similar {
		queryMethod = " LIKE ?"
	}

	if searchCondition == and {
		queryCondition = " AND "
	}

	if searchCondition == or {
		queryCondition = " OR "
	}

	if order == latest {
		orderMethod = "desc"
	}

	if order == oldest {
		orderMethod = "asc"
	}

	if perPage > maxPerPage {
		c.Set("args", fmt.Sprintf("Max:%d", maxPerPage))
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextPerPage)
	}

	if len(search) > 0 {
		for key, value := range search {
			if value != "" {
				if searchMethod == similar {
					value = "%" + value.(string) + "%"
				}

				queryKey = append(queryKey, key+queryMethod)
				queryValue = append(queryValue, value)
			}
		}
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
		Offset:     offset,
		Limit:      limit,
		PerPage:    perPage,
		Page:       page,
		Order:      orderBy + " " + orderMethod,
		QueryKey:   strings.Join(queryKey, queryCondition),
		QueryValue: queryValue,
	}
}

// NewMeta construct of metadata for response.
func NewMeta(p *Parameters, total int64) *Meta {
	return &Meta{
		Page:    p.Page,
		PerPage: p.PerPage,
		Total:   total,
	}
}
