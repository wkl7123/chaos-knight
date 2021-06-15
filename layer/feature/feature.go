package feature

import (
	"chaos-knight/src"
	"context"
)

type Feature interface {
	Do(ctx context.Context, sid int64, req *src.Request, mrc *src.Mrc) (ft *src.Ft, err error)
}
