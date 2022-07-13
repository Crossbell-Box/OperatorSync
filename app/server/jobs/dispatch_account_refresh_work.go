package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	"time"
)

func StartDispatchAccountRefreshWork() {
	go func() {
		t := time.NewTicker(1 * time.Hour)
		for {
			select {
			case <-t.C:
				dispatchAccountRefreshWork()
			}
		}
	}()
}

func dispatchAccountRefreshWork() {
	global.Logger.Debug("Start dispatching account refresh works...")

	nowTime := time.Now()

	// Accounts need update
	var charactersNeedUpdate []models.Character

	global.DB.Find(&charactersNeedUpdate, "account_last_updated_at < ?", nowTime.Add(-consts.ACCOUNT_AUTOREFRESH_INTERVAL))

	if len(charactersNeedUpdate) == 0 {
		global.Logger.Debug("No characters need update currently")
		return
	}

	for _, character := range charactersNeedUpdate {
		_, _ = utils.RefreshAccountForCharacter(character.CrossbellCharacter) // Ignore results & errors
	}

	global.Logger.Debugf("Accounts refresh work finished for %d characters.", len(charactersNeedUpdate))
}
