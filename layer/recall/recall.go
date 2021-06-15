package recall

import (
	"chaos-knight/src"
	"context"
)

type Recall interface {
	Do(ctx context.Context, sid int64, req *src.Request) (ret *src.Rc, err error)
}
