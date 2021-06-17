package rerank

import (
	"chaos-knight/src"
	"context"
	"fmt"
)

type Rerank interface {
	Do(ctx context.Context, req *src.Request, input *src.Rk) (ret *src.Rk, err error)
}

const (
	Idea1Tag = "idea1"
	Idea2Tag = "idea2"
	Idea3Tag = "idea3"
)

func getRerankInst(tag string) Rerank {
	switch tag {
	case Idea1Tag:
		return &Idea1{}
	case Idea2Tag:
		return &Idea2{}
	case Idea3Tag:
		return &Idea3{}
	default:
		return nil
	}
}

func DoRerank(ctx context.Context, req *src.Request, rk *src.Rk, tags []string) (ret *src.Rk, err error) {
	for _, tag := range tags {
		inst := getRerankInst(tag)
		rk, _ = inst.Do(ctx, req, rk)
	}
	if len(rk.Iss) == 0 {
		err = fmt.Errorf("candidates is empty after rerank")
		return
	}
	ret = rk
	return
}
