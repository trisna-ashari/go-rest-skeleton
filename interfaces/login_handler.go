package interfaces

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces/middleware"
	"golang.org/x/exp/errors"
	"net/http"
	"os"
)

type Authenticate struct {
	us application.UserAppInterface
	rd authorization.AuthInterface
	tk authorization.TokenInterface
}

func NewAuthenticate(uApp application.UserAppInterface, rd authorization.AuthInterface, tk authorization.TokenInterface) *Authenticate {
	return &Authenticate{
		us: uApp,
		rd: rd,
		tk: tk,
	}
}

func (au *Authenticate) Profile(c *gin.Context) {
	metadata, err := au.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, errors.New("api.msg.error.unauthorized"))
		return
	}

	UUID := metadata.UUID
	userData, _ := au.us.GetUser(UUID)
	middleware.Formatter(c, userData.DetailUser(), "api.msg.success.successfully_get_profile", nil)
}

func (au *Authenticate) SwitchLanguage(c *gin.Context) {
	_, err := au.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, errors.New("api.msg.error.unauthorized"))
		return
	}
	var language map[string]interface{}
	if err := c.ShouldBindJSON(&language); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, errors.New("api.msg.error.unprocessable_entity"))
		return
	}
	selectedLanguage := language["language"].(string)
	if persistence.IsValidAcceptLanguage(selectedLanguage) != true {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, errors.New("api.msg.error.unprocessable_entity"))
		return
	}

	c.Header("AcceptLanguage", selectedLanguage)

	userData := make(map[string]interface{})
	userData["language"] = selectedLanguage

	middleware.Formatter(c, userData, "api.msg.success.successfully_switch_language", nil)
}

func (au *Authenticate) Login(c *gin.Context) {
	var user *entity.User
	var tokenErr = map[string]string{}

	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, errors.New("api.msg.error.unprocessable_entity"))
		return
	}
	validateUser := user.Validate("login")
	if len(validateUser) > 0 {
		c.Set("data", validateUser)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, errors.New("api.msg.error.unprocessable_entity"))
		return
	}
	u, userErr := au.us.GetUserByEmailAndPassword(user)
	if userErr != nil {
		c.Set("data", userErr)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, errors.New("api.msg.error.unprocessable_entity"))
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

func (au *Authenticate) Logout(c *gin.Context) {
	metadata, err := au.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, errors.New("api.msg.error.unauthorized"))
		return
	}

	// If the access token exist and it is still valid, then delete both the access token and the refresh token
	deleteErr := au.rd.DeleteTokens(metadata)
	if deleteErr != nil {
		c.JSON(http.StatusUnauthorized, deleteErr.Error())
		return
	}
	middleware.Formatter(c, nil, "api.msg.success.successfully_logout", nil)
}

// Refresh is the function that uses the refresh_token to generate new pairs of refresh and access tokens.
func (au *Authenticate) Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	// Verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	// Any error may be due to token expiration
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	// Is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	// Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			_ = c.AbortWithError(http.StatusUnprocessableEntity, errors.New("api.msg.error.unprocessable_entity"))
			return
		}
		UUID := fmt.Sprintf("%.f", claims["UUID"])

		// Delete the previous Refresh Token
		delErr := au.rd.DeleteRefresh(refreshUuid)
		if delErr != nil {
			c.JSON(http.StatusUnauthorized, "api.msg.error.unauthorized")
			return
		}

		// Create new pairs of refresh and access tokens
		ts, createErr := au.tk.CreateToken(UUID)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}

		// Save the tokens metadata to redis
		saveErr := au.rd.CreateAuth(UUID, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		middleware.Formatter(c, tokens, "api.msg.success.successfully_refresh_token", nil)
	} else {
		_ = c.AbortWithError(http.StatusUnauthorized, errors.New("api.msg.error.refresh_token_expired"))
	}
}