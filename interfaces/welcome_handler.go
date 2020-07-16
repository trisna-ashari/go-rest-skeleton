package interfaces

import (
	"fmt"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

// Welcome is a struct to store user repository and token interface.
type Welcome struct {
	ur repository.UserRepository
	tk authorization.TokenInterface
}

// NewWelcomeHandler will initialize welcome handler.
func NewWelcomeHandler(ur repository.UserRepository, tk authorization.TokenInterface) *Welcome {
	return &Welcome{
		ur: ur,
		tk: tk,
	}
}

// Index will handle user request.
func (s *Welcome) Index(c *gin.Context) {
	metadata, errExtract := s.tk.ExtractTokenMetadata(c)
	if errExtract == nil {
		UUID := metadata.UUID
		userData, err := s.ur.GetUser(UUID)
		if err == nil {
			middleware.Formatter(c, nil, fmt.Sprintf("Pak Pong yuk %s :)", userData.FirstName), nil)
			return
		}
	}
	middleware.Formatter(c, nil, "APP v1.0", nil)
}
