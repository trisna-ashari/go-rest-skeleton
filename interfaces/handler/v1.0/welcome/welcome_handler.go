package welcomev1point00

import (
	"go-rest-skeleton/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

// WelcomeHandler is a struct.
type WelcomeHandler struct {
}

// NewWelcomeHandler will initialize welcome handler.
func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}

// Index will handle request.
func (s *WelcomeHandler) Index(c *gin.Context) {
	middleware.Formatter(c, nil, "PONG v1.0", nil)
}
