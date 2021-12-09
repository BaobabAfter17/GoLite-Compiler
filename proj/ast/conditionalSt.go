package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
)

type ConditionalSt struct {
	Condition Expression
	IfBlock   *Block
	ElseBlock *Block
	Line int
}

func (con *ConditionalSt) TokenLiteral() string { return "if" }
func (con *ConditionalSt) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString("(")
	out.WriteString((con.Condition).String())
	out.WriteString(") ")
	out.WriteString((*con.IfBlock).String())

	if con.ElseBlock != nil {
		out.WriteString("else\n")
		out.WriteString((*con.ElseBlock).String())
	}
	return out.String()
}

func (con *ConditionalSt) TranslateToILOC(table *st.SymbolTable) []ir.Instruction {
	var res []ir.Instruction

	condExprFrag :=  con.Condition.TranslateToILOC(table)
	res = append(res, condExprFrag.Body...)

	if con.ElseBlock == nil {
		res = append(res, ir.NewCmpC(condExprFrag.Reg, 1))
		newSkipIfLable := ir.NewLabelWithPre("skipIf")
		res = append(res, ir.NewBne(newSkipIfLable))
		res = append(res, con.IfBlock.TranslateToILOC(table)...)
		res = append(res, ir.NewLabelIns(newSkipIfLable))
	} else {
		res = append(res, ir.NewCmpC(condExprFrag.Reg, 1))
		newElseLable := ir.NewLabelWithPre("else")
		res = append(res, ir.NewBne(newElseLable))
		res = append(res, con.IfBlock.TranslateToILOC(table)...)
		newDoneLable := ir.NewLabelWithPre("ifDone")
		res = append(res, ir.NewBl(newDoneLable))
		res = append(res, ir.NewLabelIns(newElseLable))
		res = append(res, con.ElseBlock.TranslateToILOC(table)...)
		res = append(res, ir.NewLabelIns(newDoneLable))
	}

	return res

}