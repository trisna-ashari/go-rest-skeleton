package routers

import (
	"go-rest-skeleton/interfaces/middleware"

	roleV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/role"

	"github.com/gin-gonic/gin"
)

func roleRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	roleV1 := roleV1Point00.NewRoles(r.dbServices.Role, r.redisServices.Auth, rg.authToken)

	v1 := e.Group("/api/v1/external")
	v1.GET("/roles", middleware.Auth(rg.authGateway), roleV1.GetRoles)
	v1.GET("/roles/:uuid", middleware.Auth(rg.authGateway), roleV1.GetRole)
}
