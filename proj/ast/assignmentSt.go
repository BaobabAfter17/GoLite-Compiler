package ast 

import (
	ir "proj/ir"
	st "proj/symTable"
)

type AssignmentSt struct {
	LVal Expression
	Expr Expression
	Line int
}

func (ass *AssignmentSt) TokenLiteral() string { return "=" }
func (ass *AssignmentSt) String() string {
	return (ass.LVal).String() + " = " + (ass.Expr).String() + ";\n"
}

func (ass *AssignmentSt) TranslateToILOC(table *st.SymbolTable) []ir.Instruction {
	var res []ir.Instruction
	var structAddr int
	var fieldName string

	rightFrag := ass.Expr.TranslateToILOC(table)
	res = append(res, rightFrag.Body...)
	assTarget := ass.LVal
	
	switch assTarget.(type) {
	case *BinOpExpr:
		structAddrFrag := assTarget.(*BinOpExpr).Left.TranslateToILOC(table)
		res = append(res, structAddrFrag.Body...)
		structAddr = structAddrFrag.Reg
		fieldName = assTarget.(*BinOpExpr).Right.String()
		res = append(res, ir.NewStoreRef(rightFrag.Reg, structAddr, fieldName))
	case *IdenLiteral:
		varFrag := assTarget.TranslateToILOC(table)
		res = append(res, ir.NewMoveReg(varFrag.Reg, rightFrag.Reg))
	}

	return res
}