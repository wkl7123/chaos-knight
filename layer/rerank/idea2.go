package rerank

import (
	"chaos-knight/src"
	"context"
)

type Idea2 struct {
}

func (inst *Idea2) Do(ctx context.Context, req *src.Request, input *src.Rk) (ret *src.Rk, err error) {
	return input, nil
}
