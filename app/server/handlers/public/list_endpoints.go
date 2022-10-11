package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListEndpoints(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []string{
		"GET    /            - List all available endpoints",
		"GET    /healthcheck - Status and healthcheck",

		"GET    /v1/:character                                      - Check character status",
		"POST   /v1/:character                                      - Activate character",
		"DELETE /v1/:character                                      - Deactivate character",
		"GET    /v1/:character/account                              - List accounts of a specified character",
		"POST   /v1/:character/account/bind/:platform/:username     - Bind new platform account",
		"POST   /v1/:character/account/sync/:platform/:username     - Manual sync an account",
		"DELETE /v1/:character/account/unbind/:platform/:username   - Unbind platform account",
		"GET    /v1/:character/media                                - Get media of a specified character",
		"GET    /v1/feed/:platform/:username                        - Get feeds of a specified account",
	})
}
