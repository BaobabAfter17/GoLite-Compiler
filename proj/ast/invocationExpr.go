package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
)

type InvocationExpr struct {
	Func      *IdenLiteral
	Arguments []Expression
	Line int
}

func (inv *InvocationExpr) TokenLiteral() string { return "" }
func (inv *InvocationExpr) String() string {
	var out bytes.Buffer

	out.WriteString((*inv.Func).String())
	out.WriteString("(")
	for i, arg := range inv.Arguments {
		if i != 0 {
			out.WriteString(", ")
		}
		out.WriteString(arg.String())
	}
	out.WriteString(")")
	return out.String()
}

func (inv *InvocationExpr) TranslateToILOC(table *st.SymbolTable) *ir.ExpressionFrag {
	res := &ir.ExpressionFrag{}
	res.Reg = ir.NewRegister()

	pushIns := ir.NewPush()
	for _, exp := range inv.Arguments {
		expFrag := exp.TranslateToILOC(table)
		res.ExtendIns(expFrag.Body)
		pushIns.AppendReg(expFrag.Reg)
	}

	if len(pushIns.GetSources()) > 0 {
		res.AppendIns(pushIns)
	}
	
	res.AppendIns(ir.NewBl(inv.Func.String()))
	res.AppendIns(ir.NewMoveReg(res.Reg, 0))
	return res	
	
}