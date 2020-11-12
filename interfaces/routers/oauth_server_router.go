package routers

import (
	"fmt"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/interfaces/handler"
	"go-rest-skeleton/interfaces/service"
	"go-rest-skeleton/pkg/util"
	"log"
	"net/http"
	"net/url"

	"github.com/go-session/session"

	"github.com/go-oauth2/oauth2/v4"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"

	"github.com/go-oauth2/oauth2/v4/manage"

	"github.com/gin-gonic/gin"
)

func oauthServerRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	config := r.conf
	serverConfig := &server.Config{
		TokenType:            config.Oauth2Config.OauthToken,
		AllowedResponseTypes: []oauth2.ResponseType{oauth2.Code, oauth2.Token},
		AllowedGrantTypes: []oauth2.GrantType{
			oauth2.AuthorizationCode,
			oauth2.PasswordCredentials,
			oauth2.ClientCredentials,
			oauth2.Refreshing,
		},
	}

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapAccessGenerate(rg.authOauth)

	clientStore := store.NewClientStore()
	errStore := clientStore.Set(config.Oauth2Config.OauthID, &models.Client{
		ID:     config.Oauth2Config.OauthID,
		Secret: config.Oauth2Config.OauthSecret,
		Domain: config.Oauth2Config.OauthDomain,
	})

	errStore = clientStore.Set("1", &models.Client{
		ID:     "1",
		Secret: "1",
		Domain: "http://localhost:8181/other",
	})

	if errStore != nil {
		log.Fatal(errStore)
	}

	manager.MapClientStorage(clientStore)
	oauthServer := server.NewServer(serverConfig, manager)
	oauthServer.SetPasswordAuthorizationHandler(func(email, password string) (userID string, err error) {
		user, _, _ := rg.authGateway.US.GetUserByEmailAndPassword(&entity.User{Email: email, Password: password})
		userID = user.UUID

		return
	})

	oauthServer.SetUserAuthorizationHandler(userAuthorizeHandler)

	oauth := handler.NewOauth(
		service.NewAuthService(r.dbService.User, r.dbService.UserForgotPassword),
		r.redisService.Auth,
		rg.authToken,
		r.notificationService.Notification)

	sessionStore := cookie.NewStore([]byte(config.KeyConfig.AppPrivateKey))
	e.Use(sessions.Sessions("go-rest-skeleton-sessions", sessionStore))
	e.LoadHTMLFiles(
		fmt.Sprintf("%s/interfaces/handler/oauth_login.html", util.RootDir()),
		fmt.Sprintf("%s/interfaces/handler/oauth_authorize.html", util.RootDir()),
	)

	o := e.Group("/oauth")
	o.GET("/", func(c *gin.Context) {
		defaultHandler(c)
	})
	o.GET("/login", func(c *gin.Context) {
		req, err := oauthServer.ValidationAuthorizeRequest(c.Request)
		if req == nil || err != nil {
			defaultHandler(c)
			return
		}

		loginHandler(c)
	})
	o.POST("/login", oauth.Login)
	o.GET("/auth", oauth.Auth)
	o.GET("/authorize", func(c *gin.Context) {
		authorizeHandler(c, oauthServer)
	})
	o.POST("/authorize", func(c *gin.Context) {
		authorizeHandler(c, oauthServer)
	})
	o.GET("/token", func(c *gin.Context) {
		tokenHandler(c, oauthServer)
	})
	o.POST("/token", func(c *gin.Context) {
		tokenHandler(c, oauthServer)
	})

}

func defaultHandler(c *gin.Context) {
	u := oauthConfig.AuthCodeURL("ok")
	http.Redirect(c.Writer, c.Request, u, http.StatusFound)
}

func loginHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset: utf-8")
	c.HTML(http.StatusOK, "oauth_login.html", gin.H{})
}

func authorizeHandler(c *gin.Context, srv *server.Server) {
	r := c.Request
	w := c.Writer

	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, ok := sessionStore.Get("LoggedInUserID")
	if ok {
		var form url.Values
		if v, ok := sessionStore.Get("ReturnUri"); ok {
			form = v.(url.Values)
		}
		r.Form = form
	}

	sessionStore.Delete("ReturnUri")
	err = sessionStore.Save()
	if err != nil {
		return
	}

	err = srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	sessionStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := sessionStore.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			err = r.ParseForm()
			if err != nil {
				return
			}
		}

		sessionStore.Set("ReturnUri", r.Form)
		err = sessionStore.Save()
		if err != nil {
			return
		}

		w.Header().Set("Location", fmt.Sprintf("/oauth/login%s", util.BuildEncodedQueryString(r.Form)))
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	sessionStore.Delete("LoggedInUserID")
	err = sessionStore.Save()
	if err != nil {
		return
	}

	return
}

func tokenHandler(c *gin.Context, srv *server.Server) {
	r := c.Request
	w := c.Writer
	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
