package main

import (
	"chaos-knight/layer/feature"
	"chaos-knight/layer/filter"
	"chaos-knight/layer/model"
	"chaos-knight/layer/recall"
	"chaos-knight/layer/rerank"
	"chaos-knight/src"
	"context"
	"fmt"
)

type Interpreter struct {
	conf src.Conf
}

func (it *Interpreter) Do(ctx context.Context, req *src.Request, scene string) (rk *src.Rk, err error) {
	if _, ok := it.conf[scene]; !ok {
		return nil, fmt.Errorf("scene not supported: %s", scene)
	}

	layers := it.conf[scene].Layers
	recallLayerOption := src.Pick(layers.Recall)
	mrc, err := recall.GetCandidates(ctx, req, recallLayerOption.Choice)
	if err != nil {
		return
	}

	featureLayerOption := src.Pick(layers.Feature)
	ft, err := feature.GetFeature(ctx, req, mrc, featureLayerOption.Choice)
	if err != nil {
		return
	}

	filterLayerOption := src.Pick(layers.Filter)
	ft, err = filter.DoFilter(ctx, req, ft, filterLayerOption.Choice)
	if err != nil {
		return
	}

	modelLayerOption := src.Pick(layers.Model)
	rk, err = model.DoPredict(ctx, req, ft, modelLayerOption.Choice)
	if err != nil {
		return
	}
	rerankLayerOption := src.Pick(layers.Rerank)
	rk, err = rerank.DoRerank(ctx, req, rk, rerankLayerOption.Choice)
	return
}
