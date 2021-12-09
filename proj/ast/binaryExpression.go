package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
	ct "proj/token"
)

type BinOpExpr struct {
	Token    ct.Token // The token from the scanner
	Operator Operator // The operator for the binary expression
	Left     Expression
	Right    Expression
	Line int
}

func (binOp *BinOpExpr) TokenLiteral() string {
	return binOp.Token.Literal
}

func (binOp *BinOpExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(binOp.Left.String())
	out.WriteString(" " + OpString(binOp.Operator) + " ")
	out.WriteString(binOp.Right.String())
	out.WriteString(")")

	return out.String()
}

func (binOp *BinOpExpr) TranslateToILOC(table *st.SymbolTable) *ir.ExpressionFrag {
	res := &ir.ExpressionFrag{}
	res.Reg = ir.NewRegister()

	leftFrag := binOp.Left.TranslateToILOC(table)
	res.ExtendIns(leftFrag.Body)
	rightFrag := binOp.Right.TranslateToILOC(table)
	switch binOp.Right.(type) {
	case *IntLiteral:
		break
	default:
		res.ExtendIns(rightFrag.Body)
	}
	
	switch binOp.Operator {
	case ADD:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewAddC(res.Reg, leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewAdd(res.Reg, leftFrag.Reg, rightFrag.Reg))
		}
	case SUB:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewSubC(res.Reg, leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewSub(res.Reg, leftFrag.Reg, rightFrag.Reg))
		}
	case MULT:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewMulC(res.Reg, leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewMul(res.Reg, leftFrag.Reg, rightFrag.Reg))
		}
	case DIV:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewDivC(res.Reg, leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewDiv(res.Reg, leftFrag.Reg, rightFrag.Reg))
		}
	case AND:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewAndC(res.Reg, leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewAnd(res.Reg, leftFrag.Reg, rightFrag.Reg))
		}
	case OR:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewOrC(res.Reg, leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewOr(res.Reg, leftFrag.Reg, rightFrag.Reg))
		}
	case EQ:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewCmpC(leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewCmp(leftFrag.Reg, rightFrag.Reg))
		}
		res.AppendIns(ir.NewMoveEq(res.Reg, 1))
		res.AppendIns(ir.NewMoveNe(res.Reg, 0))
	case NE:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewCmpC(leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewCmp(leftFrag.Reg, rightFrag.Reg))
		}
		res.AppendIns(ir.NewMoveEq(res.Reg, 0))
		res.AppendIns(ir.NewMoveNe(res.Reg, 1))
	case GE:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewCmpC(leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewCmp(leftFrag.Reg, rightFrag.Reg))
		}
		res.AppendIns(ir.NewMoveGe(res.Reg, 1))
		res.AppendIns(ir.NewMoveLt(res.Reg, 0))
	case LE:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewCmpC(leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewCmp(leftFrag.Reg, rightFrag.Reg))
		}
		res.AppendIns(ir.NewMoveLe(res.Reg, 1))
		res.AppendIns(ir.NewMoveGt(res.Reg, 0))
	case GT:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewCmpC(leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewCmp(leftFrag.Reg, rightFrag.Reg))
		}
		res.AppendIns(ir.NewMoveGt(res.Reg, 1))
		res.AppendIns(ir.NewMoveLe(res.Reg, 0))
	case LT:
		switch binOp.Right.(type) {
		case *IntLiteral:
			res.AppendIns(ir.NewCmpC(leftFrag.Reg, binOp.Right.(*IntLiteral).Value))
		default:
			res.AppendIns(ir.NewCmp(leftFrag.Reg, rightFrag.Reg))
		}
		res.AppendIns(ir.NewMoveLt(res.Reg, 1))
		res.AppendIns(ir.NewMoveGe(res.Reg, 0))
	case ACCESS:
		res.AppendIns(ir.NewLoadRef(res.Reg, leftFrag.Reg, binOp.Right.String()))
	}
	return res
}