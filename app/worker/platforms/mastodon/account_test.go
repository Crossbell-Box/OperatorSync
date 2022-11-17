package mastodon

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"go.uber.org/zap"
	"testing"
)

func TestAccount(t *testing.T) {
	// Init deps
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	usernameAtInstance := "candinya@thenomads.social"
	validateString := "candinya@crossbell"

	t.Log(Account(usernameAtInstance, validateString))

}
