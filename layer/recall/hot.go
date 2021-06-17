package recall

import (
	"chaos-knight/src"
	"context"
)

type Hot struct {
}

func (inst *Hot) Do(ctx context.Context, req *src.Request) (ret *src.Rc, err error) {
	ret = &src.Rc{
		Tag:     HotTag,
		ItemIds: []int64{5001, 5002, 5006},
	}
	return
}
