package ir

import (
	"fmt"
)

// Iloc instruction with three registers
type CommandRRR struct{
	target    int        // The target register for the instruction
	reg2      int
	reg3      int  
	method    string
}

func NewCommandRRR(method string, target int, reg2 int, reg3 int) *CommandRRR {
	return &CommandRRR{target, reg2, reg3, method}
}

func (instr *CommandRRR) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *CommandRRR) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.reg2)
	sources = append(sources, instr.reg3)
	return sources
}
func (instr *CommandRRR) GetImmediate() *int {
	return nil
}

func (instr *CommandRRR) GetLabel() string {
	return ""
}

func (instr *CommandRRR) SetLabel(newLabel string){
}

func (instr *CommandRRR) String() string {
	return fmt.Sprintf("%s r%v,r%v,r%v", instr.method, instr.target, instr.reg2, instr.reg3)
}