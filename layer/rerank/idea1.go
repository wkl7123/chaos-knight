package rerank

import (
	"chaos-knight/src"
	"context"
)

type Idea1 struct {
}

func (inst *Idea1) Do(ctx context.Context, req *src.Request, input *src.Rk) (ret *src.Rk, err error) {
	return input, nil
}
