package public

import (
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListPlatforms(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonConsts.SUPPORTED_PLATFORM)
}
