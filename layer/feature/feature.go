package feature

import (
	"chaos-knight/src"
	"context"
	"fmt"
)

type Feature interface {
	Do(ctx context.Context, req *src.Request, mrc *src.Mrc) (ft *src.Ft, err error)
}

const (
	CommonTag = "common"
)

func getFeatureInst(tag string) Feature {
	switch tag {
	case CommonTag:
		return &CommonFeature{}
	default:
		return nil
	}
}

func GetFeature(ctx context.Context, req *src.Request, mrc *src.Mrc, tags []string) (ft *src.Ft, err error) {
	if len(tags) != 1 {
		err = fmt.Errorf("len(tags) must be 1")
		return
	}
	inst := getFeatureInst(tags[0])
	ft, err = inst.Do(ctx, req, mrc)
	return
}
