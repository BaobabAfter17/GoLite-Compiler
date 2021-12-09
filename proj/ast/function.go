package ast

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
	tp "proj/type"
)

type Function struct {
	Name       string
	Parameters []*Parameter
	ReturnTyp  tp.Type
	Dcl        []*VarDeclaration
	Stmnts     []Statement
	Line       int
}

func (fu *Function) TokenLiteral() string { return "func" }
func (fu *Function) String() string {
	var out bytes.Buffer
	out.WriteString("func ")
	out.WriteString(fu.Name)
	out.WriteString("(")

	for i, para := range fu.Parameters {
		out.WriteString(para.String())
		if i != len(fu.Parameters)-1 {
			out.WriteString(", ")
		}
	}

	out.WriteString(") ")
	out.WriteString(fu.ReturnTyp.Literal())
	out.WriteString(" {\n")
	for _, dc := range fu.Dcl {
		out.WriteString((*dc).String())
	}
	for _, st := range fu.Stmnts {
		out.WriteString(st.String())
	}
	out.WriteString("}\n")
	return out.String()
}

func (fu *Function) TranslateToILOC(table *st.SymbolTable) *ir.FuncFrag {
	res := &ir.FuncFrag{}
	res.Label = fu.Name

	// change symbol table
	fentry := table.LookUpAll(fu.Name).(*st.FuncEntry)
	table = fentry.Scope

	// add label
	res.AppendIns(ir.NewLabelIns(res.Label))
	
	// push all parameters into the stack
	pushInsPara := ir.NewPush()
	for _, para := range fu.Parameters {
		// possible TODO: pop arguments from the stack
		entry := table.LookUpLocal(para.Name).(*st.VarEntry)
		entry.Reg = ir.NewRegister()
		pushInsPara.AppendReg(entry.Reg)
	}

	if len(pushInsPara.GetSources()) > 0 {
		res.AppendIns(pushInsPara)
	}

	// push all local variables
	pushInsVar := ir.NewPush()
	for _, dcl := range fu.Dcl {
		for _, vname := range dcl.Names {
			entry := table.LookUpLocal(vname).(*st.VarEntry)
			entry.Reg = ir.NewRegister()
			pushInsVar.AppendReg(entry.Reg)
		}
	}

	if len(pushInsVar.GetSources()) > 0 {
		res.AppendIns(pushInsVar)
	}

	// translate all statements
	for _, st := range fu.Stmnts {
		res.ExtendIns(st.TranslateToILOC(table))
	}

	// possible TODO: pop all local variables and parameters


	return res
}