package model

import (
	"chaos-knight/src"
	"context"
)

type Model interface {
	Init(conf *RecSysModelConf) (err error)
	Predict(ctx context.Context, ft *src.Ft) (rank *src.Rk, err error)
}

type RecSysModelConf struct {
	Repo    string
	Version string
}
