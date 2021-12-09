package coordinator

import (
	"os"
	"io/ioutil"
	"fmt"
	"proj/scanner"
	ct "proj/token"
	"proj/parser"
	ast "proj/ast"
	sa "proj/semantic"
	ir "proj/ir"
)

func DoLex(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	// Define a new scanner for some Cal program 
	scanner := scanner.New(string(content))
	for true {
		tok := scanner.NextToken()
		fmt.Printf("Type: %-17s, Literal: %-15s, Line: %-4d\n", tok.Type, tok.Literal, tok.Line)
		if tok.Type == ct.EOF {
			break
		}
	}
}

func ConstructAST(filePath string, toPrint bool) *ast.Program {
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
	if toPrint && ast != nil {
		fmt.Print(ast.String())
	}
	return ast
}

func GenerateILoc(filePath string) {
	root := ConstructAST(filePath, false)
	se := sa.New(root, filePath)
	se.Analyse()
	pFrag := root.TranslateToILOC(se.GetGlobalST())
	
	for _, fun := range pFrag.Funcs {
		for _, ins := range fun.Body {
			switch ins.(type) {
			case *ir.LabelIns:
				fmt.Printf(ins.String() + "\n")
			default:
				fmt.Printf("    " + ins.String() + "\n")
			}
		}
	}
}