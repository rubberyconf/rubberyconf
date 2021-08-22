package metrics

import (
	"context"
	"sync"
	"time"
)

type MongoMetrics struct {
	Feature   string
	Counter   int64
	UpdatedAt time.Time
	ctx       context.Context
}

type Metrics struct {
	//store chan MongoMetrics
}

var (
	metrics     *Metrics
	metricsOnce sync.Once
)

func NewMetrics() *Metrics {

	metricsOnce.Do(func() {
		metrics = new(Metrics)
		//metrics.store = make(chan MongoMetrics)
	})
	return metrics
}

func (metricRegister *MongoMetrics) Update() {

	metricRegister.Counter += 1
	metricRegister.UpdatedAt = time.Now()
}

func (metric *Metrics) Update(ctx context.Context, feature string) (*MongoMetrics, error) {

}

func (metric *Metrics) Remove(ctx context.Context, feature string) (bool, error) {

}
