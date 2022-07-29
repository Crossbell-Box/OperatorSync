package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ListMedias(ctx *gin.Context) {
	// Parse request params
	reqCharacterID := ctx.Param("character")

	nowTime := time.Now()
	reqBeforeTimeString := ctx.DefaultQuery("before", nowTime.Format(time.RFC3339))
	reqBeforeTime, err := time.Parse(time.RFC3339, reqBeforeTimeString)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Invalid time format (please use RFC3339)",
		})
		return
	}
	//if nowTime.Before(reqBeforeTime) {
	//	reqBeforeTime = nowTime
	//}

	reqLimitString := ctx.DefaultQuery("limit", fmt.Sprintf("%d", consts.MAX_MEDIA_LIMIT_PER_PAGE))
	reqLimit, err := strconv.ParseUint(reqLimitString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Invalid limit (should be uint)",
		})
		return
	}
	if reqLimit > consts.MAX_MEDIA_LIMIT_PER_PAGE {
		reqLimit = consts.MAX_MEDIA_LIMIT_PER_PAGE
	}

	// Get requests
	var medias []models.Media

	// Use redis hash, and clear entire key when medias update
	cacheBaseKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "medias", reqCharacterID)
	cacheQueryKey := fmt.Sprintf("%d:", reqLimit)

	// Check cache
	getCacheCtx, getCacheCtxCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer getCacheCtxCancel()
	// NowTime is much more accurate than parsed time, so we regard time diff smaller than 1 sec is realtime
	if nowTime.Sub(reqBeforeTime) < 1*time.Second {
		cacheQueryKey += "latest"
	} else {
		cacheQueryKey += fmt.Sprintf("%s", reqBeforeTime)
	}
	if exist, err := global.Redis.HExists(getCacheCtx, cacheBaseKey, cacheQueryKey).Result(); err != nil {
		global.Logger.Error("Unable to check medias cache")
	} else if exist {
		if mediasBytes, err := global.Redis.HGet(getCacheCtx, cacheBaseKey, cacheQueryKey).Bytes(); err != nil {
			global.Logger.Error("Unable to get medias cache")
		} else if err = json.Unmarshal(mediasBytes, &medias); err != nil {
			global.Logger.Error("Unable to parse medias cache")
			// Clear invalid cache
			global.Redis.HDel(getCacheCtx, cacheBaseKey, cacheQueryKey)
		} else {
			// Successfully parsed

			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Records found",
				"result":  medias,
			})

			return
		}
	}

	global.DB.
		Order("created_at DESC").
		Where("crossbell_character_id = ? AND created_at < ?", reqCharacterID, reqBeforeTime).
		Limit(int(reqLimit)).
		Find(&medias)

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "records found",
		"result":  medias,
	})

	// Set cache
	setCacheCtx := context.Background()
	if mediasBytes, err := json.Marshal(&medias); err != nil {
		// WTF?
		global.Logger.Error("Failed to parse medias")
	} else {
		global.Redis.HSet(setCacheCtx, cacheBaseKey, cacheQueryKey, mediasBytes)
	}

}
