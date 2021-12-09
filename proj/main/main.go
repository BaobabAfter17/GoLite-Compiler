package main

import (
	"proj/ast"
	"proj/semantic"
)

func main() {
	root := &ast.Program{}

	// adding root.Typs
	fields := make([]*ast.Field, 1)
	field := &ast.Field{
		Name: "x",
		Typ:  ast.IntType,
	}
	fields[0] = field

	tds := make([]*ast.TypeDeclaration, 1)

	td := &ast.TypeDeclaration{
		Name:   "animal",
		Fields: fields,
	}

	tds[0] = td
	tps := &ast.Types{TypDcls: tds}
	root.Typs = tps

	// root.Dcl
	dcls := make([]*ast.VarDeclaration, 1)

	names := make([]string, 3)
	names[0] = "x"
	names[1] = "y"
	names[2] = "z"
	dcl := &ast.VarDeclaration{
		Names: names,
		Typ:   ast.IntType,
	}

	dcls[0] = dcl
	root.Dcl = dcls

	// root.Funcs
	dcls1 := make([]*ast.VarDeclaration, 1)

	names1 := make([]string, 3)
	names1[0] = "someBool"
	names1[1] = "otherBool"
	names1[2] = "againBool"
	dcl1 := &ast.VarDeclaration{
		Names: names1,
		Typ:   ast.BoolType,
	}
	dcls1[0] = dcl1

	mainf := &ast.Function{
		Name: "main",
	}

	f1 := &ast.Function{
		Name:       "foo",
		Parameters: dcls,
		Dcl:        dcls1,
	}

	var fns []*ast.Function
	fns = append(fns, mainf)
	fns = append(fns, f1)
	root.Funcs = fns

	sa := semantic.New(root)
	sa.Analyse()
	// fmt.Println(sa.String())
}
