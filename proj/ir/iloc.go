package ir

import (
	// "proj/codegen"
)

type OperandTy int

const (
	REGISTER OperandTy = iota
	IMMEDIATE
)

type Instruction interface {

	GetTargets() []int 	// Get the registers targeted by this instruction

	GetSources() []int 	// Get the source registers for this instruction

	GetImmediate() *int  // Get the immediate value (i.e., constant) of this instruction

	GetLabel() string    // Get the label for this instruction

	SetLabel(newLabel string)  //Set the label for this instruction

	String() string  // Return a string representation of this instruction

}

type Frag struct {
	Body    []Instruction    // Function body of ILOC instructions
}

func (fr *Frag) AppendIns(ins Instruction) {
	fr.Body = append(fr.Body, ins)
}

func (fr *Frag) ExtendIns(arrIns []Instruction) {
	fr.Body = append(fr.Body, arrIns...)
}


type FuncFrag struct {
	Frag
	Label 	string           // Function name
	// Body    []Instruction    // Function body of ILOC instructions
	// Frame   *codegen.Frame	 // Activation Records (i.e., stack frame) for this function
}

type GlobalVarFrag struct {
	Frag
	Name string
}

type ProgramFrag struct {
	Dcls  []*GlobalVarFrag
	Funcs []*FuncFrag
}

type ExpressionFrag struct {
	Frag
	Reg int
}