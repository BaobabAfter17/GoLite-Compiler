package ir

import (
	"fmt"
)

type StoreRef struct{
	target    int      
	reg2      int  
	field    string      
}

//NewAdd is a constructor and initialization function for a new Add instruction
func NewStoreRef(target int, reg2 int, field string) *StoreRef {
	return &StoreRef{target, reg2, field}
}

func (instr *StoreRef) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}


func (instr *StoreRef) GetSources() []int {
	return []int{ instr.reg2 }
}

func (instr *StoreRef) GetFieldName() string {
	return instr.field
}

func (instr *StoreRef) GetImmediate() *int {
	return nil
}
func (instr *StoreRef) GetLabel() string {
	return ""
}
func (instr *StoreRef) SetLabel(newLabel string) {
}

func (instr *StoreRef) String() string {
	return fmt.Sprintf("strRef r%v,r%v,@%s", instr.target, instr.reg2, instr.field)
}