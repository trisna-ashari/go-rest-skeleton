package handler

import (
	"go-rest-skeleton/config"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler is a struct.
type PingHandler struct {
	cf *config.Config
}

// PingResponse is a struct.
type PingResponse struct {
	DB    string `json:"db"`
	Redis string `json:"redis"`
}

// NewPingHandler will initialize Ping handler.
func NewPingHandler(config *config.Config) *PingHandler {
	return &PingHandler{cf: config}
}

// @Summary Ping server
// @Description Check server response.
// @Tags development
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.successOutput
// @Failure 400 {string} middleware.errOutput
// @Failure 404 {object} middleware.errOutput
// @Failure 500 {object} middleware.errOutput
// @Router /api/ping [get]
// Ping will handle ping request.
func (p *PingHandler) Ping(c *gin.Context) {
	var pingData PingResponse

	dbConnection, errDBConn := persistence.NewDBConnection(p.cf.DBConfig)
	if errDBConn != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errDBConn)
		pingData.DB = "Not OK"
	} else {
		defer dbConnection.Close()
		pingData.DB = "OK"
	}

	redisConnection, errRedisConn := persistence.NewRedisConnection(p.cf.RedisConfig)
	if errRedisConn != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errRedisConn)
		pingData.Redis = "Not OK"
	} else {
		_ = redisConnection.Close()
		pingData.Redis = "OK"
	}

	middleware.Formatter(c, pingData, "PONG", nil)
}
