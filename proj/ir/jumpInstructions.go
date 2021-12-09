package ir

import (
	"fmt"
)

type JumpCommand struct {
	label string
	method string
}

func NewJumpCommand(method string, label string) *JumpCommand {
	return &JumpCommand{label, method}
}

func (instr *JumpCommand) GetTargets() []int {
	return nil
}

func (instr *JumpCommand) GetSources() []int {
	return nil
}

func (instr *JumpCommand) GetImmediate() *int {
	return nil
}

func (instr *JumpCommand) GetLabel() string {
	return instr.label
}

func (instr *JumpCommand) SetLabel(newLabel string) {
}

func (instr *JumpCommand) String() string {
	return fmt.Sprintf("%s %s", instr.method, instr.label)
}

type Beq struct {
	JumpCommand
}

func NewBeq(label string) *Beq {
	res := &Beq{}
	res.JumpCommand = *NewJumpCommand("beq", label)
	return res
}

type Bne struct {
	JumpCommand
}

func NewBne(label string) *Bne {
	res := &Bne{}
	res.JumpCommand = *NewJumpCommand("bne", label)
	return res
}

type Bl struct {
	JumpCommand
}

func NewBl(label string) *Bl {
	res := &Bl{}
	res.JumpCommand = *NewJumpCommand("b", label)
	return res
}