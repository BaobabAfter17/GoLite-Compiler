package ir

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	addIns := NewAdd(1, 2, 3)
	
	fmt.Printf("%v\n", addIns.GetSources())
	fmt.Printf("%s\n", addIns.String())
}