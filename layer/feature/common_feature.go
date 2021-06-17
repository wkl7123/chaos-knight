package feature

import (
	"chaos-knight/src"
	"context"
)

type CommonFeature struct {
}

func (inst *CommonFeature) Do(ctx context.Context, req *src.Request, mrc *src.Mrc) (ft *src.Ft, err error) {
	m := make(map[int64]*src.Feature)
	// ...
	ft = &src.Ft{
		Fm: m,
	}
	return
}
