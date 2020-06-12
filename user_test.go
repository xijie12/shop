package main

import (
	"testing"
	"fmt"
)

func TestUpLevel(t *testing.T) {
	UpLevel(10)
	if level != 2 || exp != 10{
		t.Fatalf("level(%d) != 2 or exp(%d) != 10", level, exp)
	}
	fmt.Printf("Level:%d, exp:%d",level,exp)
}
