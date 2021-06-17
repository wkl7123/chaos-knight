package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Conf map[string]Scene

type Scene struct {
	Layers FullLayers `json:"layers"`
}

type FullLayers struct {
	Recall  []LayerOption `json:"recall"`
	Feature []LayerOption `json:"feature"`
	Filter  []LayerOption `json:"filter"`
	Model   []LayerOption `json:"model"`
	Rerank  []LayerOption `json:"rerank"`
}

type LayerOption struct {
	Prob   float32  `json:"prob"`
	Choose []string `json:"choose"`
}

func GetConfFromFile(confPath string) (conf *Conf, err error) {
	conf = &Conf{}
	confData, err := ioutil.ReadFile("conf.json")
	if err != nil {
		err = fmt.Errorf("config is not valid json: err: %v", err)
		return
	}
	err = json.Unmarshal(confData, conf)
	if err != nil {
		err = fmt.Errorf("conf not valid: err: %v", err)
		return
	}
	return
}
