package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
)

func RPC() error {
	global.RPC = types.NewPool(fmt.Sprintf("%s:%s", config.Config.WorkerRPCEndpoint, config.Config.WorkerRPCPort), false)

	return nil
}
