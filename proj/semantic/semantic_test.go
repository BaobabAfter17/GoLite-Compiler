package semantic

import (
	"fmt"
	"os"
	"proj/ast"
	"testing"
	"io/ioutil"
	tp "proj/type"
	ct "proj/token"
	"proj/scanner"
	"proj/parser"
)

func ConstructAST(filePath string) *ast.Program {
	file, err := os.Open(filePath)
	var tokens []ct.Token
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	scanner := scanner.New(string(content))
	for true {
		tok := scanner.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == ct.EOF { break }
	}

	parser := parser.New(tokens, filePath)
	ast, _ := parser.Parse()
	if ast != nil {
		fmt.Print(ast.String())
	}
	return ast
}

func Test1(t *testing.T) {

	// Build up an Program AST
	root := &ast.Program{}

	// adding root.Typs
	fields := make([]*ast.Field, 1)
	field := &ast.Field{
		Name: "x",
		Typ:  tp.IntType,
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
		Typ:   tp.IntType,
	}

	dcls[0] = dcl
	root.Dcl = dcls

	paras := make([]*ast.Parameter, 1)
	para := &ast.Parameter{
		Name: "x",
		Typ:  tp.IntType,
		Line: 10,
	}
	paras[0] = para

	// root.Funcs
	dcls1 := make([]*ast.VarDeclaration, 1)

	names1 := make([]string, 3)
	names1[0] = "someBool"
	names1[1] = "otherBool"
	names1[2] = "againBool"
	dcl1 := &ast.VarDeclaration{
		Names: names1,
		Typ:   tp.BoolType,
	}
	dcls1[0] = dcl1

	mainf := &ast.Function{
		Name: "main",
	}

	f1 := &ast.Function{
		Name:       "foo",
		Parameters: paras,
		Dcl:        dcls1,
	}

	var fns []*ast.Function
	fns = append(fns, mainf)
	fns = append(fns, f1)
	root.Funcs = fns

	sa := New(root, "fake.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test2(t *testing.T) {

	// Build up an Program AST
	root := &ast.Program{}

	// root.Dcl
	dcls := make([]*ast.VarDeclaration, 2)

	names := make([]string, 3)
	names[0] = "x"
	names[1] = "y"
	names[2] = "z"
	dcl := &ast.VarDeclaration{
		Names: names,
		Typ:   tp.IntType,
	}

	dcl1 := &ast.VarDeclaration{
		Names: []string{"x"},
		Typ:   tp.BoolType,
	}

	dcls[0] = dcl
	dcls[1] = dcl1
	root.Dcl = dcls

	// root.Funcs
	mainf := &ast.Function{
		Name: "main",
	}

	var fns []*ast.Function
	fns = append(fns, mainf)
	root.Funcs = fns

	sa := New(root, "fake.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test3(t *testing.T) {

	// Build up an Program AST
	root := &ast.Program{}

	// root.Dcl
	dcls := make([]*ast.VarDeclaration, 1)

	names := make([]string, 3)
	names[0] = "x"
	names[1] = "y"
	names[2] = "z"
	dcl := &ast.VarDeclaration{
		Names: names,
		Typ:   tp.IntType,
	}

	dcls[0] = dcl
	root.Dcl = dcls

	// root.Funcs
	paras := make([]*ast.Parameter, 1)
	para := &ast.Parameter{
		Name: "x",
		Typ:  tp.IntType,
		Line: 10,
	}
	paras[0] = para
	mainf := &ast.Function{
		Name: "main",
		Parameters: paras,
	}

	var fns []*ast.Function
	fns = append(fns, mainf)
	root.Funcs = fns

	sa := New(root, "fake.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test4(t *testing.T) { 
	root := ConstructAST("test_files/test1.golite")
	sa := New(root, "test_files/test1.golite")
	if ok := sa.Analyse(); !ok {
		t.Errorf("Expected Sementic Analysis to pass!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test5(t *testing.T) { 
	// field type not seen so far
	root := ConstructAST("test_files/test2.golite")
	sa := New(root, "test_files/test2.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test6(t *testing.T) { 
	// package main
	root := ConstructAST("test_files/test3.golite")
	sa := New(root, "test_files/test3.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test7(t *testing.T) { 
	// if boolean expression
	root := ConstructAST("test_files/test4.golite")
	sa := New(root, "test_files/test4.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test8(t *testing.T) { 
	// struct var = nil
	root := ConstructAST("test_files/test5.golite")
	sa := New(root, "test_files/test5.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test9(t *testing.T) { 
	// check print
	root := ConstructAST("test_files/test6.golite")
	sa := New(root, "test_files/test6.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test10(t *testing.T) { 
	// delete Struct
	root := ConstructAST("test_files/test7.golite")
	sa := New(root, "test_files/test7.golite")
	if ok := sa.Analyse(); !ok {
		t.Errorf("Expected Sementic Analysis to pass!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test11(t *testing.T) { 
	// check print
	root := ConstructAST("test_files/test8.golite")
	sa := New(root, "test_files/test8.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}

func Test12(t *testing.T) { 
	// check  all control-flow paths have a return statement
	root := ConstructAST("test_files/test9.golite")
	sa := New(root, "test_files/test9.golite")
	if ok := sa.Analyse(); ok {
		t.Errorf("Expected Sementic Analysis to fail!")
	}
	fmt.Println("======Symbol Table======")
	fmt.Println(sa.String())
}