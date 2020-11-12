package routers

import (
	"go-rest-skeleton/config"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces/middleware"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

// Router is a struct contains needed dependencies to init MainRouter().
type Router struct {
	conf                *config.Config
	dbService           *persistence.Repositories
	redisService        *persistence.RedisService
	storageService      *persistence.StorageService
	notificationService *persistence.NotificationService
}

// RouterAuthGateway is a struct contains needed dependencies to init Routes.
type RouterAuthGateway struct {
	authGateway *authorization.Gateway
	authToken   *authorization.Token
	authOauth   *authorization.OauthAuth
}

// NewRouter is a constructor uses to construct Router.
func NewRouter(
	conf *config.Config,
	dbService *persistence.Repositories,
	redisService *persistence.RedisService,
	storageService *persistence.StorageService,
	notificationService *persistence.NotificationService) *Router {
	return &Router{
		conf:                conf,
		dbService:           dbService,
		redisService:        redisService,
		storageService:      storageService,
		notificationService: notificationService,
	}
}

// NewRouter is a constructor uses to construct RouterAuthGateway.
func NewRouterAuthGateway(ag *authorization.Gateway, at *authorization.Token, oa *authorization.OauthAuth) *RouterAuthGateway {
	return &RouterAuthGateway{
		authGateway: ag,
		authToken:   at,
		authOauth:   oa,
	}
}

// MainRouter is a method to initialize gin engine.
func (r *Router) Init() *gin.Engine {
	// Logging
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// Init response options
	optResponse := middleware.ResponseOptions{
		Environment:     r.conf.AppEnvironment,
		DebugMode:       r.conf.DebugMode,
		DefaultLanguage: r.conf.AppLanguage,
		DefaultTimezone: r.conf.AppTimezone,
	}

	// Init authorization
	authBasic := authorization.NewBasicAuth(r.dbService.User)
	authJWT := authorization.NewJWTAuth(r.conf.KeyConfig, r.redisService.Client)
	authToken := authorization.NewToken(r.conf.KeyConfig, r.redisService.Client)
	authOauth := authorization.NewOauthAuth(r.conf.KeyConfig, r.redisService.Client)
	authGateway := authorization.NewAuthGateway(authBasic, authJWT, authOauth, r.dbService.User, r.dbService.Role)

	// Init gin
	if !r.conf.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.Default()
	e.Use(middleware.NewResponse(optResponse).Handler())
	e.Use(middleware.SetRequestID(middleware.RequestIDOptions{AllowSetting: r.conf.EnableRequestID}))
	e.Use(middleware.CORS(middleware.CORSOptions{AllowSetting: r.conf.EnableCors}))
	e.Use(middleware.SetLogger(middleware.LoggerOptions{AllowSetting: r.conf.EnableLogger}))
	e.Use(middleware.APIVersion())
	e.Use(middleware.Recovery())

	// Init Routes
	rg := NewRouterAuthGateway(authGateway, authToken, authOauth)
	authRoutes(e, r, rg)
	devRoutes(e, r)
	noRoutes(e)
	oauthServerRoutes(e, r, rg)
	oauthClientRoutes(e, r, rg)
	roleRoutes(e, r, rg)
	tourRoutes(e, r, rg)
	userRoutes(e, r, rg)
	welcomeRoutes(e)
	graphqlRoute(e, r, rg)

	return e
}
