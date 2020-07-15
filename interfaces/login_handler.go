package interfaces

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

// SwitchLanguage will switch active language for current user.
func (au *Authenticate) SwitchLanguage(c *gin.Context) {
	_, exists := c.Get("UUID")
	if !exists {
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return
	}

	var language map[string]interface{}
	if err := c.ShouldBindJSON(&language); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	selectedLanguage := language["language"].(string)

	if !util.IsValidAcceptLanguage(selectedLanguage) {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	c.Header("AcceptLanguage", selectedLanguage)

	userData := make(map[string]interface{})
	userData["language"] = selectedLanguage

	middleware.Formatter(c, userData, "api.msg.success.successfully_switch_language", nil)
}

// Login will handle login request.
func (au *Authenticate) Login(c *gin.Context) {
	var user *entity.User
	var tokenErr = map[string]string{}

	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	validateUser := user.ValidateLogin(c)
	if len(validateUser) > 0 {
		c.Set("data", validateUser)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	u, userErr := au.us.GetUserByEmailAndPassword(user)
	if userErr != nil {
		c.Set("data", userErr)
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
		c.JSON(http.StatusInternalServerError, saveErr.Error())
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

// Refresh will handle request to generate new pairs of refresh and access tokens.
func (au *Authenticate) Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	refreshToken := mapToken["refresh_token"]

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
