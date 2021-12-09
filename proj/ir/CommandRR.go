package ir

import (
	"fmt"
)

// Iloc instruction with exactly two register
type CommandRR struct{
	target    int        // The target register for the instruction
	reg2      int
	method    string
}

func NewCommandRR(method string, target int, reg2 int) *CommandRR {
	return &CommandRR{target, reg2, method}
}

func (instr *CommandRR) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *CommandRR) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.reg2)
	return sources
}
func (instr *CommandRR) GetImmediate() *int {
	return nil
}

func (instr *CommandRR) GetLabel() string {
	return ""
}

func (instr *CommandRR) SetLabel(newLabel string){
}

func (instr *CommandRR) String() string {
	return fmt.Sprintf("%s r%v,r%v", instr.method, instr.target, instr.reg2)
}