package handler

import (
	"errors"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/infrastructure/message/success"
	"go-rest-skeleton/infrastructure/notify"
	"go-rest-skeleton/infrastructure/notify/notification"
	"go-rest-skeleton/interfaces/service"
	"go-rest-skeleton/pkg/response"
	"go-rest-skeleton/pkg/security"
	"go-rest-skeleton/pkg/translation"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authenticate is a struct to store interfaces.
type Authenticate struct {
	aa service.AuthService
	rd authorization.AuthInterface
	tk authorization.TokenInterface
	ni application.NotifyAppInterface
}

// NewAuthenticate will initialize authenticate interface for login handler.
func NewAuthenticate(
	aa service.AuthService,
	rd authorization.AuthInterface,
	tk authorization.TokenInterface,
	ni application.NotifyAppInterface) *Authenticate {
	return &Authenticate{
		aa: aa,
		rd: rd,
		tk: tk,
		ni: ni,
	}
}

// @Summary Authentication profile
// @Description Get current user profile using Authorization Header.
// @Tags authentication
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Security OauthGrantTypePassword
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/profile [get]
// Profile will return detail user of current logged in user.
func (au *Authenticate) Profile(c *gin.Context) {
	UUID, exists := c.Get("UUID")
	if !exists {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	userData, _ := au.aa.GetUserWithRoles(UUID.(string))

	response.NewSuccess(c, userData.DetailUser(), success.AuthSuccessfullyGetProfile).JSON()
}

// @Summary Authentication login
// @Description Login by email and password.
// @Tags authentication
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Param email formData string true "User email"
// @Param password formData string true "User password"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 422 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/login [post]
// Login will handle login request.
func (au *Authenticate) Login(c *gin.Context) {
	var user *entity.User
	var errToken = map[string]string{}

	if err := c.ShouldBind(&user); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := user.ValidateLogin()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	u, _, errException := au.aa.GetUserByEmailAndPassword(user)
	if errException != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, errException)
		return
	}

	ts, tErr := au.tk.CreateToken(u.UUID, "")
	if tErr != nil {
		errToken["token_error"] = tErr.Error()
		_ = c.AbortWithError(http.StatusUnprocessableEntity, tErr)
		return
	}

	errSave := au.rd.CreateAuth(u.UUID, ts)
	if errSave != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errSave)
		return
	}

	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken

	response.NewSuccess(c, userData, success.AuthSuccessfullyLogin).JSON()
}

// @Summary Authentication logout
// @Description Logout using Authorization Header.
// @Tags authentication
// @Produce json
// @Security JWTAuth
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/logout [post]
// Logout will handle logout request.
func (au *Authenticate) Logout(c *gin.Context) {
	metadata, err := au.tk.ExtractTokenMetadata(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	// If the access token exist and it is still valid, then delete both the access token and the refresh token
	errDelete := au.rd.DeleteTokens(metadata)
	if errDelete != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, errDelete)
		return
	}

	response.NewSuccess(c, nil, success.AuthSuccessfullyLogout).JSON()
}

// @Summary Authentication refresh
// @Description Retrieve a new access_token using refresh token.
// @Tags authentication
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param refresh_token formData string true "Refresh token from /login"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 422 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/refresh [post]
// Refresh will handle request to generate new pairs of refresh and access tokens.
func (au *Authenticate) Refresh(c *gin.Context) {
	var mapToken *authorization.JWTRefreshToken
	if err := c.ShouldBind(&mapToken); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	refreshToken := mapToken.RefreshToken

	// Verify the token
	token, errToken := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, exception.ErrorTextAnErrorOccurred
		}
		return []byte(os.Getenv("APP_PRIVATE_KEY")), nil
	})

	// Any error may be due to token expiration
	if errToken != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, errToken)
		return
	}

	// Is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		_ = c.AbortWithError(http.StatusUnauthorized, errToken)
		return
	}

	// Since token is valid, get the uuid:
	claims, okClaims := token.Claims.(jwt.MapClaims)
	if !okClaims {
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextAnErrorOccurred)
	}

	if !token.Valid {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextRefreshTokenIsExpired)
	}

	refreshUUID, ok := claims["refresh_uuid"].(string)
	if !ok {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	UUID := claims["uuid"].(string)

	// Delete the previous Refresh Token
	errDelete := au.rd.DeleteRefresh(refreshUUID)
	if errDelete != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	// Create new pairs of refresh and access tokens
	ts, errCreate := au.tk.CreateToken(UUID, "")
	if errCreate != nil {
		_ = c.AbortWithError(http.StatusForbidden, errCreate)
		return
	}

	// Save the tokens metadata to redis
	errSave := au.rd.CreateAuth(UUID, ts)
	if errSave != nil {
		_ = c.AbortWithError(http.StatusForbidden, errSave)
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	response.NewSuccess(c, tokens, success.AuthSuccessfullyRefreshToken).JSON()
}

// @Summary Authentication forgot password
// @Description Retrieve an email contain link to reset password.
// @Tags authentication
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param email formData string true "Email address"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 422 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/password/forgot [post]
// ForgotPassword will handle request to send email contain link to reset password.
func (au *Authenticate) ForgotPassword(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBind(&user); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := user.ValidateForgotPassword()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	userData, _, errException := au.aa.GetUserByEmail(&user)
	if errException != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
		return
	}

	forgotPasswordToken, _, errException := au.aa.CreateToken(userData)
	if errException != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextAnErrorOccurred)
		return
	}

	nOptions := notify.NotificationOptions{URLPath: forgotPasswordToken.Token}
	nForgotPassword := notification.NewForgotPassword(userData, au.ni, translation.GetLanguage(c), nOptions)
	err := nForgotPassword.Send()
	if len(err) > 0 {
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextAnErrorOccurred)
		return
	}

	response.NewSuccess(c, nil, success.AuthSuccessfullyForgotPassword).JSON()
}

// @Summary Authentication reset password
// @Description Set a new password.
// @Tags authentication
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param token formData string true "Token from forgot password request"
// @Param new_password formData string true "New password"
// @Param confirm_password formData string true "Confirm password"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 422 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/password/reset/{token} [post]
// ForgotPassword will handle request to set a new password.
func (au *Authenticate) ResetPassword(c *gin.Context) {
	var userForgotPassword entity.UserForgotPassword
	var userResetPassword entity.UserResetPassword
	if err := c.ShouldBindUri(&userForgotPassword.Token); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBind(&userResetPassword); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	token := c.Param("token")
	tokenForgotPassword, err := au.aa.GetToken(token)
	if err != nil {
		if errors.Is(err, exception.ErrorTextUserForgotPasswordTokenNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextUserForgotPasswordTokenNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user, err := au.aa.GetUser(tokenForgotPassword.UserUUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextUserNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextUserNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextUserForgotPasswordTokenNotFound)
		return
	}

	validateErr := userResetPassword.ValidateResetPassword()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	hashPassword, errHash := security.Hash(userResetPassword.NewPassword)
	if errHash != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user.Password = string(hashPassword)
	_, errDesc, errException := au.aa.UpdateUser(user.UUID, user)
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

	_ = au.aa.DeactivateToken(tokenForgotPassword.UUID)

	response.NewSuccess(c, nil, success.AuthSuccessfullyResetPassword).JSON()
}
