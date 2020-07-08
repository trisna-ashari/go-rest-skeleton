package welcome_v1_0_0

import (
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/application"
	"go-rest-skeleton/interfaces/middleware"
)

type WelcomeHandler struct {
	wa application.WelcomeAppInterface
}

func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}

func (s *WelcomeHandler) Index(c *gin.Context) {
	middleware.Formatter(c, nil, "PONG v1.0")
}

func (s *WelcomeHandler) Greeting(c *gin.Context) {
	middleware.Formatter(c, nil, "GREETING v1.0")
}
