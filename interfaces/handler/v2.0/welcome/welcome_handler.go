package welcome_v2_0_0

import (
	"go-rest-skeleton/application"
	"go-rest-skeleton/interfaces/middleware"
	"github.com/gin-gonic/gin"
)

type WelcomeHandler struct {
	wa application.WelcomeAppInterface
}

func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}

func (s *WelcomeHandler) Index(c *gin.Context) {
	middleware.Formatter(c, nil, "PONG v2.0")
}

func (s *WelcomeHandler) Greeting(c *gin.Context) {
	middleware.Formatter(c, nil, "GREETING v2.0")
}
