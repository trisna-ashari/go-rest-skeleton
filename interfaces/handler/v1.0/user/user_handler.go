package userv1point00

import (
	"errors"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/infrastructure/message/success"
	"go-rest-skeleton/infrastructure/storage"
	"go-rest-skeleton/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	fileCategory = storage.CategoryAvatar
)

// Users is a struct defines the dependencies that will be used.
type Users struct {
	us application.UserAppInterface
	ss application.StorageAppInterface
}

// NewUsers is constructor will initialize user handler.
func NewUsers(us application.UserAppInterface, ss application.StorageAppInterface) *Users {
	return &Users{
		us: us,
		ss: ss,
	}
}

// @Summary Create a new user
// @Description Create a new user.
// @Tags users
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param name formData string true "User name"
// @Param password formData string true "User password"
// @Param email formData string true "User email"
// @Param phone formData string true "User phone"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/users [post]
// SaveUser is a function uses to handle create a new user.
func (s *Users) SaveUser(c *gin.Context) {
	var userEntity entity.User
	if err := c.ShouldBindJSON(&userEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	validateErr := userEntity.ValidateSaveUser()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newUser, errDesc, errException := s.us.SaveUser(&userEntity)
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
	response.NewSuccess(c, newUser.DetailUser(), success.UserSuccessfullyCreateUser).JSON()
}

// @Summary Update user
// @Description Update an existing user.
// @Tags users
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "User UUID"
// @Param name formData string true "User name"
// @Param password formData string true "User password"
// @Param email formData string true "User email"
// @Param phone formData string true "User phone"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/users/uuid [put]
// UpdateUser is a function uses to handle update user by UUID.
func (s *Users) UpdateUser(c *gin.Context) {
	var userEntity entity.User
	if err := c.ShouldBindUri(&userEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&userEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := userEntity.ValidateUpdateUser()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	UUID := c.Param("uuid")
	updatedUser, errDesc, errException := s.us.UpdateUser(UUID, &userEntity)
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
	c.Status(http.StatusOK)
	response.NewSuccess(c, updatedUser.DetailUser(), success.UserSuccessfullyUpdateUser).JSON()
}

// @Summary Delete user
// @Description Delete an existing user.
// @Tags users
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "User UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/users/{uuid} [delete]
// DeleteUser is a function uses to handle delete user by UUID.
func (s *Users) DeleteUser(c *gin.Context) {
	var userEntity entity.User
	if err := c.ShouldBindUri(&userEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteUser(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextUserNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextUserNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.UserSuccessfullyDeleteUser).JSON()
}

// @Summary Get users
// @Description Get list of existing users.
// @Tags users
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/users [get]
// GetUsers is a function uses to handle get user list.
func (s *Users) GetUsers(c *gin.Context) {
	var users entity.Users
	var err error
	parameters := repository.NewParameters(c)
	users, meta, err := s.us.GetUsers(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, users.DetailUsers(), success.UserSuccessfullyGetUserList).WithMeta(meta).JSON()
}

// @Summary Get user
// @Description Get detail of existing user.
// @Tags users
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "User UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/users/{uuid} [get]
// GetUser is a function uses to handle get user detail by UUID.
func (s *Users) GetUser(c *gin.Context) {
	var userEntity entity.User
	if err := c.ShouldBindUri(&userEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	user, err := s.us.GetUser(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextUserNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextUserNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	avatarURL, errAvatar := s.ss.GetFile(user.AvatarUUID)
	if errAvatar != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errAvatar)
		return
	}
	response.NewSuccess(c, user.DetailUserAvatar(avatarURL), success.UserSuccessfullyGetUserDetail).JSON()
}

// @Summary Update user avatar
// @Description Update user avatar of existing user.
// @Tags users
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "User UUID"
// @Param avatar formData file true "User Avatar"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/users/{uuid}/avatar [put]
// UpdateAvatar is a function uses to handle update user avatar by UUID.
func (s *Users) UpdateAvatar(c *gin.Context) {
	UUID, exists := c.Get("UUID")
	if !exists {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	var userEntity entity.User
	avatarData, _, errException, errArgs := s.ss.UploadFile(avatar, fileCategory)
	if errException != nil {
		if errors.Is(errException, exception.ErrorTextStorageUploadInvalidSize) {
			c.Set("args", errArgs)
		}
		if errors.Is(errException, exception.ErrorTextStorageUploadInvalidFileType) {
			c.Set("args", errArgs)
		}
		_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
		return
	}

	userEntity.AvatarUUID = avatarData
	updatedUser, errDesc, errException := s.us.UpdateUserAvatar(UUID.(string), &userEntity)
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
	avatarURL, errAvatar := s.ss.GetFile(updatedUser.AvatarUUID)
	if errAvatar != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errAvatar)
		return
	}
	c.Status(http.StatusOK)
	response.NewSuccess(c, updatedUser.DetailUserAvatar(avatarURL), success.UserSuccessfullyUpdateUserAvatar).JSON()
}
