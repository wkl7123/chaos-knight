package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Conf map[string]*Scene

type Scene struct {
	Layers FullLayers `json:"layers"`
}

type FullLayers struct {
	Recall  []*LayerOption `json:"recall"`
	Feature []*LayerOption `json:"feature"`
	Filter  []*LayerOption `json:"filter"`
	Model   []*LayerOption `json:"model"`
	Rerank  []*LayerOption `json:"rerank"`
}

type LayerOption struct {
	Prob   float32  `json:"prob"`
	Choice []string `json:"choice"`
}

func (inst LayerOption) String() string {
	return strings.Join(inst.Choice, ";")
}

func CheckOptions(options []*LayerOption) bool {
	var sum float32 = 0.0
	for _, option := range options {
		sum += option.Prob
	}
	return math.Abs(float64(sum)-1.0) < 0.000001
}

func Pick(options []*LayerOption) *LayerOption {
	flt := rand.Float32()
	var acc float32 = 0.0
	for _, option := range options {
		if flt < (acc + option.Prob) {
			return option
		}
		acc += option.Prob
	}
	return options[len(options)-1]
}

func GetConfFromFile(confPath string) (conf Conf, err error) {
	conf = Conf{}
	confData, err := ioutil.ReadFile(confPath)
	if err != nil {
		err = fmt.Errorf("config is not valid json: err: %v", err)
		return
	}
	err = json.Unmarshal(confData, &conf)
	if err != nil {
		err = fmt.Errorf("conf not valid: err: %v", err)
		return
	}
	return
}
