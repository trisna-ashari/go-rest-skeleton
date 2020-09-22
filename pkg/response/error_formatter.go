package response

import (
	"go-rest-skeleton/pkg/translation"
	"strings"

	"github.com/gin-gonic/gin"
)

type errorOutput struct {
	c                *gin.Context
	language         string
	ErrorHTTPCode    int         `json:"code"`
	Data             interface{} `json:"data"`
	Message          string      `json:"message"`
	Args             string      `json:"args,omitempty"`
	ErrorTracingCode string      `json:"error_code,omitempty"`
}

// NewError will is a constructor will initialize errorOutput.
func NewError(c *gin.Context, message string) *errorOutput {
	var errorTracingCode interface{}
	var errorMessageData = make(map[string]interface{})

	errorHTTPCode := c.Writer.Status()
	errorData, _ := c.Get("data")
	errorArgs, _ := c.Get("args")
	errorTracingCode, _ = c.Get("errorTracingCode")

	if errorArgs != nil {
		for _, arg := range strings.Split(errorArgs.(string), ";") {
			splitArg := strings.Split(arg, ":")
			argKey := splitArg[0]
			argVal := splitArg[1]
			errorMessageData[argKey] = argVal
		}
	}

	errMessage, language := translation.NewTranslation(c, "error", message, errorMessageData)

	errOutput := &errorOutput{
		c:             c,
		language:      language,
		ErrorHTTPCode: errorHTTPCode,
		Data:          errorData,
		Message:       errMessage,
	}

	if errorTracingCode != nil {
		errOutput.ErrorTracingCode = errorTracingCode.(string)
	}

	return errOutput
}

// WithArgs is a function uses to attach args to the errorOutput.
func (eo *errorOutput) WithArgs() *errorOutput {
	return eo
}

// JSON is a function uses to format response then return as json.
func (eo *errorOutput) JSON() {
	eo.c.Header("Accept-Language", eo.language)
	eo.c.JSON(eo.c.Writer.Status(), eo)
}
