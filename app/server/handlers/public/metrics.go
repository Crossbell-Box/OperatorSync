package public

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Metrics(ctx *gin.Context) {
	// Try cache
	cacheKey := fmt.Sprintf("%s:metrics", consts.CACHE_PREFIX)
	if cacheExist, err := commonGlobal.Redis.Exists(context.Background(), cacheKey).Result(); err != nil {
		// Cannot check cache
		global.Logger.Errorf("Failed to check metrics cache with error: %s", err.Error())
	} else if cacheExist > 0 {
		// Cache exist
		if cacheBytes, err := commonGlobal.Redis.Get(context.Background(), cacheKey).Bytes(); err != nil {
			// Cannot get cache
			global.Logger.Errorf("Failed to get metrics cache with error: %s", err.Error())
		} else {
			// Success
			// Check style
			var m types.Metrics
			if err := json.Unmarshal(cacheBytes, &m); err != nil {
				// Invalid
				global.Logger.Errorf("Failed to parse cached metrics cache with error: %s", err.Error())
			} else {
				ctx.JSON(http.StatusOK, &m)
				return
			}
		}
	}

	// Get latest metrics
	m := calcMetrics()

	// Return response
	ctx.JSON(http.StatusOK, m)

	// Save to cache
	mBytes, err := json.Marshal(m)
	if err != nil {
		global.Logger.Errorf("Failed to marshall metrics data to JSON with error: %s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Failed to marshall metrics data to JSON",
		})
		return
	}

	// Save to cache
	commonGlobal.Redis.Set(context.Background(), cacheKey, mBytes, consts.METRICS_CACHE_EXPIRE)

}

func calcMetrics() *types.Metrics {
	var m types.Metrics

	m.GeneratedAt = time.Now()

	global.DB.Model(&models.Character{}).Count(&m.Character.Total)
	global.DB.Model(&models.Account{}).Distinct("crossbell_character_id").Count(&m.Character.Valid)

	m.Account.Total = 0
	m.Account.Valid = 0

	m.Platform = make(map[string]types.PlatformMetrics)
	for platformID, _ := range commonConsts.SUPPORTED_PLATFORM {
		var pm types.PlatformMetrics

		global.DB.Scopes(models.FeedTable(models.Feed{
			Feed: types.Feed{
				Platform: platformID,
			},
		})).Count(&pm.Feed)

		global.DB.Model(&models.Account{}).Where("platform = ?", platformID).Count(&pm.Account.Total)
		global.DB.Model(&models.Account{}).Where("platform = ?", platformID).Where("feeds_count > ?", 0).Count(&m.Character.Valid)

		m.Account.Total += pm.Account.Total
		m.Account.Valid += pm.Account.Valid

		m.Platform[platformID] = pm
	}

	return &m
}
