package ir

import (
	"fmt"
)

// Iloc instruction with only one register target and one constant
type CommandRC struct{
	target    int        // The target register for the instruction
	constant  int
	method    string
}

func NewCommandRC(method string, target int, reg2 int) *CommandRC {
	return &CommandRC{target, reg2, method}
}

func (instr *CommandRC) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *CommandRC) GetSources() []int {
	return nil
}
func (instr *CommandRC) GetImmediate() *int {
	return &instr.constant
}

func (instr *CommandRC) GetLabel() string {
	return ""
}

func (instr *CommandRC) SetLabel(newLabel string){
}

func (instr *CommandRC) String() string {
	return fmt.Sprintf("%s r%v,#%v", instr.method, instr.target, instr.constant)
}