package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
)

type ReadSt struct {
	Iden string
	Line int
}

func (rea *ReadSt) TokenLiteral() string { return "scan" }
func (rea *ReadSt) String() string {
	var out bytes.Buffer

	out.WriteString("fmt.Scan(&")
	out.WriteString(rea.Iden)
	out.WriteString(");\n")
	return out.String()
}

func (rea *ReadSt) TranslateToILOC(table *st.SymbolTable) []ir.Instruction {
	var res []ir.Instruction
	entry := table.LookUpAll(rea.Iden)
	
	res = append(res, ir.NewRead(entry.(*st.VarEntry).Reg))
	return res
}