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

	v1.GET("/roles", guard.Authenticate(), guard.Authorize("read"), roleV1.GetRoles)
	v1.POST("/roles", guard.Authenticate(), guard.Authorize("create"), roleV1.SaveRole)
	v1.GET("/roles/:uuid", guard.Authenticate(), guard.Authorize("detail"), roleV1.GetRole)
	v1.PUT("/roles/:uuid", guard.Authenticate(), guard.Authorize("update"), roleV1.UpdateRole)
	v1.DELETE("/roles/:uuid", guard.Authenticate(), guard.Authorize("delete"), roleV1.DeleteRole)
}
