package tour

import (
	"errors"
	"fmt"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/infrastructure/message/success"
	"go-rest-skeleton/pkg/response"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// Tours is a struct defines the dependencies that will be used.
type Tours struct {
	ur application.TourAppInterface
	rd authorization.AuthInterface
	tk authorization.TokenInterface
}

// NewTours is constructor will initialize Tour handler.
func NewTours(
	ur application.TourAppInterface,
	rd authorization.AuthInterface,
	tk authorization.TokenInterface) *Tours {
	return &Tours{
		ur: ur,
		rd: rd,
		tk: tk,
	}
}

// SaveTour is a function uses to handle create a new tour.
func (s *Tours) SaveTour(c *gin.Context) {
	var tourEntity entity.Tour
	if err := c.ShouldBindJSON(&tourEntity); err != nil {
		fmt.Println(err)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	validateErr := tourEntity.ValidateSaveTour()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	_, errDesc, errException := s.ur.GetTourBySlug(&tourEntity)
	if errException == nil {
		exceptionData := response.TranslateErrorForm(c, []response.ErrorForm{
			{
				Field: "slug",
				Msg:   exception.ErrorTextTourSlugAlreadyExists.Error(),
				Data: map[string]interface{}{
					"Slug": &tourEntity.Slug,
				},
			},
		})
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	newTour, errDesc, errException := s.ur.SaveTour(&tourEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextUnprocessableEntity) {
			_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	response.NewSuccess(c, newTour.DetailTour(), success.TourSuccessfullyCreateTour).JSON()
}

func (s *Tours) UpdateTour(c *gin.Context) {

}

func (s *Tours) DeleteTour(c *gin.Context) {

}

func (s *Tours) GetTours(c *gin.Context) {
	var tour entity.Tour
	var tours entity.Tours
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(tour.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	tours, meta, err := s.ur.GetTours(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, tours.DetailTours(), success.RoleSuccessfullyGetRoleList).WithMeta(meta).JSON()
}

func (s *Tours) GetTour(c *gin.Context) {
	var tourEntity entity.Role
	if err := c.ShouldBindUri(&tourEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	tour, err := s.ur.GetTour(UUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextUserNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, tour.DetailTour(), success.RoleSuccessfullyGetRoleDetail).JSON()
}
