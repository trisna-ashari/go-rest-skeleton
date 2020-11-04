package routers

import (
	"go-rest-skeleton/interfaces/middleware"

	roleV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/role"

	"github.com/gin-gonic/gin"
)

func roleRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	roleV1 := roleV1Point00.NewRoles(r.dbService.Role, r.redisService.Auth, rg.authToken)

	guard := middleware.Guard(rg.authGateway)

	v1 := e.Group("/api/v1/external")

	v1.GET("/roles", guard.Authenticate(), guard.Authorize("role_read"), roleV1.GetRoles)
	v1.POST("/roles", guard.Authenticate(), guard.Authorize("role_create"), roleV1.SaveRole)
	v1.GET("/roles/:uuid", guard.Authenticate(), guard.Authorize("role_detail"), roleV1.GetRole)
	v1.PUT("/roles/:uuid", guard.Authenticate(), guard.Authorize("role_update"), roleV1.UpdateRole)
	v1.DELETE("/roles/:uuid", guard.Authenticate(), guard.Authorize("role_delete"), roleV1.DeleteRole)
}
