package main

import (
	"chaos-knight/src"
	"context"
)

func main() {
	conf, err := src.GetConfFromFile("conf.json")
	if err != nil {
		panic(err)
	}

	parser := &Interpreter{conf: conf}
	ctx := context.TODO()
	req := &src.Request{
		Uid:   925082,
		Buvid: "DJKFSJKSJAKFSF",
		Scene: "场景1",
	}
	_, _ = parser.Do(ctx, req, req.Scene)

}
