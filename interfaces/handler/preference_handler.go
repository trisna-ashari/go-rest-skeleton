package handler

import (
	"errors"
	"fmt"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/infrastructure/message/success"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Preference is a struct.
type Preference struct {
	up application.UserPreferenceAppInterface
	rd authorization.AuthInterface
	tk authorization.TokenInterface
}

// NewPreference will initialize interface for handler.Preference handler.
func NewPreference(
	up application.UserPreferenceAppInterface,
	rd authorization.AuthInterface,
	tk authorization.TokenInterface) *Preference {
	return &Preference{
		up: up,
		rd: rd,
		tk: tk,
	}
}

// @Summary User preference
// @Description Get current user preference using Authorization Header.
// @Tags preference
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 401 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
// @Router /api/v1/external/preference [get]
// GerPreference will return detail user preference of current logged in user.
func (up *Preference) GerPreference(c *gin.Context) {
	UUID, exists := c.Get("UUID")
	if !exists {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	userPreferenceData, err := up.up.GetUserPreference(UUID.(string))
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(userPreferenceData)

	userPreference := userPreferenceData.DetailUserPreference()
	middleware.Formatter(c, userPreference, success.UserSuccessfullyGetUserPreference, nil)
}

// @Summary Update user preference
// @Description Update current user preference using Authorization Header.
// @Tags preference
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Param preference body entity.DetailUserPreference true "User preference"
// @Security BasicAuth
// @Security JWTAuth
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 401 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 422 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
// @Router /api/v1/external/preference [put]
// UpdatePreference will update user preference of current logged in user.
func (up *Preference) UpdatePreference(c *gin.Context) {
	UUID, exists := c.Get("UUID")
	if !exists {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	var preference entity.DetailUserPreference
	if err := c.ShouldBind(&preference); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	userPreferenceData, errDesc, errException := up.up.UpdateUserPreference(UUID.(string), &preference)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextUserNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, errException)
			return
		}
		if errors.Is(errException, exception.ErrorTextUnprocessableEntity) {
			_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
		return
	}

	fmt.Println(userPreferenceData)

	userPreference := userPreferenceData.DetailUserPreference()
	middleware.Formatter(c, userPreference, success.UserSuccessfullyUpdateUserPreference, nil)
}

// @Summary Reset user preference
// @Description Reset current user preference to default using Authorization Header.
// @Tags preference
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 401 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
// @Router /api/v1/external/preference/reset [post]
// GerPreference will return detail user preference of current logged in user.
func (up *Preference) ResetPreference(c *gin.Context) {
	UUID, exists := c.Get("UUID")
	if !exists {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	userPreferenceData, err := up.up.ResetUserPreference(UUID.(string))
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(userPreferenceData)

	userPreference := userPreferenceData.DetailUserPreference()
	middleware.Formatter(c, userPreference, success.UserSuccessfullyResetUserPreference, nil)
}
