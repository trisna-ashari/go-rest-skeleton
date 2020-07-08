package application

import (
	"github.com/gin-gonic/gin"
)

type welcomeApp struct {

}

type WelcomeAppInterface interface {
	Index(c *gin.Context) (interface{}, error)
	Greeting(c *gin.Context) (interface{}, error)
}

var _ WelcomeAppInterface = &welcomeApp{}

func (w *welcomeApp) Index(c *gin.Context) (interface{}, error) {
	return "Hai", nil
}

func (w *welcomeApp) Greeting(c *gin.Context) (interface{}, error){
	return w.Index(c)
}
