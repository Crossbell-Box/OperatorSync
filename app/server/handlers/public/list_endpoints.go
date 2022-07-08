package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListEndpoints(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"GET  /":            "List all available endpoints",
		"GET  /healthcheck": "Status and healthcheck",

		"GET  /v1/account/:character":         "List accounts of a specified character",
		"POST /v1/account/:character/refresh": "Trigger once account refresh",
		"GET  /v1/feed/:platform/:username":   "Get feeds of a specified account",
		"GET  /v1/media/:character":           "Get media of a specified character",
	})
}
