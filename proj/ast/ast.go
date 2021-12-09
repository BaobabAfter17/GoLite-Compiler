package ast

import (
	"bytes"
	ct "proj/token"
	tp "proj/type"
	st "proj/symTable"
	ir "proj/ir"
)

// The base Node interface that all ast nodes have to access
type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
	TranslateToILOC(*st.SymbolTable) *ir.ExpressionFrag
}

type Statement interface {
	Node
	TranslateToILOC(*st.SymbolTable) []ir.Instruction
}

type Package struct {
	Name string
	Line int
}

func (pkg *Package) TokenLiteral() string { return "package" }
func (pkg *Package) String() string {
	var out bytes.Buffer

	out.WriteString("package ")
	out.WriteString(pkg.Name)
	out.WriteString(";\n")
	return out.String()
}

type Import struct {
}

func (pkg *Import) TokenLiteral() string { return "import" }
func (pkg *Import) String() string {
	return `import "fmt";` + "\n"
}

type Field struct {
	Name string
	Typ  tp.Type
	Line int
}

func (fie *Field) TokenLiteral() string { return "field" }
func (fie *Field) String() string {
	var out bytes.Buffer

	out.WriteString(fie.Name)
	out.WriteString(" ")
	out.WriteString(fie.Typ.Literal())
	out.WriteString(";\n")
	return out.String()
}

type TypeDeclaration struct {
	Token  ct.Token
	Name   string
	Fields []*Field
	Line   int
}

func (typdcl *TypeDeclaration) TokenLiteral() string { return "type struct" }
func (typdcl *TypeDeclaration) String() string {
	var out bytes.Buffer

	out.WriteString("type ")
	out.WriteString(typdcl.Name)
	out.WriteString(" struct {\n")

	for _, fie := range typdcl.Fields {
		out.WriteString((*fie).String())
	}

	out.WriteString("}\n")
	return out.String()
}

type Types struct {
	TypDcls []*TypeDeclaration
}

func (tps *Types) TokenLiteral() string { return "Types" }
func (tps *Types) String() string {
	var out bytes.Buffer
	for _, tp := range tps.TypDcls {
		out.WriteString((*tp).String())
	}
	return out.String()
}

type Parameter struct {
	Name string
	Typ  tp.Type
	Line int
}

func (para *Parameter) TokenLiteral() string { return "Parameter" }
func (para *Parameter) String() string {
	return para.Name + " " + para.Typ.Literal()
}

type Operator int

const (
	ADD Operator = iota
	MULT
	SUB
	DIV
	AND
	OR
	EQ
	GE
	LE
	GT
	LT
	NE
	NOT
	ADDRESS
	ACCESS
)

func OpString(op Operator) string {

	switch op {
	case ADD:
		return "+"
	case MULT:
		return "*"
	case SUB:
		return "-"
	case DIV:
		return "/"
	case AND:
		return "&&"
	case OR:
		return "||"
	case EQ:
		return "=="
	case GE:
		return ">="
	case LE:
		return "<="
	case GT:
		return ">"
	case LT:
		return "<"
	case NE:
		return "!="
	case NOT:
		return "!"
	case ADDRESS:
		return "&"
	case ACCESS:
		return "."
	}
	panic("Error: Could not determine operator")
}