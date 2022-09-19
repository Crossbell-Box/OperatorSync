package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/rpc/service"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"net"
	"net/rpc"
)

func RPC() error {

	err := rpc.RegisterName(commonConsts.RPCSETTINGS_BaseServiceName, new(service.WorkerRPC))
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Config.WorkerRPCPort))
	if err != nil {
		return err
	}

	go func() {
		global.Logger.Debug("Start waiting for RPC connections...")

		for {
			// Should listen forever

			conn, err := listener.Accept() // Block till next incoming connection
			if err != nil {
				global.Logger.Fatalf("Failed to accept connections with error: %s", err.Error())
			}

			// Start a new session to serve
			go func() {
				global.Logger.Debug("A new RPC connection starting...")

				conn := conn // Copy value to prevent shared memory change

				rpc.ServeConn(conn)

				// If ended
				global.Logger.Debug("An RPC connection closed.")
			}()
		}
	}()

	return nil
}
