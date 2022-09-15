package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
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

	conn, err := listener.Accept()
	if err != nil {
		return err
	}

	go func() {
		rpc.ServeConn(conn)
	}()

	return nil
}
