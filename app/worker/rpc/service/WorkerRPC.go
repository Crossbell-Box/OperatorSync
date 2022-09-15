package service

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/rpc/jobs"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

type WorkerRPC struct{}

func (rpc *WorkerRPC) Validate(request commonTypes.ValidateRequest, response *commonTypes.ValidateResponse) error {
	response = jobs.ValidateAccounts(&request)
	return nil
}

func (rpc *WorkerRPC) OnChain(request commonTypes.OnChainRequest, response *commonTypes.OnChainRespond) error {
	response = jobs.OnChainNotes(&request)
	return nil
}
