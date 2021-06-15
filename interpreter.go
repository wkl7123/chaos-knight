package main

import (
	"chaos-knight/src"
	"context"
)

type Interpreter struct {
	conf *src.Conf
}

func (it *Interpreter) Do(ctx context.Context, req *src.Request, scene string) {

}
