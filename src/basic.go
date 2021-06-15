package src

// recall
type Rc struct {
	Tag     string
	ItemIds []int64
}

// multi recall
type Mrc struct {
	TagListMap map[string][]int64
	ItemIdMap  map[int64][]string
}

// feature
type Ft struct {
	//Landing int64
	//Ctx     context.Context
	Fm map[int64]*Feature
}

// rank
type Rk struct {
	Iss *ItemScoreSlice
}

type ItemInfo struct {
	ItemFeaStrMap map[string]string
	ItemFeaFltMap map[string]float32
}

type UserInfo struct {
	UserFeaStrMap map[string]string
	UserFeaFltMap map[string]float32
}

type Feature struct {
	*ItemInfo
	*UserInfo
	CrossFeaStrMap map[string]string
	CrossFeaFltMap map[string]float32
}

type ItemScore struct {
	// 房间号
	ItemId int64
	// 模型阶段的得分
	ModelScore float64
	// 最终得分
	Score float64
	// 传递下去的特征
	Feature *Feature
	// 复合模型得分
	AssembleScore map[string]interface{}
}

type ItemScoreSlice []*ItemScore

func (s ItemScoreSlice) Len() int {
	return len(s)
}

func (s ItemScoreSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ItemScoreSlice) Less(i, j int) bool {
	return s[j].Score < s[i].Score
}
