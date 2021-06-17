package filter

import (
	"chaos-knight/src"
	"context"
	"fmt"
)

type Filter interface {
	Do(ctx context.Context, req *src.Request, ft *src.Ft) (ret *src.Ft, err error)
}

const (
	DislikeTag     = "dislike"
	RecommendedTag = "recommended"
)

func getFilterInst(tag string) Filter {
	switch tag {
	case DislikeTag:
		return &Dislike{}
	case RecommendedTag:
		return &Recommended{}
	default:
		return nil
	}
}

func DoFilter(ctx context.Context, req *src.Request, ft *src.Ft, tags []string) (ret *src.Ft, err error) {
	for _, tag := range tags {
		inst := getFilterInst(tag)
		ft, _ = inst.Do(ctx, req, ft)
	}
	if len(ft.Fm) == 0 {
		err = fmt.Errorf("empty candidates after filter")
		return
	}
	ret = ft
	return
}
