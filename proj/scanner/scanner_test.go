package scanner

import (
	"proj/token"
	"testing"
)

type ExpectedResult struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func VerifyTest(t *testing.T, tests []ExpectedResult, scanner *Scanner) {

	for i, tt := range tests {
		tok := scanner.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("FAILED[%d] - incorrect token.\nexpected=%v\ngot=%v\n",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("FAILED[%d] - incorrect token literal.\nexpected=%v\ngot=%v\n",
				i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Type == token.ILLEGAL {
			break
		}
	}
}

func Test1(t *testing.T) {
	input1 := `package main;
import "fmt";

func main() {
	var a int;
	a = 3 + 4 * 5 / 6;
	fmt.Println(a);
}
`

	expected := []ExpectedResult{
		{token.PACKAGE, "package"},
		{token.ID, "main"},
		{token.SEMICOLON, ";"},
		{token.IMPORT, "import"},
		{token.QUOTATION, "\""},
		{token.FMT, "fmt"},
		{token.QUOTATION, "\""},
		{token.SEMICOLON, ";"},
		{token.FUNC, "func"},
		{token.ID, "main"},
		{token.LEFTPARENTHESIS, "("},
		{token.RIGHTPARENTHESIS, ")"},
		{token.LEFTCURLY, "{"},
		{token.VAR, "var"},
		{token.ID, "a"},
		{token.INT, "int"},
		{token.SEMICOLON, ";"},
		{token.ID, "a"},
		{token.ASSIGN, "="},
		{token.NUMBER, "3"},
		{token.PLUS, "+"},
		{token.NUMBER, "4"},
		{token.ASTEROID, "*"},
		{token.NUMBER, "5"},
		{token.SLASH, "/"},
		{token.NUMBER, "6"},
		{token.SEMICOLON, ";"},
		{token.FMT, "fmt"},
		{token.DOT, "."},
		{token.PRINTLN, "Println"},
		{token.LEFTPARENTHESIS, "("},
		{token.ID, "a"},
		{token.RIGHTPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.RIGHTCURLY, "}"},
		{token.EOF, "EOF"},
	}

	scanner := New(input1)
	VerifyTest(t, expected, scanner)
}

func Test2(t *testing.T) {

	input := `
type Scanner struct {
	idx int;
	rawInput bool;
	p *SomeType;
}
`

	expected := []ExpectedResult{
		{token.TYPE, "type"},
		{token.ID, "Scanner"},
		{token.STRUCT, "struct"},
		{token.LEFTCURLY, "{"},
		{token.ID, "idx"},
		{token.INT, "int"},
		{token.SEMICOLON, ";"},
		{token.ID, "rawInput"},
		{token.BOOL, "bool"},
		{token.SEMICOLON, ";"},
		{token.ID, "p"},
		{token.ASTEROID, "*"},
		{token.ID, "SomeType"},
		{token.SEMICOLON, ";"},
		{token.RIGHTCURLY, "}"},
		{token.EOF, "EOF"},
	}

	scanner := New(input)
	VerifyTest(t, expected, scanner)
}

func Test3(t *testing.T) {

	input := `
func foo(a, b, file) bool {
	if (a == 1 && b < 2) {
		fmt.Print(a);
	}  else {
		fmt.Scan(&file);
	}
	return true;
}
`

	expected := []ExpectedResult{
		{token.FUNC, "func"},
		{token.ID, "foo"},
		{token.LEFTPARENTHESIS, "("},
		{token.ID, "a"},
		{token.COMMA, ","},
		{token.ID, "b"},
		{token.COMMA, ","},
		{token.ID, "file"},
		{token.RIGHTPARENTHESIS, ")"},
		{token.BOOL, "bool"},
		{token.LEFTCURLY, "{"},
		{token.IF, "if"},
		{token.LEFTPARENTHESIS, "("},
		{token.ID, "a"},
		{token.EQUAL, "=="},
		{token.NUMBER, "1"},
		{token.AND, "&&"},
		{token.ID, "b"},
		{token.LESSTHAN, "<"},
		{token.NUMBER, "2"},
		{token.RIGHTPARENTHESIS, ")"},
		{token.LEFTCURLY, "{"},
		{token.FMT, "fmt"},
		{token.DOT, "."},
		{token.PRINT, "Print"},
		{token.LEFTPARENTHESIS, "("},
		{token.ID, "a"},
		{token.RIGHTPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.RIGHTCURLY, "}"},
		{token.ELSE, "else"},
		{token.LEFTCURLY, "{"},
		{token.FMT, "fmt"},
		{token.DOT, "."},
		{token.SCAN, "Scan"},
		{token.LEFTPARENTHESIS, "("},
		{token.AMPERSAND, "&"},
		{token.ID, "file"},
		{token.RIGHTPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.RIGHTCURLY, "}"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RIGHTCURLY, "}"},
		{token.EOF, "EOF"},
	}

	scanner := New(input)
	VerifyTest(t, expected, scanner)
}

func Test4(t *testing.T) {

	input := `
	func multiple(a, b) int {
		// take two integers and return their product
		var i, sum int;
		i = 1;
		sum = 0;
		for (i <= b) {
			sum = sum + a;
			i = i + 1;
		}
		return sum;
	}
	multiple(2, 3);
`

	expected := []ExpectedResult{
		{token.FUNC, "func"},
		{token.ID, "multiple"},
		{token.LEFTPARENTHESIS, "("},
		{token.ID, "a"},
		{token.COMMA, ","},
		{token.ID, "b"},
		{token.RIGHTPARENTHESIS, ")"},
		{token.INT, "int"},
		{token.LEFTCURLY, "{"},
		{token.VAR, "var"},
		{token.ID, "i"},
		{token.COMMA, ","},
		{token.ID, "sum"},
		{token.INT, "int"},
		{token.SEMICOLON, ";"},
		{token.ID, "i"},
		{token.ASSIGN, "="},
		{token.NUMBER, "1"},
		{token.SEMICOLON, ";"},
		{token.ID, "sum"},
		{token.ASSIGN, "="},
		{token.NUMBER, "0"},
		{token.SEMICOLON, ";"},
		{token.FOR, "for"},
		{token.LEFTPARENTHESIS, "("},
		{token.ID, "i"},
		{token.LESSOREQUAL, "<="},
		{token.ID, "b"},
		{token.RIGHTPARENTHESIS, ")"},
		{token.LEFTCURLY, "{"},
		{token.ID, "sum"},
		{token.ASSIGN, "="},
		{token.ID, "sum"},
		{token.PLUS, "+"},
		{token.ID, "a"},
		{token.SEMICOLON, ";"},
		{token.ID, "i"},
		{token.ASSIGN, "="},
		{token.ID, "i"},
		{token.PLUS, "+"},
		{token.NUMBER, "1"},
		{token.SEMICOLON, ";"},
		{token.RIGHTCURLY, "}"},
		{token.RETURN, "return"},
		{token.ID, "sum"},
		{token.SEMICOLON, ";"},
		{token.RIGHTCURLY, "}"},
		{token.ID, "multiple"},
		{token.LEFTPARENTHESIS, "("},
		{token.NUMBER, "2"},
		{token.COMMA, ","},
		{token.NUMBER, "3"},
		{token.RIGHTPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	scanner := New(input)
	VerifyTest(t, expected, scanner)
}

func Test5(t *testing.T) {

	input := `
	func main() {
		123.numberMethod();
		nil.nilMethod();
		foo().bar().boo();
		b = !true || false;
		3 > 4;
		a >= b;
		k != 0;
	}
`

	expected := []ExpectedResult{
		{token.FUNC, "func"},
		{token.ID, "main"},
		{token.LEFTPARENTHESIS, "("},
		{token.RIGHTPARENTHESIS, ")"},
		{token.LEFTCURLY, "{"},
		{token.NUMBER, "123"},
		{token.DOT, "."},
		{token.ID, "numberMethod"},
		{token.LEFTPARENTHESIS, "("},
		{token.RIGHTPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.NIL, "nil"},
		{token.DOT, "."},
		{token.ID, "nilMethod"},
		{token.LEFTPARENTHESIS, "("},
		{token.RIGHTPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.ID, "foo"},
		{token.LEFTPARENTHESIS, "("},
		{token.RIGHTPARENTHESIS, ")"},
		{token.DOT, "."},
		{token.ID, "bar"},
		{token.LEFTPARENTHESIS, "("},
		{token.RIGHTPARENTHESIS, ")"},
		{token.DOT, "."},
		{token.ID, "boo"},
		{token.LEFTPARENTHESIS, "("},
		{token.RIGHTPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.ID, "b"},
		{token.ASSIGN, "="},
		{token.EXCLAMATION, "!"},
		{token.TRUE, "true"},
		{token.OR, "||"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "3"},
		{token.GREATERTHAN, ">"},
		{token.NUMBER, "4"},
		{token.SEMICOLON, ";"},
		{token.ID, "a"},
		{token.GREATEROREQUAL, ">="},
		{token.ID, "b"},
		{token.SEMICOLON, ";"},
		{token.ID, "k"},
		{token.NOTEQUAL, "!="},
		{token.NUMBER, "0"},
		{token.SEMICOLON, ";"},
		{token.RIGHTCURLY, "}"},
		{token.EOF, "EOF"},
	}

	scanner := New(input)
	VerifyTest(t, expected, scanner)
}
