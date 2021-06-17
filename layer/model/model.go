package model

import (
	"chaos-knight/src"
	"context"
	"fmt"
	"path"
	"sync"
)

var Map sync.Map

func init() {
	modelLr1 := &LR{}
	c1 := &RecSysModelConf{
		Repo:    "lr",
		Version: "v0.0.1",
	}
	_ = modelLr1.Init(c1)
	Map.Store(c1.FormKey(), modelLr1)

	modelLr2 := &LR{}
	c2 := &RecSysModelConf{
		Repo:    "lr",
		Version: "v0.0.2",
	}
	_ = modelLr1.Init(c2)
	Map.Store(c2.FormKey(), modelLr2)
}

type Model interface {
	Init(conf *RecSysModelConf) (err error)
	Predict(ctx context.Context, ft *src.Ft) (rank *src.Rk, err error)
}

type RecSysModelConf struct {
	Repo    string
	Version string
}

func (c *RecSysModelConf) GetModelPath() string {
	return path.Join("/data/model", c.Repo, c.Version)
}

func (c *RecSysModelConf) FormKey() string {
	return c.Repo + "@" + c.Version
}

func GetModel(tags []string) (m Model, err error) {
	if len(tags) != 1 {
		err = fmt.Errorf("len(tags) must be 1")
		return
	}
	inst, ok := Map.Load(tags[0])
	if !ok {
		err = fmt.Errorf("model not exist")
		return
	}
	if m, ok = inst.(Model); !ok {
		err = fmt.Errorf("failed to transfer model")
		return
	}
	return
}

func DoPredict(ctx context.Context, req *src.Request, ft *src.Ft, tags []string) (rk *src.Rk, err error) {
	m, err := GetModel(tags)
	if err != nil {
		return
	}
	rk, err = m.Predict(ctx, ft)
	return
}
