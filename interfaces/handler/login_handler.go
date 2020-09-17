package handler

import (
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/infrastructure/message/success"
	"go-rest-skeleton/infrastructure/notify"
	"go-rest-skeleton/infrastructure/notify/notification"
	"go-rest-skeleton/interfaces/middleware"
	"go-rest-skeleton/pkg/translation"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authenticate is a struct to store interfaces.
type Authenticate struct {
	us application.UserAppInterface
	rd authorization.AuthInterface
	tk authorization.TokenInterface
	ni application.NotifyAppInterface
}

// NewAuthenticate will initialize authenticate interface for login handler.
func NewAuthenticate(
	uApp application.UserAppInterface,
	rd authorization.AuthInterface,
	tk authorization.TokenInterface,
	ni application.NotifyAppInterface) *Authenticate {
	return &Authenticate{
		us: uApp,
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
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 401 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
// @Router /api/v1/external/profile [get]
// Profile will return detail user of current logged in user.
func (au *Authenticate) Profile(c *gin.Context) {
	UUID, exists := c.Get("UUID")
	if !exists {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	userData, _ := au.us.GetUserWithRoles(UUID.(string))

	middleware.Formatter(c, userData.DetailUser(), success.AuthSuccessfullyGetProfile, nil)
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
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 401 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 422 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
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
		exceptionData := exception.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	u, _, errException := au.us.GetUserByEmailAndPassword(user)
	if errException != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, errException)
		return
	}
	ts, tErr := au.tk.CreateToken(u.UUID)
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

	middleware.Formatter(c, userData, success.AuthSuccessfullyLogin, nil)
}

// @Summary Authentication logout
// @Description Logout using Authorization Header.
// @Tags authentication
// @Produce json
// @Security JWTAuth
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 401 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
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

	middleware.Formatter(c, nil, success.AuthSuccessfullyLogout, nil)
}

// @Summary Authentication refresh
// @Description Retrieve a new access_token using refresh token.
// @Tags authentication
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param refresh_token formData string true "Refresh token from /login"
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 422 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
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
	ts, errCreate := au.tk.CreateToken(UUID)
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

	middleware.Formatter(c, tokens, success.AuthSuccessfullyRefreshToken, nil)
}

// @Summary Authentication forgot password
// @Description Retrieve an email contain link to reset password.
// @Tags authentication
// @Accept mpfd
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param email formData string true "Email address"
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 422 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
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
		exceptionData := exception.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	userData, _, errException := au.us.GetUserByEmail(&user)
	if errException != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
		return
	}

	nOptions := notify.NotificationOptions{URLPath: "token"}
	nForgotPassword := notification.NewForgotPassword(userData, au.ni, translation.GetLanguage(c), nOptions)
	err := nForgotPassword.Send()
	if len(err) > 0 {
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextAnErrorOccurred)
		return
	}

	middleware.Formatter(c, nil, success.AuthSuccessfullyForgotPassword, nil)
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
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 422 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
// @Router /api/v1/external/password/reset/{token} [post]
// ForgotPassword will handle request to set a new password.
func (au *Authenticate) ResetPassword(c *gin.Context) {
	middleware.Formatter(c, nil, success.AuthSuccessfullyResetPassword, nil)
}
