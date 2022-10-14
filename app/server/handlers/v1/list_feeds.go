package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func ListSingleAccountFeeds(ctx *gin.Context) {
	// Parse request params
	reqPlatform := ctx.Param("platform")
	reqUsername := ctx.Param("username")

	if _, ok := commonConsts.SUPPORTED_PLATFORM[reqPlatform]; !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Platform not supported",
		})
		return
	}

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

	reqLimitString := ctx.DefaultQuery("limit", fmt.Sprintf("%d", consts.MAX_FEED_LIMIT_PER_PAGE))
	reqLimit, err := strconv.ParseUint(reqLimitString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Invalid limit (should be uint)",
		})
		return
	}
	if reqLimit > consts.MAX_FEED_LIMIT_PER_PAGE {
		reqLimit = consts.MAX_FEED_LIMIT_PER_PAGE
	}

	// Validate request
	var account models.Account
	if err := global.DB.First(&account, "platform = ? AND username = ?", reqPlatform, reqUsername).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Account not found",
		})
		return
	}

	// Get requests
	var feeds []models.Feed

	// Use redis hash, and clear entire key when feeds update
	cacheBaseKey := fmt.Sprintf("%s:%s:%d", consts.CACHE_PREFIX, "feeds", account.ID)
	cacheQueryKey := fmt.Sprintf("%d:", reqLimit)

	// Check cache
	getCacheCtx, getCacheCtxCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer getCacheCtxCancel()
	if nowTime.Sub(reqBeforeTime) < 1*time.Second {
		cacheQueryKey += "latest" // Regard as realtime
	} else {
		cacheQueryKey += fmt.Sprintf("%s", reqBeforeTime)
	}
	if exist, err := commonGlobal.Redis.HExists(getCacheCtx, cacheBaseKey, cacheQueryKey).Result(); err != nil {
		global.Logger.Error("Unable to check feeds cache")
	} else if exist {
		if feedsBytes, err := commonGlobal.Redis.HGet(getCacheCtx, cacheBaseKey, cacheQueryKey).Bytes(); err != nil {
			global.Logger.Error("Unable to get feeds cache")
		} else if err = json.Unmarshal(feedsBytes, &feeds); err != nil {
			global.Logger.Error("Unable to parse feeds cache")
			// Clear invalid cache
			commonGlobal.Redis.HDel(getCacheCtx, cacheBaseKey, cacheQueryKey)
		} else {
			// Successfully parsed

			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Records found",
				"result":  feeds,
			})

			return
		}
	}

	// Get data from database
	var feedWithPlatform models.Feed
	feedWithPlatform.Platform = account.Platform // Cannot be initialized with platform specified

	global.DB.
		Scopes(models.FeedTable(feedWithPlatform)).
		Order("created_at DESC").
		Where("account_id = ? AND collected_at < ?", account.ID, reqBeforeTime).
		Limit(int(reqLimit)).
		Find(&feeds)

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "records found",
		"result":  feeds,
	})

	// Set cache
	/*
		setCacheCtx := context.Background()
		if feedsBytes, err := json.Marshal(&feeds); err != nil {
			// WTF?
			global.Logger.Error("Failed to parse feeds")
		} else {
			commonGlobal.Redis.HSet(setCacheCtx, cacheBaseKey, cacheQueryKey, feedsBytes)
		}

	*/

}
