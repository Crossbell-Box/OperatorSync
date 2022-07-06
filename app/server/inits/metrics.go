package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/rcrowley/go-metrics"
)

func Metrics() error {

	global.MetricsDispatchedWorkCount = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_PREFIX+"dispatched", global.MetricsDispatchedWorkCount); err != nil {
		global.Logger.Error("Failed to register Dispatched Works metrics: ", err.Error())
		return err
	}

	global.MetricsSucceededWorkCount = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_PREFIX+"succeeded", global.MetricsSucceededWorkCount); err != nil {
		global.Logger.Error("Failed to register Succeeded Works metrics: ", err.Error())
		return err
	}

	global.MetricsFailedWorkCount = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_PREFIX+"failed", global.MetricsFailedWorkCount); err != nil {
		global.Logger.Error("Failed to register Failed Works metrics: ", err.Error())
		return err
	}

	defer metrics.Unregister(consts.WORK_COUNTS_METRICS_PREFIX)
	return nil
}
