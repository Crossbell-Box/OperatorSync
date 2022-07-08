package public

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type worksCount struct {
	Dispatched int64 `json:"dispatched"`
	Succeeded  int64 `json:"succeeded"`
	Failed     int64 `json:"failed"`
}

type healthStatus struct {
	OK         bool       `json:"ok"`
	WorksCount worksCount `json:"works_count"`
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, healthStatus{
		OK: true,
		WorksCount: worksCount{
			Dispatched: global.Metrics.Work.Dispatched.Count(),
			Succeeded:  global.Metrics.Work.Succeeded.Count(),
			Failed:     global.Metrics.Work.Failed.Count(),
		},
	})
}
