package middleware

import (
	"github.com/ansel1/merry"
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/infrastructure/persistence"
	"os"
)

const DefaultGenericError = `an_error_occurred`

type ResponseOptions struct {
	DebugMode       bool
	DefaultLanguage string
}

type errOutput struct {
	Code    int                    `json:"code"`
	Data    interface{}            `json:"data"`
	Args    map[string]interface{} `json:"details,omitempty"`
	Message string                 `json:"message"`
}

type successOutput struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type successOutputWithMeta struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
}

type Middleware struct {
	Debug bool
	GenericError string
	LogFunc func(err string, code int, messages map[string]interface{})
}

func New(options ResponseOptions) *Middleware {
	return &Middleware{Debug: options.DebugMode, GenericError: DefaultGenericError}
}

func (m *Middleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		c.Header("Content-Type", "application/json; charset: utf-8")
		c.Next()

		// No errors then skip
		if c.Errors.Last() == nil {
			return
		}

		// Get env
		env := os.Getenv("APP_ENV")

		// Get last error, clear all errors
		err := c.Errors.Last().Err
		c.Errors = c.Errors[:0]

		// Form the response
		response := errOutput{Message: err.Error(), Args: map[string]interface{}{}}
		for key, val := range merry.Values(err) {
			if key == "message" || key == "http status code" {
				continue
			}
			if key, ok := key.(string); ok {
				response.Args[key] = val
			}
		}

		// Set error code
		errCode := c.Writer.Status()
		response.Code = errCode

		// Set error data
		errData, _ := c.Get("data")
		response.Data = errData

		// Set translations
		translatedMessage, language := persistence.NewTranslation(c, "error", response.Message)
		c.Header("Accept-Language", language)

		// If environment is production
		if env == "production" && errCode == 500 {
			errCode := c.Writer.Status()
			response.Message = m.GenericError
			response.Args = nil
			c.JSON(errCode, response)
			return
		}

		// Set message
		response.Message = translatedMessage

		// Add the error's stack if Debug is enabled
		if m.Debug == true {
			response.Args[`stack`] = merry.Stacktrace(err)
		}

		// Log the error
		if m.LogFunc != nil {
			m.LogFunc(err.Error(), errCode, response.Args)
		}

		// Return error response
		c.JSON(errCode, response)
		return
	}
}

func Formatter(c *gin.Context, data interface{}, message string) {
	response := successOutput{Code: c.Writer.Status(), Message: "OK"}
	response.Data = data
	if message != "" {
		response.Message = message
	}
	translatedMessage, language := persistence.NewTranslation(c, "error", response.Message)
	response.Message = translatedMessage

	c.Header("Accept-Language", language)
	c.JSON(c.Writer.Status(), response)
}

func FormatterWithMeta(c *gin.Context, data interface{}, message string, meta interface{}) {
	response := successOutputWithMeta{Code: c.Writer.Status(), Message: "OK"}
	response.Data = data
	response.Meta = meta
	if message != "" {
		response.Message = message
	}
	translatedMessage, language := persistence.NewTranslation(c, "error", response.Message)
	response.Message = translatedMessage

	c.Header("Accept-Language", language)
	c.JSON(c.Writer.Status(), response)
}
