package routers

import (
	"go-rest-skeleton/interfaces/handler"
	"go-rest-skeleton/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func authRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	authenticate := handler.NewAuthenticate(r.dbService.User, r.redisService.Auth, rg.authToken)

	v1 := e.Group("/api/v1/external")

	v1.GET("/profile", middleware.Auth(rg.authGateway), authenticate.Profile)
	v1.POST("/login", authenticate.Login)
	v1.POST("/logout", middleware.Auth(rg.authGateway), authenticate.Logout)
	v1.POST("/refresh", authenticate.Refresh)
	v1.POST("/language", middleware.Auth(rg.authGateway), authenticate.SwitchLanguage)
}
