package routers

import (
	"go-rest-skeleton/interfaces/middleware"

	userV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/user"
	userV2Point00 "go-rest-skeleton/interfaces/handler/v2.0/user"

	"github.com/gin-gonic/gin"
)

func userRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	userV1 := userV1Point00.NewUsers(r.dbService.User, r.storageService.Storage)
	userV2 := userV2Point00.NewUsers(r.dbService.User)

	guard := middleware.Guard(rg.authGateway)

	v1 := e.Group("/api/v1/external")
	v2 := e.Group("/api/v2/external")

	v1.GET("/users", guard.Authenticate(), guard.Authorize("user_read"), userV1.GetUsers)
	v1.POST("/users", guard.Authenticate(), guard.Authorize("user_create"), userV1.SaveUser)
	v1.GET("/users/:uuid", guard.Authenticate(), guard.Authorize("user_detail"), userV1.GetUser)
	v1.PUT("/users/:uuid", guard.Authenticate(), guard.Authorize("user_update"), userV1.UpdateUser)
	v1.PUT("/users/:uuid/avatar", guard.Authenticate(), guard.Authorize("user_update"), userV1.UpdateAvatar)
	v1.DELETE("/users/:uuid", guard.Authenticate(), guard.Authorize("user_delete"), userV1.DeleteUser)

	v2.GET("/users", guard.Authenticate(), userV2.GetUsers)
	v2.POST("/users", guard.Authenticate(), userV2.SaveUser)
	v2.GET("/users/:uuid", guard.Authenticate(), userV2.GetUser)
}
