package global

import "github.com/rcrowley/go-metrics"

type gMetrics struct {
	Work struct {
		Dispatched metrics.Counter
		Succeeded  metrics.Counter
		Failed     metrics.Counter
	}
}

var Metrics gMetrics
