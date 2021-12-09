package ir

import (
	"fmt"
)

type LoadStoreGlobal struct{
	target    int        
	source    string      
	method    string   
}

//NewAdd is a constructor and initialization function for a new Add instruction
func NewLoadStoreGlobal(method string, target int, source string) *LoadStoreGlobal {
	return &LoadStoreGlobal{target, source, method}
}

func (instr *LoadStoreGlobal) GetTargets() []int {
	targets := make([]int, 1)
	targets = append(targets, instr.target)
	return targets
}


func (instr *LoadStoreGlobal) GetSources() []int {
	return nil
}

func (instr *LoadStoreGlobal) GetSourceName() string {
	return instr.source
}

func (instr *LoadStoreGlobal) GetImmediate() *int {
	return nil
}
func (instr *LoadStoreGlobal) GetLabel() string {
	return ""
}
func (instr *LoadStoreGlobal) SetLabel(newLabel string) {
}

func (instr *LoadStoreGlobal) String() string {
	return fmt.Sprintf("%s %v,%s", instr.method, instr.target, instr.source)
}