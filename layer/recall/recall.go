package recall

import (
	"chaos-knight/src"
	"context"
	"fmt"
	"sync"
)

type Recall interface {
	Do(ctx context.Context, req *src.Request) (ret *src.Rc, err error)
}

const (
	U2r2rTag = "u2r2r"
	U2rTag   = "u2r"
	HotTag   = "hot"
)

func getRecallInst(tag string) Recall {
	switch tag {
	case U2r2rTag:
		return &U2R2R{}
	case U2rTag:
		return &U2R{}
	case HotTag:
		return &Hot{}
	default:
		return nil
	}
}

// concurrent get recall result and merge
func GetCandidates(ctx context.Context, req *src.Request, tags []string) (mrc *src.Mrc, err error) {
	m := make(map[string][]int64)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, tag := range tags {
		wg.Add(1)
		defer wg.Done()
		inst := getRecallInst(tag)
		if inst != nil {
			if rc, err := inst.Do(ctx, req); err != nil && rc != nil {
				mutex.Lock()
				m[rc.Tag] = rc.ItemIds
				mutex.Unlock()
			}
		}
	}

	m1 := make(map[int64][]string)
	for tag, list := range m {
		for _, id := range list {
			m1[id] = append(m1[id], tag)
		}
	}

	if len(m1) == 0 {
		err = fmt.Errorf("recall empty")
		return
	}

	mrc = &src.Mrc{
		TagListMap: m,
		ItemIdMap:  m1,
	}

	return
}
