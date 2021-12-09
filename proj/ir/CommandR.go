package ir

import (
	"fmt"
)

// Iloc instruction with only one register target
type CommandR struct{
	target    int        // The target register for the instruction
	method    string
}

func NewCommandR(method string, target int) *CommandR {
	return &CommandR{target, method}
}

func (instr *CommandR) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *CommandR) GetSources() []int {
	return nil
}
func (instr *CommandR) GetImmediate() *int {
	return nil
}

func (instr *CommandR) GetLabel() string {
	return ""
}

func (instr *CommandR) SetLabel(newLabel string){
}

func (instr *CommandR) String() string {
	return fmt.Sprintf("%s r%v", instr.method, instr.target)
}