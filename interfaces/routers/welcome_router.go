package routers

import (
	"github.com/gin-gonic/gin"
	welcomeV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/welcome"
	welcomeV2Point00 "go-rest-skeleton/interfaces/handler/v2.0/welcome"
)

func welcomeRoutes(e *gin.Engine) {
	welcomeV1 := welcomeV1Point00.NewWelcomeHandler()
	welcomeV2 := welcomeV2Point00.NewWelcomeHandler()

	v1 := e.Group("/api/v1/external")
	v2 := e.Group("/api/v2/external")

	v1.GET("/welcome", welcomeV1.Index)

	// Welcome
	v2.GET("/welcome", welcomeV2.Index)
}
