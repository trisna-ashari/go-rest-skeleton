package interfaces

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/interfaces/middleware"
)

type Welcome struct {
	ur repository.UserRepository
	tk authorization.TokenInterface
}

func NewWelcomeHandler(ur repository.UserRepository, tk authorization.TokenInterface) *Welcome {
	return &Welcome{
		ur: ur,
		tk: tk,
	}
}

func (s *Welcome) Index(c *gin.Context) {
	metadata, errExtract := s.tk.ExtractTokenMetadata(c.Request)
	if errExtract == nil {
		UUID := metadata.UUID
		userData, err := s.ur.GetUser(UUID)
		if err == nil {
			middleware.Formatter(c, nil, fmt.Sprintf("Pak Pong yuk %s :)", userData.FirstName),nil)
			return
		}
	}
	middleware.Formatter(c, nil, "APP v1.0", nil)
}

func (s *Welcome) Greeting(c *gin.Context) {
	middleware.Formatter(c, nil, "GREETING v1.0", nil)
}

