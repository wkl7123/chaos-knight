package main

import (
	"chaos-knight/src"
	"context"
	"fmt"
)

func main() {
	conf, err := src.GetConfFromFile("conf.json")
	if err != nil {
		panic(err)
	}

	parser := &Interpreter{conf: conf}
	for i := 0; i < 1000000; i++ {
		ctx := context.TODO()
		req := &src.Request{
			Uid:   925082,
			Buvid: "DJKFSJKSJAKFSF",
			Scene: "场景1",
		}
		_, trackInfo, err := parser.Do(ctx, req, req.Scene)
		// form msg to message queue, and resp to client
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		fmt.Printf("%s,%s,%s,%s,%s\n", trackInfo["recall"], trackInfo["feature"], trackInfo["filter"], trackInfo["model"], trackInfo["rerank"])
	}
}
