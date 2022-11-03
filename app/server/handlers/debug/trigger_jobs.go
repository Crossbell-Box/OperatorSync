package debug

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/jobs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TriggerJobResumePausedAccounts(ctx *gin.Context) {
	go jobs.TryToResumeAllPausedAccounts()

	ctx.JSON(http.StatusAccepted, "Start ResumePausedAccounts job.")
}
