package ast 

import (
	ir "proj/ir"
	st "proj/symTable"
	ct "proj/token"
)

// ======================================================================
//  Single Indentifier
// ======================================================================
type IdenLiteral struct {
	Token ct.Token
	Value string
}

func (iden *IdenLiteral) TokenLiteral() string { return iden.Value }
func (iden *IdenLiteral) String() string       { return iden.Value }
func (iden *IdenLiteral) TranslateToILOC(table *st.SymbolTable) *ir.ExpressionFrag {
	res := &ir.ExpressionFrag{}
	entry := table.LookUpAll(iden.String())
	if entry == nil { 
		res.Reg = -1 
	} else {
		res.Reg = entry.(*st.VarEntry).Reg
	}

	return res
}

// ======================================================================
//  Integer Constant
// ======================================================================
type IntLiteral struct {
	Token ct.Token
	Value int
}

func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntLiteral) String() string       { return il.Token.Literal }
func (il *IntLiteral) TranslateToILOC(table *st.SymbolTable) *ir.ExpressionFrag {
	res := &ir.ExpressionFrag{}
	res.Reg = ir.NewRegister()
	res.AppendIns(ir.NewMove(res.Reg, il.Value))
	return res
}

// ======================================================================
//  Boolean Constant
// ======================================================================
type BoolLiteral struct {
	Token ct.Token
	Value bool
}

func (bl *BoolLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BoolLiteral) String() string       { return bl.Token.Literal }
func (bl *BoolLiteral) TranslateToILOC(table *st.SymbolTable) *ir.ExpressionFrag {
	res := &ir.ExpressionFrag{}
	res.Reg = ir.NewRegister()
	if bl.Value == false {
		res.AppendIns(ir.NewMove(res.Reg, 0))
	} else {
		res.AppendIns(ir.NewMove(res.Reg, 1))
	}
	
	return res
}


// ======================================================================
//  NIL
// ======================================================================
type NilLiteral struct {
	Token ct.Token
}

func (nl *NilLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NilLiteral) String() string       { return nl.Token.Literal }
func (nl *NilLiteral) TranslateToILOC(table *st.SymbolTable) *ir.ExpressionFrag {
	res := &ir.ExpressionFrag{}
	res.Reg = ir.NewRegister()
	res.AppendIns(ir.NewMove(res.Reg, 0))
	return res
}
