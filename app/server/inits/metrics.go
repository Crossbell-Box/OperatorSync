package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/rcrowley/go-metrics"
)

func Metrics() error {

	global.Metrics.Work.Dispatched = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_PREFIX+"dispatched", global.Metrics.Work.Dispatched); err != nil {
		global.Logger.Error("Failed to register Dispatched Works metrics: ", err.Error())
		return err
	}

	global.Metrics.Work.Succeeded = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_PREFIX+"succeeded", global.Metrics.Work.Succeeded); err != nil {
		global.Logger.Error("Failed to register Succeeded Works metrics: ", err.Error())
		return err
	}

	global.Metrics.Work.Failed = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_PREFIX+"failed", global.Metrics.Work.Failed); err != nil {
		global.Logger.Error("Failed to register Failed Works metrics: ", err.Error())
		return err
	}

	defer metrics.Unregister(consts.WORK_COUNTS_METRICS_PREFIX)
	return nil
}
