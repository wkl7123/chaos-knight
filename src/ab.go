package src

type Conf map[string]Scene

type Scene struct {
	Layers FullLayers `json:"layers"`
}

type FullLayers struct {
	Recall  RecallLayer  `json:"recall"`
	Feature FeatureLayer `json:"feature"`
	Filter  FilterLayer  `json:"filter"`
	Model   ModelLayer   `json:"model"`
	Rerank  RerankLayer  `json:"rerank"`
}

type RecallLayer struct {
	Options []Option `json:"options"`
}

type FeatureLayer struct {
	Options []Option `json:"options"`
}

type FilterLayer struct {
	Options []Option `json:"options"`
}

type ModelLayer struct {
	Options []Option `json:"options"`
}

type RerankLayer struct {
	Options []Option `json:"options"`
}

type Option struct {
	Prob   float32  `json:"prob"`
	Choose []string `json:"choose"`
}
