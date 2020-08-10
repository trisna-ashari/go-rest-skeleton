package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// LoggerOptions is a struct to store options for SetLogger.
type LoggerOptions struct {
	AllowSetting bool
}

// Config is a struct to store config for SetLogger.
type Config struct {
	Logger         *zerolog.Logger
	UTC            bool
	SkipPath       []string
	SkipPathRegexp *regexp.Regexp
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func makeSkip(newConfig Config) map[string]struct{} {
	var skip map[string]struct{}
	if length := len(newConfig.SkipPath); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range newConfig.SkipPath {
			skip[path] = struct{}{}
		}
	}
	return skip
}

func makeSubLogger(newConfig Config) zerolog.Logger {
	var subLog zerolog.Logger
	if newConfig.Logger == nil {
		subLog = log.Logger
	} else {
		subLog = *newConfig.Logger
	}

	return subLog
}

func getRequestHeader(c *gin.Context) []byte {
	var headerBytes []byte
	var rawHeader = make(map[string]interface{})
	for headerKey, headerValue := range c.Request.Header {
		if len(headerValue) == 1 {
			rawHeader[headerKey] = headerValue[0]
		} else {
			rawHeader[headerKey] = headerValue
		}
	}

	headerBytes, err := json.Marshal(rawHeader)
	if err != nil {
		return headerBytes
	}

	return headerBytes
}

func getRequestForm(c *gin.Context) []byte {
	var formBytes []byte
	var rawForm = make(map[string]interface{})
	formData, _ := c.MultipartForm()
	if formData != nil {
		for formKey, formValue := range formData.Value {
			if len(formValue) == 1 {
				rawForm[formKey] = formValue[0]
			}
		}
	}

	formBytes, err := json.Marshal(rawForm)
	if err != nil {
		return formBytes
	}

	return formBytes
}

func getRequestBody(c *gin.Context) []byte {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	var rawBody map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &rawBody); err != nil {
		bodyBytes = []byte("{}")
	}

	return bodyBytes
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func getResponseBody(c *gin.Context) []byte {
	w := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()
	if c.Writer.Status() >= http.StatusBadRequest {
		return []byte("{}")
	}
	return w.body.Bytes()
}

func printLogger(c *gin.Context, l zerolog.Logger, m string) {
	switch {
	case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
		{
			l.Warn().
				Msg(m)
		}
	case c.Writer.Status() >= http.StatusInternalServerError:
		{
			l.Error().
				Msg(m)
		}
	default:
		l.Info().
			Msg(m)
	}
}

// SetLogger is a middleware function uses to log all incoming request and print it to console.
func SetLogger(options LoggerOptions, config ...Config) gin.HandlerFunc {
	var newConfig Config
	if len(config) > 0 {
		newConfig = config[0]
	}

	skip := makeSkip(newConfig)
	subLog := makeSubLogger(newConfig)

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		requestHeader := getRequestHeader(c)
		requestBody := getRequestBody(c)
		requestForm := getRequestForm(c)
		requestID := c.Writer.Header().Get("X-Request-Id")
		responseBody := getResponseBody(c)

		c.Next()

		track := options.AllowSetting

		if _, ok := skip[path]; ok {
			track = false
		}

		if track &&
			newConfig.SkipPathRegexp != nil &&
			newConfig.SkipPathRegexp.MatchString(path) {
			track = false
		}

		if track {
			end := time.Now()
			latency := end.Sub(start)
			if newConfig.UTC {
				end = end.UTC()
			}

			msg := "Request"
			if len(c.Errors) > 0 {
				msg = c.Errors.String()
			}

			dumpLogger := subLog.With().
				Int("status", c.Writer.Status()).
				Str("method", c.Request.Method).
				Str("path", path).
				Str("ip", c.ClientIP()).
				Dur("latency", latency).
				Str("user-agent", c.Request.UserAgent()).
				Str("request-id", requestID).
				RawJSON("headers", requestHeader).
				RawJSON("request-payload", requestBody).
				RawJSON("request-form", requestForm).
				RawJSON("response", responseBody).
				Logger()

			printLogger(c, dumpLogger, msg)
		}
	}
}
