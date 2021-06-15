package src

type Request struct {
	uid   int64
	buvid string
}

type Response struct {
	List []*Item
}

type Item struct {
	id      int64
	score   float32
	sources []string
	feature *Feature
}
