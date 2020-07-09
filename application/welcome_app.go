package application

import (
	"github.com/gin-gonic/gin"
)

type welcomeApp struct {
}

// WelcomeAppInterface is an interface.
type WelcomeAppInterface interface {
	Index(c *gin.Context) (interface{}, error)
}

var _ WelcomeAppInterface = &welcomeApp{}

// Index is implementation of method Index.
func (w *welcomeApp) Index(c *gin.Context) (interface{}, error) {
	return "Hai", nil
}
