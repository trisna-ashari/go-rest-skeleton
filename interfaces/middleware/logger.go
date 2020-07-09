package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

// SetLogger is a middleware function uses to log all incoming request and print it to console.
func SetLogger(options LoggerOptions, config ...Config) gin.HandlerFunc {
	var newConfig Config
	if len(config) > 0 {
		newConfig = config[0]
	}
	var skip map[string]struct{}
	if length := len(newConfig.SkipPath); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range newConfig.SkipPath {
			skip[path] = struct{}{}
		}
	}

	var subLog zerolog.Logger
	if newConfig.Logger == nil {
		subLog = log.Logger
	} else {
		subLog = *newConfig.Logger
	}

	return func(c *gin.Context) {
		var bodyBytes []byte
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		var rawBody map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &rawBody); err != nil {
			bodyBytes = []byte("{}")
		}

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
				RawJSON("payloads", bodyBytes).
				Logger()

			switch {
			case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
				{
					dumpLogger.Warn().
						Msg(msg)
				}
			case c.Writer.Status() >= http.StatusInternalServerError:
				{
					dumpLogger.Error().
						Msg(msg)
				}
			default:
				dumpLogger.Info().
					Msg(msg)
			}
		}
	}
}
