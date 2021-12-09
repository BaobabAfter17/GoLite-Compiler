package ast 

import (
	ir "proj/ir"
	st "proj/symTable"
)

type PrintSt struct {
	MethodName string
	Iden       string
	Line int
}

func (pri *PrintSt) TokenLiteral() string { return pri.MethodName }
func (pri *PrintSt) String() string {
	return "fmt." + pri.MethodName + " (" + pri.Iden + ");\n"
}

func (pri *PrintSt) TranslateToILOC(table *st.SymbolTable) []ir.Instruction {
	var res []ir.Instruction
	entry := table.LookUpAll(pri.Iden)
	if pri.MethodName == "print" {
		res = append(res, ir.NewPrint(entry.(*st.VarEntry).Reg))
	} else {
		res = append(res, ir.NewPrintln(entry.(*st.VarEntry).Reg))
	}
	return res
}