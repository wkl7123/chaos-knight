package feature

import (
	"chaos-knight/src"
	"context"
)

type CommonFeature struct {
}

func (inst *CommonFeature) Do(ctx context.Context, req *src.Request, mrc *src.Mrc) (ft *src.Ft, err error) {
	m := make(map[int64]*src.Feature)
	for id, sources := range mrc.ItemIdMap {
		m[id] = &src.Feature{
			ItemInfo: &src.ItemInfo{}, // 真实环境可以从全局的map中获取
			UserInfo: &src.UserInfo{}, // 真实环境需要实时构建
			Sources:  sources,
		}
	}
	ft = &src.Ft{
		Fm: m,
	}
	return
}
