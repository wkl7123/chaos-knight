package rerank

import (
	"chaos-knight/src"
	"context"
)

type Idea3 struct {
}

func (inst *Idea3) Do(ctx context.Context, req *src.Request, input *src.Rk) (ret *src.Rk, err error) {
	return input, nil
}
