package middleware

import (
	"go-rest-skeleton/pkg/response"

	"github.com/gin-gonic/gin"
)

// DefaultGenericError is to define default generic error for production environment.
const DefaultGenericError = `an_error_occurred`

// ResponseOptions is a struct to store options for error response.
type ResponseOptions struct {
	Environment     string
	DebugMode       bool
	DefaultLanguage string
	DefaultTimezone string
	GenericError    string
	LogFunc         func(err string, code int, messages map[string]interface{})
}

// NewResponse will initialize response middleware.
func NewResponse(o ResponseOptions) *ResponseOptions {
	return &ResponseOptions{
		DebugMode:       o.DebugMode,
		DefaultLanguage: o.DefaultLanguage,
		DefaultTimezone: o.DefaultTimezone,
		GenericError:    DefaultGenericError,
	}
}

// Handler will handle any error response.
func (r *ResponseOptions) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset: utf-8")
		c.Next()

		if c.Errors.Last() == nil {
			return
		}

		err := c.Errors.Last().Err
		c.Errors = c.Errors[:0]
		message := err.Error()

		if r.Environment == "production" && c.Writer.Status() == 500 {
			message = r.GenericError
			response.NewError(c, message).JSON()
			return
		}

		response.NewError(c, message).JSON()
	}
}
