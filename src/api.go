package src

type Request struct {
	Uid   int64
	Buvid string
	Scene string
}

type Response struct {
	List []*Item
}

type Item struct {
	Id      int64
	Score   float32
	Sources []string
	Feature *Feature
}
