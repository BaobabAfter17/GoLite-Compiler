package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
	tp "proj/type"
)

type VarDeclaration struct {
	Names []string
	Typ   tp.Type
	Line  int
}

func (vardcl *VarDeclaration) TokenLiteral() string { return "var" }
func (vardcl *VarDeclaration) String() string {
	var out bytes.Buffer

	out.WriteString("var ")
	for i, name := range vardcl.Names {
		out.WriteString(name)
		if i != len(vardcl.Names)-1 {
			out.WriteString(", ")
		}
	}
	switch vardcl.Typ.(type) {
	case *tp.StructType:
		out.WriteString(" *" + vardcl.Typ.Literal())
	default:
		out.WriteString(" " + vardcl.Typ.Literal())
	}
	
	out.WriteString(";\n")
	return out.String()
}

func (vardcl *VarDeclaration) GlTranslateToILOC(table *st.SymbolTable) []*ir.GlobalVarFrag {
	var res []*ir.GlobalVarFrag
	typ := vardcl.Typ
	switch typ.(type) {
	case *tp.StructType:
		for _, name := range vardcl.Names {
			varFrag := &ir.GlobalVarFrag{Name: name}
			newReg := ir.NewRegister()
			varFrag.AppendIns(ir.NewAllocate(newReg, typ.(*tp.StructType).Name))
			varFrag.AppendIns(ir.NewLoadStoreGlobal("str", newReg, "@" + name))
			res = append(res, varFrag)
		}
	default:
		for _, name := range vardcl.Names {
			varFrag := &ir.GlobalVarFrag{Name: name}
			newReg := ir.NewRegister()
			varFrag.AppendIns(ir.NewMove(newReg, 0))
			varFrag.AppendIns(ir.NewLoadStoreGlobal("str", newReg, "@" + name))
			res = append(res, varFrag)
		}
	}
	return res
}