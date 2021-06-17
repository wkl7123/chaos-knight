package recall

import (
	"chaos-knight/src"
	"context"
)

type U2R2R struct {
}

func (inst *U2R2R) Do(ctx context.Context, req *src.Request) (ret *src.Rc, err error) {
	ret = &src.Rc{
		Tag:     U2r2rTag,
		ItemIds: []int64{5003, 5002},
	}
	return
}
