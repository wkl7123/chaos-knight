package filter

import (
	"chaos-knight/src"
	"context"
)

type Dislike struct {
}

func (inst *Dislike) Do(ctx context.Context, req *src.Request, ft *src.Ft) (ret *src.Ft, err error) {
	dislikeTypes := map[string]struct{}{"type1": {}, "type2": {}} // user dislike type (personally)
	for id, item := range ft.Fm {
		ty := item.ItemFeaStrMap["type"]
		if _, ok := dislikeTypes[ty]; ok {
			delete(ft.Fm, id)
		}
	}
	return ft, nil
}
