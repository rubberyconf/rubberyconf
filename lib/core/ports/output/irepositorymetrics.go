package output

import (
	"context"
	"time"
)

type IMetricsRepository interface {
	Fetch(ctx context.Context, feature string) (*Metrics, error)

	Store(ctx context.Context, metricRegister *Metrics) (bool, error)

	Remove(ctx context.Context, feature string) (bool, error)
}

type Metrics struct {
	Feature   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Counter   int64
}
