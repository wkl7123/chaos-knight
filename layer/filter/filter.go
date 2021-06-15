package filter

import (
	"chaos-knight/src"
	"context"
)

type Filter interface {
	Do(ctx context.Context, sid int64, req *src.Request, ft *src.Ft) (ret *src.Ft, err error)
}
