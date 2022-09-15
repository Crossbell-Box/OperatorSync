package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"net/rpc"
)

func RPC() error {
	var err error
	global.RPC, err = rpc.Dial("tcp", fmt.Sprintf("%s:%s", config.Config.WorkerRPCEndpoint, config.Config.WorkerRPCPort))
	if err != nil {
		return fmt.Errorf("failed to connect to worker with error: %v", err)
	}

	return nil
}
