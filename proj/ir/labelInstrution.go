package ir

import (
	"fmt"
)

type LabelIns struct{
	label string
}

func NewLabelIns(label string) *LabelIns {
	return &LabelIns{label}
}

func (instr *LabelIns) GetTargets() []int {
	return nil
}

func (instr *LabelIns) GetSources() []int {
	return nil
}

func (instr *LabelIns) GetImmediate() *int {
	return nil
}

func (instr *LabelIns) GetLabel() string {
	return instr.label
}

func (instr *LabelIns) SetLabel(newLabel string){
}

func (instr *LabelIns) String() string {
	return fmt.Sprintf("%s:",instr.label)
}