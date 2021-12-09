package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
)

type ReturnSt struct {
	Expr Expression
	Line int
}

func (ret *ReturnSt) TokenLiteral() string { return "return" }
func (ret *ReturnSt) String() string {
	var out bytes.Buffer

	out.WriteString("return ")
	out.WriteString(ret.Expr.String())
	out.WriteString(";\n")
	return out.String()
}

func (ret *ReturnSt) TranslateToILOC(table *st.SymbolTable) []ir.Instruction {
	var res []ir.Instruction
	
	retExprFrag := ret.Expr.TranslateToILOC(table)
	res = append(res, retExprFrag.Body...)
	res = append(res, ir.NewReturn(retExprFrag.Reg))
	return res
}