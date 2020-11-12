package handler

import (
	"fmt"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/interfaces/service"
	"go-rest-skeleton/pkg/response"
	"net/http"

	"github.com/go-session/session"

	"github.com/gin-gonic/gin"
)

type Oauth struct {
	aa service.AuthService
	rd authorization.AuthInterface
	tk authorization.TokenInterface
	ni application.NotifyAppInterface
}

func NewOauth(aa service.AuthService,
	rd authorization.AuthInterface,
	tk authorization.TokenInterface,
	ni application.NotifyAppInterface) *Oauth {
	return &Oauth{
		aa: aa,
		rd: rd,
		tk: tk,
		ni: ni,
	}
}

func (o *Oauth) Login(c *gin.Context) {
	r := c.Request
	w := c.Writer
	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user *entity.User
	if err := c.ShouldBind(&user); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	if user != nil {
		validateErr := user.ValidateLogin()
		if len(validateErr) > 0 {
			exceptionData := response.TranslateErrorForm(c, validateErr)
			c.Set("data", exceptionData)
			_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
			return
		}

		u, _, errException := o.aa.GetUserByEmailAndPassword(user)
		if errException != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, errException)
			return
		}

		sessionStore.Set("LoggedInUserID", u.UUID)
		errStore := sessionStore.Save()
		if errStore != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, errStore)
			return
		}

		c.Redirect(http.StatusFound, "/oauth/auth")
		return
	}
}

func (o *Oauth) Auth(c *gin.Context) {
	r := c.Request
	w := c.Writer
	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(sessionStore)

	if _, ok := sessionStore.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/oauth/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	UUID, _ := sessionStore.Get("LoggedInUserID")
	if UUID == nil {
		c.Header("Content-Type", "text/html; charset: utf-8")
		c.HTML(http.StatusOK, "oauth_login.html", gin.H{})
		return
	}

	c.Header("Content-Type", "text/html; charset: utf-8")
	c.HTML(http.StatusOK, "oauth_authorize.html", gin.H{})
}
