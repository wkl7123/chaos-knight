package model

import (
	"chaos-knight/src"
	"context"
	"sort"
)

type LR struct {
	m map[string]float32
}

func (inst *LR) Init(conf *RecSysModelConf) (err error) {
	//dir := conf.GetModelPath()
	//get lr feature weight from dir, then fill m
	return
}

func (inst *LR) Predict(ctx context.Context, ft *src.Ft) (rank *src.Rk, err error) {
	rank = &src.Rk{}
	for id, item := range ft.Fm {
		var score float32
		fea_1 := "fea_1" + item.ItemFeaStrMap["fea_1"]
		// ... all features
		score += inst.m[fea_1]
		rank.Iss = append(rank.Iss, &src.ItemScore{
			ItemId:     id,
			ModelScore: score,
			Score:      score,
			Feature:    item,
		})
	}
	sort.Sort(rank.Iss)
	return
}
