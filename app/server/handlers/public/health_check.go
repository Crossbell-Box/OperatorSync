package public

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type healthStatus struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}

func HealthCheck(ctx *gin.Context) {
	if config.Config.IsMainServer {
		if time.Now().Sub(config.Status.Jobs.FeedCollectLastRun) > 2*consts.JOBS_INTERVAL_FEED_COLLECT {
			ctx.JSON(http.StatusInternalServerError, healthStatus{
				OK:      false,
				Message: "Feed collect work not running",
			})
			return
		}

		if time.Now().Sub(config.Status.Jobs.ResumePausedAccountsLastRun) > 2*consts.JOBS_INTERVAL_RESUME_PAUSED_ACCOUNTS {
			ctx.JSON(http.StatusInternalServerError, healthStatus{
				OK:      false,
				Message: "Resume paused account work not running",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, healthStatus{
		OK: true,
	})
}
