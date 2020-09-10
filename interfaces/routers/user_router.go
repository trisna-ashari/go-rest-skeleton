package routers

import (
	"go-rest-skeleton/interfaces/handler"
	"go-rest-skeleton/interfaces/middleware"

	userV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/user"

	"github.com/gin-gonic/gin"
)

func userRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	userV1 := userV1Point00.NewUsers(r.dbService.User, r.storageService.Storage)
	userPreference := handler.NewPreference(r.dbService.UserPreference, r.redisService.Auth, rg.authToken)

	guard := middleware.Guard(rg.authGateway)

	v1 := e.Group("/api/v1/external")

	v1.GET("/users", guard.Authenticate(), guard.Authorize("read"), userV1.GetUsers)
	v1.POST("/users", guard.Authenticate(), guard.Authorize("create"), userV1.SaveUser)
	v1.GET("/users/:uuid", guard.Authenticate(), guard.Authorize("detail"), userV1.GetUser)
	v1.PUT("/users/:uuid", guard.Authenticate(), guard.Authorize("update"), userV1.UpdateUser)
	v1.PUT("/users/:uuid/avatar", guard.Authenticate(), guard.Authorize("update"), userV1.UpdateAvatar)
	v1.DELETE("/users/:uuid", guard.Authenticate(), guard.Authorize("delete"), userV1.DeleteUser)

	v1.GET("/preference", guard.Authenticate(), userPreference.GerPreference)
	v1.PUT("/preference", guard.Authenticate(), userPreference.UpdatePreference)
	v1.POST("/preference/reset", guard.Authenticate(), userPreference.ResetPreference)

}
