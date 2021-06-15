package rerank

import (
	"chaos-knight/src"
	"context"
)

type Rerank interface {
	Do(ctx context.Context, sid int64, req *src.Request, input *src.Rk) (ret *src.Rk, err error)
}
