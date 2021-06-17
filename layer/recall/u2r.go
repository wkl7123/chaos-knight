package recall

import (
	"chaos-knight/src"
	"context"
)

type U2R struct {
}

func (inst *U2R) Do(ctx context.Context, req *src.Request) (ret *src.Rc, err error) {
	ret = &src.Rc{
		Tag:     U2rTag,
		ItemIds: []int64{5007, 5009},
	}
	return
}
