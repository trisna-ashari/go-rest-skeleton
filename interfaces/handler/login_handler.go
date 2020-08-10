package handler

import (
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/exception"
	"go-rest-skeleton/infrastructure/util"
	"go-rest-skeleton/interfaces/middleware"
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
}

// NewAuthenticate will initialize authenticate interface for login handler.
func NewAuthenticate(
	uApp application.UserAppInterface,
	rd authorization.AuthInterface,
	tk authorization.TokenInterface) *Authenticate {
	return &Authenticate{
		us: uApp,
		rd: rd,
		tk: tk,
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
	middleware.Formatter(c, userData.DetailUser(), "api.msg.success.successfully_get_profile", nil)
}

// @Summary Switch language preference
// @Description Change language preference.
// @Tags authentication
// @Accept mpfd
// @Produce json
// @Param language formData string true "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Success 200 {object} middleware.successOutput
// @Failure 400 {object} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
// @Router /api/v1/external/language [post]
// SwitchLanguage will switch active language for current user.
func (au *Authenticate) SwitchLanguage(c *gin.Context) {
	_, exists := c.Get("UUID")
	if !exists {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	var language *util.TranslationLanguage
	if err := c.ShouldBind(&language); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	selectedLanguage := language.Language

	if !util.IsValidAcceptLanguage(selectedLanguage) {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	c.Header("AcceptLanguage", selectedLanguage)

	userData := make(map[string]interface{})
	userData["language"] = selectedLanguage

	middleware.Formatter(c, userData, "api.msg.success.successfully_switch_language", nil)
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
// @Failure 404 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
// @Router /api/v1/external/login [post]
// Login will handle login request.
func (au *Authenticate) Login(c *gin.Context) {
	var user *entity.User
	var tokenErr = map[string]string{}

	if err := c.ShouldBind(&user); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	validateUser := user.ValidateLogin(c)
	if len(validateUser) > 0 {
		c.Set("data", validateUser)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	u, errDesc, _ := au.us.GetUserByEmailAndPassword(user)
	if errDesc != nil {
		c.Set("data", errDesc)
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}
	ts, tErr := au.tk.CreateToken(u.UUID)
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		_ = c.AbortWithError(http.StatusUnprocessableEntity, tErr)
		return
	}
	saveErr := au.rd.CreateAuth(u.UUID, ts)
	if saveErr != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, saveErr)
		return
	}
	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["uuid"] = u.UUID
	userData["first_name"] = u.FirstName
	userData["last_name"] = u.LastName
	userData["language"] = os.Getenv("APP_LANG")

	middleware.Formatter(c, userData, "api.msg.success.successfully_login", nil)
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
	deleteErr := au.rd.DeleteTokens(metadata)
	if deleteErr != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, deleteErr)
		return
	}
	middleware.Formatter(c, nil, "api.msg.success.successfully_logout", nil)
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
// @Failure 500 {object} middleware.errOutput
// @Router /api/v1/external/refresh [post]
// Logout will handle logout request.
// Refresh will handle request to generate new pairs of refresh and access tokens.
func (au *Authenticate) Refresh(c *gin.Context) {
	var mapToken *authorization.JWTRefreshToken
	if err := c.ShouldBind(&mapToken); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	refreshToken := mapToken.RefreshToken

	// Verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, exception.ErrorTextAnErrorOccurred
		}
		return []byte(os.Getenv("APP_PRIVATE_KEY")), nil
	})

	// Any error may be due to token expiration
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Since token is valid, get the uuid:
	claims, okClaims := token.Claims.(jwt.MapClaims)
	if !okClaims {
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextAnErrorOccurred)
	}

	if token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
			return
		}
		UUID := claims["uuid"].(string)

		// Delete the previous Refresh Token
		delErr := au.rd.DeleteRefresh(refreshUUID)
		if delErr != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
			return
		}

		// Create new pairs of refresh and access tokens
		ts, createErr := au.tk.CreateToken(UUID)
		if createErr != nil {
			_ = c.AbortWithError(http.StatusForbidden, createErr)
			return
		}

		// Save the tokens metadata to redis
		saveErr := au.rd.CreateAuth(UUID, ts)
		if saveErr != nil {
			_ = c.AbortWithError(http.StatusForbidden, saveErr)
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		middleware.Formatter(c, tokens, "api.msg.success.successfully_refresh_token", nil)
	} else {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextRefreshTokenIsExpired)
	}
}
