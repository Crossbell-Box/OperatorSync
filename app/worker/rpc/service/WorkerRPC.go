package service

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/rpc/jobs"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

type WorkerRPC struct{}

func (rpc *WorkerRPC) Validate(request commonTypes.ValidateRequest, response *commonTypes.ValidateResponse) error {
	jobs.ValidateAccounts(&request, response)
	return nil
}

func (rpc *WorkerRPC) OnChain(request commonTypes.OnChainRequest, response *commonTypes.OnChainResponse) error {
	jobs.OnChainNotes(&request, response)
	return nil
}
