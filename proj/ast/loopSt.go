package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
)

type LoopSt struct {
	Condition Expression
	LoopBlock *Block
	Line int
}

func (loo *LoopSt) TokenLiteral() string { return "for" }
func (loo *LoopSt) String() string {
	var out bytes.Buffer

	out.WriteString("for ")
	out.WriteString(loo.Condition.String())
	out.WriteString((*loo.LoopBlock).String())
	return out.String()
}

func (loo *LoopSt) TranslateToILOC(table *st.SymbolTable) []ir.Instruction {
	var res []ir.Instruction
	newLoopBodyLable := ir.NewLabelWithPre("loopBody")
	newTestCondLable := ir.NewLabelWithPre("testCond")

	res = append(res, ir.NewBl(newTestCondLable))
	res = append(res, ir.NewLabelIns(newLoopBodyLable))
	res = append(res, loo.LoopBlock.TranslateToILOC(table)...)
	res = append(res, ir.NewLabelIns(newTestCondLable))
	condExprFrag :=  loo.Condition.TranslateToILOC(table)
	res = append(res, condExprFrag.Body...)
	res = append(res, ir.NewCmpC(condExprFrag.Reg, 1))
	res = append(res, ir.NewBeq(newLoopBodyLable))
	return res
}