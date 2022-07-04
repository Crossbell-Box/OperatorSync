package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/rcrowley/go-metrics"
)

func Metrics() error {

	global.MetricsDispatchedWorkCount = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_SCOPE, global.MetricsDispatchedWorkCount); err != nil {
		global.Logger.Error("Failed to register Dispatched Works metrics: ", err)
		return err
	}

	global.MetricsSucceededWorkCount = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_SCOPE, global.MetricsSucceededWorkCount); err != nil {
		global.Logger.Error("Failed to register Succeeded Works metrics: ", err)
		return err
	}

	global.MetricsFailedWorkCount = metrics.NewCounter()
	if err := metrics.Register(consts.WORK_COUNTS_METRICS_SCOPE, global.MetricsFailedWorkCount); err != nil {
		global.Logger.Error("Failed to register Failed Works metrics: ", err)
		return err
	}

	defer metrics.Unregister(consts.WORK_COUNTS_METRICS_SCOPE)
	return nil
}
