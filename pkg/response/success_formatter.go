package response

import (
	"go-rest-skeleton/pkg/translation"

	"github.com/gin-gonic/gin"
)

type successOutput struct {
	c               *gin.Context
	language        string
	SuccessHTTPCode int         `json:"code"`
	Data            interface{} `json:"data"`
	Message         string      `json:"message"`
	Meta            interface{} `json:"meta,omitempty"`
}

// NewSuccess will is a constructor will initialize successOutput.
func NewSuccess(c *gin.Context, data interface{}, message string) *successOutput {
	successHTTPCode := c.Writer.Status()
	successData := data
	if message == "" {
		message = "OK"
	}
	successMessage, language := translation.NewTranslation(c, "success", message, map[string]interface{}{})

	return &successOutput{
		c:               c,
		language:        language,
		SuccessHTTPCode: successHTTPCode,
		Data:            successData,
		Message:         successMessage,
		Meta:            nil,
	}
}

// WithMeta is a function uses to attach meta into successOutput.
func (so *successOutput) WithMeta(meta interface{}) *successOutput {
	so.Meta = meta

	return so
}

// JSON is a function uses to format response then return as json.
func (so *successOutput) JSON() {
	so.c.Header("Accept-Language", so.language)
	so.c.JSON(so.SuccessHTTPCode, so)
}

// XML is a function uses to format response then return as xml.
func (so *successOutput) XML() {
	so.c.Header("Accept-Language", so.language)
	so.c.XML(so.SuccessHTTPCode, so)
}
