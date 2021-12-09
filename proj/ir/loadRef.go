package ir

import (
	"fmt"
)

type LoadRef struct{
	target    int      
	reg2      int  
	field    string      
}

//NewAdd is a constructor and initialization function for a new Add instruction
func NewLoadRef(target int, reg2 int, field string) *LoadRef {
	return &LoadRef{target, reg2, field}
}

func (instr *LoadRef) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}


func (instr *LoadRef) GetSources() []int {
	return []int{ instr.reg2 }
}

func (instr *LoadRef) GetFieldName() string {
	return instr.field
}

func (instr *LoadRef) GetImmediate() *int {
	return nil
}
func (instr *LoadRef) GetLabel() string {
	return ""
}
func (instr *LoadRef) SetLabel(newLabel string) {
}

func (instr *LoadRef) String() string {
	return fmt.Sprintf("loadRef r%v,r%v,@%s", instr.target, instr.reg2, instr.field)
}