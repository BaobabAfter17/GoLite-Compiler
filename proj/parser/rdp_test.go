package parser

import (
	"fmt"
	token "proj/token"
	"testing"
)

func Test1(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []token.Token{
		{token.PACKAGE, "package", 0},
		{token.ID, "main", 0},
		{token.SEMICOLON, ";", 0},
		{token.IMPORT, "import", 0},
		{token.QUOTATION, "\"", 0},
		{token.FMT, "fmt", 0},
		{token.QUOTATION, "\"", 0},
		{token.SEMICOLON, ";", 0},
		{token.FUNC, "func", 0},
		{token.ID, "main", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.LEFTCURLY, "{", 0},
		{token.VAR, "var", 0},
		{token.ID, "a", 0},
		{token.INT, "int", 0},
		{token.SEMICOLON, ";", 0},
		{token.ID, "a", 0},
		{token.ASSIGN, "=", 0},
		{token.NUMBER, "3", 0},
		{token.PLUS, "+", 0},
		{token.NUMBER, "4", 0},
		{token.ASTEROID, "*", 0},
		{token.NUMBER, "5", 0},
		{token.SLASH, "/", 0},
		{token.NUMBER, "6", 0},
		{token.SEMICOLON, ";", 0},
		{token.FMT, "fmt", 0},
		{token.DOT, ".", 0},
		{token.PRINTLN, "Println", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.ID, "a", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.EOF, "EOF", 0},
	}

	parser := New(tokens, "fake.golite")
	ast, _ := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println(ast)
}


func Test2(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []token.Token{
		{token.PACKAGE, "package", 0},
		{token.ID, "main", 0},
		{token.SEMICOLON, ";", 0},
		{token.IMPORT, "import", 0},
		{token.QUOTATION, "\"", 0},
		{token.FMT, "fmt", 0},
		{token.QUOTATION, "\"", 0},
		{token.SEMICOLON, ";", 0},
		{token.TYPE, "type", 0},
		{token.ID, "Scanner", 0},
		{token.STRUCT, "struct", 0},
		{token.LEFTCURLY, "{", 0},
		{token.ID, "idx", 0},
		{token.INT, "int", 0},
		{token.SEMICOLON, ";", 0},
		{token.ID, "rawInput", 0},
		{token.BOOL, "bool", 0},
		{token.SEMICOLON, ";", 0},
		{token.ID, "p", 0},
		{token.ASTEROID, "*", 0},
		{token.ID, "SomeType", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.SEMICOLON, ";", 0},
		{token.EOF, "EOF", 0},
	}

	parser := New(tokens, "fake.golite")
	ast, _ := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println(ast)
}

func Test3(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []token.Token{
		{token.PACKAGE, "package", 0},
		{token.ID, "main", 0},
		{token.SEMICOLON, ";", 0},
		{token.IMPORT, "import", 0},
		{token.QUOTATION, "\"", 0},
		{token.FMT, "fmt", 0},
		{token.QUOTATION, "\"", 0},
		{token.SEMICOLON, ";", 0},
		{token.FUNC, "func", 0},
		{token.ID, "main", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.LEFTCURLY, "{", 0},
		{token.FMT, "fmt", 0},
		{token.DOT, ".", 0},
		{token.SCAN, "Scan", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.AMPERSAND, "&", 0},
		{token.ID, "variable", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.SEMICOLON, ";", 0},
		{token.FMT, "fmt", 0},
		{token.DOT, ".", 0},
		{token.PRINT, "Print", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.ID, "var1", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.SEMICOLON, ";", 0},
		{token.FMT, "fmt", 0},
		{token.DOT, ".", 0},
		{token.PRINTLN, "Println", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.ID, "var2", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.EOF, "EOF", 0},
	}

	parser := New(tokens, "fake.golite")
	ast, _ := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println(ast)
}

func Test4(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []token.Token{
		{token.PACKAGE, "package", 0},
		{token.ID, "main", 0},
		{token.SEMICOLON, ";", 0},
		{token.IMPORT, "import", 0},
		{token.QUOTATION, "\"", 0},
		{token.FMT, "fmt", 0},
		{token.QUOTATION, "\"", 0},
		{token.SEMICOLON, ";", 0},
		{token.FUNC, "func", 0},
		{token.ID, "main", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.LEFTCURLY, "{", 0},
		{token.IF, "if", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.ID, "a", 0},
		{token.EQUAL, "==", 0},
		{token.NUMBER, "1", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.LEFTCURLY, "{", 0},
		{token.ID, "b", 0},
		{token.ASSIGN, "=", 0},
		{token.NUMBER, "1", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.ELSE, "else", 0},
		{token.LEFTCURLY, "{", 0},
		{token.ID, "b", 0},
		{token.ASSIGN, "=", 0},
		{token.NUMBER, "2", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.EOF, "EOF", 0},
	}

	parser := New(tokens, "fake.golite")
	ast, _ := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println(ast)
}

func Test5(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []token.Token{
		{token.PACKAGE, "package", 0},
		{token.ID, "main", 0},
		{token.SEMICOLON, ";", 0},
		{token.IMPORT, "import", 0},
		{token.QUOTATION, "\"", 0},
		{token.FMT, "fmt", 0},
		{token.QUOTATION, "\"", 0},
		{token.SEMICOLON, ";", 0},
		{token.FUNC, "func", 0},
		{token.ID, "main", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.LEFTCURLY, "{", 0},
		{token.FOR, "for", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.ID, "a", 0},
		{token.EQUAL, "==", 0},
		{token.NUMBER, "1", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.LEFTCURLY, "{", 0},
		{token.ID, "b", 0},
		{token.ASSIGN, "=", 0},
		{token.NUMBER, "1", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.EOF, "EOF", 0},
	}

	parser := New(tokens, "fake.golite")
	ast, _ := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println(ast)
}

func Test6(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []token.Token{
		{token.PACKAGE, "package", 0},
		{token.ID, "main", 0},
		{token.SEMICOLON, ";", 0},
		{token.IMPORT, "import", 0},
		{token.QUOTATION, "\"", 0},
		{token.FMT, "fmt", 0},
		{token.QUOTATION, "\"", 0},
		{token.SEMICOLON, ";", 0},
		{token.FUNC, "func", 0},
		{token.ID, "main", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.ID, "v1", 0},
		{token.INT, "int", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.INT, "int", 0},
		{token.LEFTCURLY, "{", 0},
		{token.RETURN, "return", 0},
		{token.ID, "foo", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.NUMBER, "1", 0},
		{token.COMMA, ",", 0},
		{token.ID, "foo2", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.NUMBER, "1", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.EOF, "EOF", 0},
	}

	parser := New(tokens, "fake.golite")
	ast, _ := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println(ast)
}

func Test7(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []token.Token{
		{token.PACKAGE, "package", 0},
		{token.ID, "main", 0},
		{token.SEMICOLON, ";", 0},
		{token.IMPORT, "import", 0},
		{token.QUOTATION, "\"", 0},
		{token.FMT, "fmt", 0},
		{token.QUOTATION, "\"", 0},
		{token.SEMICOLON, ";", 0},
		{token.FUNC, "func", 0},
		{token.ID, "foo", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.ID, "v2", 0},
		{token.INT, "int", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.INT, "int", 0},
		{token.LEFTCURLY, "{", 0},
		{token.RETURN, "return", 0},
		{token.ID, "sum", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.FUNC, "func", 0},
		{token.ID, "main", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.ID, "v1", 0},
		{token.INT, "int", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.INT, "int", 0},
		{token.LEFTCURLY, "{", 0},
		{token.RETURN, "return", 0},
		{token.ID, "multiple", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.NUMBER, "2", 0},
		{token.COMMA, ",", 0},
		{token.NUMBER, "1", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.EOF, "EOF", 0},
	}

	parser := New(tokens, "fake.golite")
	ast, _ := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println(ast)
}

func Test8(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []token.Token{
		{token.PACKAGE, "package", 0},
		{token.ID, "main", 0},
		{token.SEMICOLON, ";", 0},
		{token.IMPORT, "import", 0},
		{token.QUOTATION, "\"", 0},
		{token.FMT, "fmt", 0},
		{token.QUOTATION, "\"", 0},
		{token.SEMICOLON, ";", 0},
		{token.FUNC, "func", 0},
		{token.ID, "main", 0},
		{token.LEFTPARENTHESIS, "(", 0},
		{token.RIGHTPARENTHESIS, ")", 0},
		{token.LEFTCURLY, "{", 0},
		{token.ID, "b", 0},
		{token.ASSIGN, "=", 0},
		{token.NUMBER, "3", 0},
		{token.GREATERTHAN, ">", 0},
		{token.NUMBER, "4", 0},
		{token.AND, "&&", 0},
		{token.TRUE, "true", 0},
		{token.OR, "||", 0},
		{token.ID, "a", 0},
		{token.GREATEROREQUAL, ">=", 0},
		{token.HYPHEN, "-", 0},
		{token.NUMBER, "5", 0},
		{token.OR, "||", 0},
		{token.HYPHEN, "-", 0},
		{token.ID, "a", 0},
		{token.EQUAL, "==", 0},
		{token.NUMBER, "3", 0},
		{token.SEMICOLON, ";", 0},
		{token.RIGHTCURLY, "}", 0},
		{token.EOF, "EOF", 0},
	}

	parser := New(tokens, "fake.golite")
	ast, _ := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println(ast)
}