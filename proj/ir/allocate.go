package ir

import (
	"fmt"
)

type Allocate struct{
	target    int        // The target register for the instruction
	specifiedTyp string
}

func NewAllocate(target int, specifiedTyp string) *Allocate {
	return &Allocate{target, specifiedTyp}
}

func (instr *Allocate) GetTargets() []int {
	targets := make([]int, 1)
	targets = append(targets, instr.target)
	return targets
}

func (instr *Allocate) GetSources() []int {
	return nil
}

func (instr *Allocate) GetSpecifiedType() string {
	return instr.specifiedTyp
}

func (instr *Allocate) GetImmediate() *int {
	return nil
}

func (instr *Allocate) GetLabel() string {
	return ""
}

func (instr *Allocate) SetLabel(newLabel string){
}

func (instr *Allocate) String() string {
	return fmt.Sprintf("new r%v, %s",instr.target, instr.specifiedTyp)
}