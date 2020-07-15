package middleware

import (
	"go-rest-skeleton/infrastructure/util"

	"github.com/ansel1/merry"
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

type errOutput struct {
	ErrorHTTPCode    int                    `json:"code"`
	Data             interface{}            `json:"data"`
	Args             map[string]interface{} `json:"details,omitempty"`
	Message          string                 `json:"message"`
	ErrorTracingCode string                 `json:"error_code,omitempty"`
}

type successOutput struct {
	SuccessHTTPCode int         `json:"code"`
	Data            interface{} `json:"data"`
	Message         string      `json:"message"`
	Meta            interface{} `json:"meta,omitempty"`
}

// New will initialize response middleware.
func New(o ResponseOptions) *ResponseOptions {
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

		// No errors then skip
		if c.Errors.Last() == nil {
			return
		}

		// Get last error, clear all errors
		err := c.Errors.Last().Err
		c.Errors = c.Errors[:0]

		// Form the response
		response := errOutput{Message: err.Error(), Args: map[string]interface{}{}}
		for key, val := range merry.Values(err) {
			if key == "message" || key == "http status code" {
				continue
			}
			if argKey, ok := key.(string); ok {
				response.Args[argKey] = val
			}
		}

		// Set error code, data, error tracing code
		errHTTPCode := c.Writer.Status()
		errData, _ := c.Get("data")
		errTracingCode, _ := c.Get("errorTracingCode")
		response.ErrorHTTPCode = errHTTPCode
		response.Data = errData
		response.ErrorTracingCode = errTracingCode.(string)

		// Set translations
		translatedMessage, language := util.NewTranslation(c, "error", response.Message, map[string]interface{}{})
		c.Header("Accept-Language", language)

		// If environment is production
		if r.Environment == "production" && errHTTPCode == 500 {
			httpCode := c.Writer.Status()
			response.Message = r.GenericError
			response.Args = nil
			c.JSON(httpCode, response)
			return
		}

		// Set message
		response.Message = translatedMessage

		// Add the error's stack if Debug is enabled
		if r.DebugMode {
			response.Args[`stack`] = merry.Stacktrace(err)
		}

		// Log the error
		if r.LogFunc != nil {
			r.LogFunc(err.Error(), errHTTPCode, response.Args)
		}

		// Return error response
		c.JSON(errHTTPCode, response)
	}
}
