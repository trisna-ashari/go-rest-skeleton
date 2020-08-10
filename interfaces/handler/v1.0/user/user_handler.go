package userv1point00

import (
	"errors"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/exception"
	"go-rest-skeleton/infrastructure/storage"
	"go-rest-skeleton/interfaces/middleware"
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

// SaveUser is a function uses to handle create a new user.
func (s *Users) SaveUser(c *gin.Context) {
	var userEntity entity.User
	if err := c.ShouldBindJSON(&userEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	validateErr := userEntity.ValidateSaveUser(c)
	if len(validateErr) > 0 {
		c.Set("data", validateErr)
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
	middleware.Formatter(c, newUser.DetailUser(), "api.msg.success.successfully_create_user", nil)
}

// UpdateUser is a function uses to handle create a new user.
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

	validateErr := userEntity.ValidateUpdateUser(c)
	if len(validateErr) > 0 {
		c.Set("data", validateErr)
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
	middleware.Formatter(c, updatedUser.DetailUser(), "api.msg.success.successfully_update_user", nil)
}

// DeleteUser is a function uses to handle get user detail by UUID.
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
	middleware.Formatter(c, nil, "api.msg.success.successfully_delete_user", nil)
}

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
	middleware.Formatter(c, users.DetailUsers(), "api.msg.success.successfully_get_user_list", meta)
}

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
	middleware.Formatter(c, user.DetailUserAvatar(avatarURL), "api.msg.success.successfully_get_user_detail", nil)
}

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
	avatarData, _, errException := s.ss.UploadFile(avatar, fileCategory)
	if errException != nil {
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
	middleware.Formatter(c, updatedUser.DetailUserAvatar(avatarURL), "api.msg.success.successfully_update_avatar", nil)
}
