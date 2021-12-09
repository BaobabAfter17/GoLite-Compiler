package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
)

type InvocationSt struct {
	FuncName  string
	Arguments []Expression
	Line int
}

func (inv *InvocationSt) TokenLiteral() string { return "" }
func (inv *InvocationSt) String() string {
	var out bytes.Buffer

	out.WriteString(inv.FuncName)
	out.WriteString("(")
	for i, arg := range inv.Arguments {
		if i != 0 {
			out.WriteString(", ")
		}
		out.WriteString(arg.String())
	}
	out.WriteString(");\n")
	return out.String()
}

func (inv *InvocationSt) TranslateToILOC(table *st.SymbolTable) []ir.Instruction {
	var res []ir.Instruction

	pushIns := ir.NewPush()
	for _, exp := range inv.Arguments {
		expFrag := exp.TranslateToILOC(table)
		pushIns.AppendReg(expFrag.Reg)
	}

	if len(pushIns.GetSources()) > 0 {
		res = append(res, pushIns)
	}

	res = append(res, ir.NewBl(inv.FuncName))
	return res	
}
