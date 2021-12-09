package ast 

import (
	"bytes"
	ir "proj/ir"
	st "proj/symTable"
)

type Block struct {
	Stmnts []Statement
}

func (blk *Block) TokenLiteral() string { return "{}" }
func (blk *Block) String() string {
	var out bytes.Buffer

	out.WriteString("{\n")
	for _, st := range blk.Stmnts {
		out.WriteString("    " + st.String())
	}
	out.WriteString("}\n")
	return out.String()
}

func (blk *Block) TranslateToILOC(table *st.SymbolTable) []ir.Instruction {
	var res []ir.Instruction

	if blk.Stmnts == nil { return res }
	for _, st := range blk.Stmnts {
		switch st.(type) {
		case *Block:
			res = append(res, st.(*Block).TranslateToILOC(table)...)
		case *AssignmentSt:
			res = append(res, st.(*AssignmentSt).TranslateToILOC(table)...)
		case *PrintSt:
			res = append(res, st.(*PrintSt).TranslateToILOC(table)...)
		case *ConditionalSt:
			res = append(res, st.(*ConditionalSt).TranslateToILOC(table)...)
		case *LoopSt:
			res = append(res, st.(*LoopSt).TranslateToILOC(table)...)
		case *ReturnSt:
			res = append(res, st.(*ReturnSt).TranslateToILOC(table)...)
		case *ReadSt:
			res = append(res, st.(*ReadSt).TranslateToILOC(table)...)
		case *InvocationSt:
			res = append(res, st.(*InvocationSt).TranslateToILOC(table)...)
		}
	}

	return res
	
}