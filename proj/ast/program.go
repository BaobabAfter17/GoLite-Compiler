package ast

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
)

type Program struct {
	Pkg   *Package
	Impt  *Import
	Typs  *Types
	Dcl   []*VarDeclaration
	Funcs []*Function
}

func (pg *Program) TokenLiteral() string { return "program" }
func (pg *Program) String() string {
	var out bytes.Buffer
	out.WriteString((*pg.Pkg).String())
	out.WriteString((*pg.Impt).String())
	out.WriteString((*pg.Typs).String())
	for _, dc := range pg.Dcl {
		out.WriteString((*dc).String())
	}
	for _, fu := range pg.Funcs {
		out.WriteString((*fu).String())
	}
	return out.String()
}

func (p *Program) TranslateToILOC(table *st.SymbolTable) *ir.ProgramFrag {
	res := &ir.ProgramFrag{}
	
	for _, gdcl := range p.Dcl {
		res.Dcls = append(res.Dcls, gdcl.GlTranslateToILOC(table)...)
	}

	for _, fun := range p.Funcs {
		res.Funcs = append(res.Funcs, fun.TranslateToILOC(table))
	}

	return res
}
