package metrics

import (
	"context"
	"time"

	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

func UpdateValue(me *output.Metrics) {

	me.Counter += 1
	me.UpdatedAt = time.Now()
}

func Update(ctx context.Context, feature string, repo *output.IMetricsRepository) (bool, error) {

	mtrs, err := repo.Fetch(ctx, feature)
	if err != nil {
		return false, err
	}

	UpdateValue(mtrs)

	res, err := repo.Store(ctx, mtrs)
	if err != nil || !res {
		return false, err
	}
	return true, nil

}
