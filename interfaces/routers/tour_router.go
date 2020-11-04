package routers

import (
	"go-rest-skeleton/interfaces/middleware"

	TourV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/Tour"

	"github.com/gin-gonic/gin"
)

func tourRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	TourV1 := TourV1Point00.NewTours(r.dbService.Tour, r.redisService.Auth, rg.authToken)

	guard := middleware.Guard(rg.authGateway)

	v1 := e.Group("/api/v1/external")

	v1.GET("/tours", guard.Authenticate(), guard.Authorize("tour_read"), TourV1.GetTours)
	v1.POST("/tours", guard.Authenticate(), guard.Authorize("tour_create"), TourV1.SaveTour)
	v1.GET("/tours/:uuid", guard.Authenticate(), guard.Authorize("tour_detail"), TourV1.GetTour)
	v1.PUT("/tours/:uuid", guard.Authenticate(), guard.Authorize("tour_update"), TourV1.UpdateTour)
	v1.DELETE("/tours/:uuid", guard.Authenticate(), guard.Authorize("tour_delete"), TourV1.DeleteTour)
}
