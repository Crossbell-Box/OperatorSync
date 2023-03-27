package utils

import (
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"gorm.io/gorm"
)

func FeedOnChainDispatchWork(account *models.Account, feeds []models.Feed) {
	if IsAccountOnChainPaused(account) {
		global.Logger.Errorf("OnChain process has been paused for account %s#%d with reason: %s", account.Platform, account.ID, account.OnChainPauseMessage)
		return
	}

	for _, feed := range feeds {

		// Check if deps is on chain
		if feed.ForURI != "" {
			var forFeed models.Feed
			// Get from database
			err := global.DB.Scopes(models.FeedTable(models.Feed{
				Feed: types.Feed{
					Platform: account.Platform,
				},
			})).First(&forFeed, "link = ?", feed.ForURI).Error
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					// Something is wrong
					global.Logger.Errorf("Failed to get for link from database with error: %v", err)
				}
			} else {
				// Got it, check if is valid
				if forFeed.CharacterID > 0 && forFeed.NoteID > 0 {
					// Is valid
					feed.ForCharacterID = forFeed.CharacterID
					feed.ForNoteID = forFeed.NoteID
				} else if forFeed.Transaction == "" {
					// Invalid
					msg := fmt.Sprintf("Depending note with link %s is not on chain yet", feed.ForURI)
					global.Logger.Errorf(msg)
					// Pause account
					AccountOnChainPause(account, msg)
					break
				}
			}
		}

		ipfsUri, tx, characterId, noteId, err, _ := OneFeedOnChain(account, &feed)
		if err != nil {
			global.Logger.Errorf("Failed to push all feeds on chain with error: %s", err.Error())
			break
		} else {
			// Accumulate notes_count
			global.DB.Exec("UPDATE accounts SET notes_count = notes_count + 1 WHERE id = ?", account.ID)

			// Update feed save into database
			if err = global.DB.Scopes(models.FeedTable(models.Feed{
				Feed: types.Feed{
					Platform: account.Platform,
				},
			})).Updates(&models.Feed{
				ID: feed.ID,
				OnChainData: types.OnChainData{
					IPFSUri:     ipfsUri,
					Transaction: tx,
					CharacterID: characterId,
					NoteID:      noteId,
				},
			}).Error; err != nil {
				global.Logger.Errorf("Failed to save updated feed %v", feed)
			}
		}

	}
}
