package filter

import (
	"chaos-knight/src"
	"context"
)

type Recommended struct {
}

func (inst *Recommended) Do(ctx context.Context, req *src.Request, ft *src.Ft) (ret *src.Ft, err error) {
	recommendedMap := map[int64]struct{}{4002: {}, 5003: {}} // already recommended ids (personally)
	for id := range ft.Fm {
		if _, ok := recommendedMap[id]; ok {
			delete(ft.Fm, id)
		}
	}
	return ft, nil
}
