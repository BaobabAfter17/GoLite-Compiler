package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
	ct "proj/token"
)

type UnOpExpr struct {
	Token    ct.Token // The token from the scanner
	Operator Operator // The unary operator
	Left     Expression
	Line int
}

func (unOp *UnOpExpr) TokenLiteral() string {
	return unOp.Token.Literal
}

func (unOp *UnOpExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(OpString(unOp.Operator) + " ")
	out.WriteString(unOp.Left.String())
	out.WriteString(")")

	return out.String()
}

func (unOp *UnOpExpr) TranslateToILOC(table *st.SymbolTable) *ir.ExpressionFrag {
	res := &ir.ExpressionFrag{}
	res.Reg = ir.NewRegister()
	leftFrag := unOp.Left.TranslateToILOC(table)
	switch unOp.Left.(type) {
	case *IntLiteral:
		break
	default:
		res.ExtendIns(leftFrag.Body)
	}

	switch unOp.Operator {
	case NOT:
		switch unOp.Left.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewNotC(res.Reg, unOp.Left.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewNot(res.Reg, leftFrag.Reg))
		}
	case SUB:
		zeroReg := ir.NewRegister()
		res.AppendIns(ir.NewMove(zeroReg, 0))

		switch unOp.Left.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewSubC(res.Reg, zeroReg, unOp.Left.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewSub(res.Reg, zeroReg, leftFrag.Reg))
		}
	}
	return res
}