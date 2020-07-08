package v1_0

import (
	"go-rest-skeleton/interfaces/middleware"
	"github.com/gin-gonic/gin"
)

type PingHandler struct {

}

func (s *PingHandler) Index(c *gin.Context) {
	middleware.Formatter(c, nil, "PONG 1.0")
}