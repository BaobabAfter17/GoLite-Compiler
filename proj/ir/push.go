package ir

import (
	"bytes"
	"strconv"
)

type Push struct{
	regs []int 
}

func NewPush() *Push {
	return &Push{}
}

func (instr *Push) GetTargets() []int {
	return nil
}

func (instr *Push) AppendReg(reg int) {
	instr.regs = append(instr.regs, reg)
}


func (instr *Push) GetSources() []int {
	return instr.regs
}

func (instr *Push) GetImmediate() *int {
	return nil
}
func (instr *Push) GetLabel() string {
	return ""
}
func (instr *Push) SetLabel(newLabel string) {
}

func (instr *Push) String() string {
	var out bytes.Buffer
	out.WriteString("push {")
	for i, v := range instr.regs {
		if i == 0 {
			out.WriteString("r" + strconv.Itoa(v))
		} else {
			out.WriteString(",r" + strconv.Itoa(v))
		}
	}
	out.WriteString("}")
	return out.String()
}