package main

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	x := math.Abs(-1)
	if x != 1.0 {
		t.Errorf("%f != -1", x)
	}
}
