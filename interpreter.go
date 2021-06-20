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
	"sort"
)

type Interpreter struct {
	conf src.Conf
}

func (it *Interpreter) Prepare() {
	for _, item := range it.conf {
		for _, option := range item.Layers.Recall {
			sort.Strings(option.Choice)
		}
		for _, option := range item.Layers.Feature {
			sort.Strings(option.Choice)
		}
		for _, option := range item.Layers.Filter {
			sort.Strings(option.Choice)
		}
		for _, option := range item.Layers.Feature {
			sort.Strings(option.Choice)
		}
		for _, option := range item.Layers.Rerank {
			sort.Strings(option.Choice)
		}
	}
}

func (it *Interpreter) Do(ctx context.Context, req *src.Request, scene string) (rk *src.Rk, trackInfo map[string]string, err error) {
	trackInfo = make(map[string]string)
	if _, ok := it.conf[scene]; !ok {
		return nil, trackInfo, fmt.Errorf("scene not supported: %s", scene)
	}

	layers := it.conf[scene].Layers
	recallLayerOption := src.Pick(layers.Recall)
	trackInfo["recall"] = recallLayerOption.String()
	mrc, err := recall.GetCandidates(ctx, req, recallLayerOption.Choice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	featureLayerOption := src.Pick(layers.Feature)
	trackInfo["feature"] = featureLayerOption.String()
	ft, err := feature.GetFeature(ctx, req, mrc, featureLayerOption.Choice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	filterLayerOption := src.Pick(layers.Filter)
	trackInfo["filter"] = filterLayerOption.String()
	ft, err = filter.DoFilter(ctx, req, ft, filterLayerOption.Choice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	modelLayerOption := src.Pick(layers.Model)
	trackInfo["model"] = modelLayerOption.String()
	rk, err = model.DoPredict(ctx, req, ft, modelLayerOption.Choice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	rerankLayerOption := src.Pick(layers.Rerank)
	trackInfo["rerank"] = rerankLayerOption.String()
	rk, err = rerank.DoRerank(ctx, req, rk, rerankLayerOption.Choice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	return
}
