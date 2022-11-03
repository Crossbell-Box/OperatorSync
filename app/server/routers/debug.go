package routers

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/handlers/debug"
	"github.com/gin-gonic/gin"
)

func DebugEndpoints(rg *gin.RouterGroup) {
	rg.POST("/job/resume-paused-accounts", debug.TriggerJobResumePausedAccounts)
}
