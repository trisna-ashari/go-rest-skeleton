package welcome_v2_0_0

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
	middleware.Formatter(c, nil, "PONG v2.0", nil)
}

func (s *WelcomeHandler) Greeting(c *gin.Context) {
	middleware.Formatter(c, nil, "GREETING v2.0", nil)
}
