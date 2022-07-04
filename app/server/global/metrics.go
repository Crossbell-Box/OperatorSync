package global

import "github.com/rcrowley/go-metrics"

var (
	MetricsDispatchedWorkCount metrics.Counter
	MetricsSucceededWorkCount  metrics.Counter
	MetricsFailedWorkCount     metrics.Counter
)
