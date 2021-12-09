package ir

import (
	"fmt"
)

// Iloc instruction with two registers and one constant
type CommandRRC struct{
	target    int        // The target register for the instruction
	reg2      int
	constant  int  
	method    string
}

func NewCommandRRC(method string, target int, reg2 int, constant int) *CommandRRC {
	return &CommandRRC{target, reg2, constant, method}
}

func (instr *CommandRRC) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *CommandRRC) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.reg2)
	return sources
}
func (instr *CommandRRC) GetImmediate() *int {
	return &instr.constant
}

func (instr *CommandRRC) GetLabel() string {
	return ""
}

func (instr *CommandRRC) SetLabel(newLabel string){
}

func (instr *CommandRRC) String() string {
	return fmt.Sprintf("%s r%v,r%v,#%v", instr.method, instr.target, instr.reg2, instr.constant)
}