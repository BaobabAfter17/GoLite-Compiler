package parser

import (
	"errors"
	"fmt"
	"proj/ast"
	ct "proj/token"
	tp "proj/type"
	"strconv"
)

type Parser struct {
	tokens    []ct.Token
	currIdx   int
	filePath  string
	const_err error
}

//New creates and initializes a new parser
func New(tokens []ct.Token, path string) *Parser {
	res := &Parser{tokens, -1, path, errors.New("ERROR")}
	return res
}

func (p *Parser) currToken() ct.Token {
	return p.tokens[p.currIdx]
}

func (p *Parser) nextToken() ct.Token {

	var token ct.Token
	if p.currIdx == len(p.tokens)-1 {
		token = p.tokens[p.currIdx]
	} else {
		p.currIdx += 1
		token = p.tokens[p.currIdx]
	}
	return token
}

// Returns true if current token matches the given token type
func (p *Parser) match(tokenTyp ct.TokenType) (ct.Token, bool) {
	if tokenTyp == p.currToken().Type {
		token := p.currToken()
		p.nextToken()
		return token, true
	}
	if p.currToken().Type == ct.ILLEGAL {
		p.parseError("illegal token", p.currToken().Line)
	}

	return ct.Token{ct.ILLEGAL, "", p.currToken().Line}, false
}

// Set the currIdx to the given idx
func (p *Parser) setIdx(idx int) {
	p.currIdx = idx
}

// Decrement the currIdx to the given idx
func (p *Parser) DecIdx() {
	if p.currIdx > 0 {
		p.currIdx--
	}
}

func (p *Parser) parseError(msg string, errLine int) {
	// fmt.Errorf("%s:%d: syntax error: %s", p.filePath, errLine, msg)
	fmt.Printf("%s:%d: syntax error: %s\n", p.filePath, errLine, msg)
}

func (p *Parser) Parse() (*ast.Program, error) {
	p.nextToken()
	return p.parseProgram()
}

func (p *Parser) parseProgram() (*ast.Program, error) {
	pg := &ast.Program{}
	var err error
	pg.Pkg, err = parsePackage(p)
	if err != nil {
		return nil, err
	}
	pg.Impt, err = parseImport(p)
	if err != nil {
		return nil, err
	}
	pg.Typs, err = parseTypeDeclarations(p)
	if err != nil {
		return nil, err
	}
	pg.Dcl, err = parseVarDeclarations(p)
	if err != nil {
		return nil, err
	}
	pg.Funcs, err = parseFunctions(p)
	if err != nil {
		return nil, err
	}
	return pg, err
}

func parseFunctions(p *Parser) ([]*ast.Function, error) {
	var res []*ast.Function
	for p.currToken().Type == ct.FUNC {
		fun, err := parseFunction(p)
		if err != nil {
			return nil, p.const_err
		}

		res = append(res, fun)
	}

	return res, nil
}

func parsePackage(p *Parser) (*ast.Package, error) {
	pkg := &ast.Package{}
	pkg.Line = p.currToken().Line
	tk, correct := p.match(ct.PACKAGE)
	if !correct {
		p.parseError("Package definition not found.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.ID)
	if !correct {
		p.parseError("Package name not found.", tk.Line)
		return nil, p.const_err
	}
	pkg.Name = tk.Literal

	tk, correct = p.match(ct.SEMICOLON)
	if !correct {
		p.parseError("Semicolon not found.", tk.Line)
		return nil, p.const_err
	}

	return pkg, nil
}

func parseImport(p *Parser) (*ast.Import, error) {
	impt := &ast.Import{}
	tk, correct := p.match(ct.IMPORT)
	if !correct {
		p.parseError("Import definition not found.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.QUOTATION)
	if !correct {
		p.parseError("Quotation not found.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.FMT)
	if !correct {
		p.parseError("Only fmt is allowed to import.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.QUOTATION)
	if !correct {
		p.parseError("Quotation not found.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.SEMICOLON)
	if !correct {
		p.parseError("Semicolon not found.", tk.Line)
		return nil, p.const_err
	}

	return impt, nil
}

func parseTypeDeclarations(p *Parser) (*ast.Types, error) {
	var err error
	typs := &ast.Types{}
	for true {
		tk, correct := p.match(ct.TYPE)
		if !correct {
			return typs, nil
		}

		typdcl := &ast.TypeDeclaration{}
		typdcl.Line = p.currToken().Line
		tk, correct = p.match(ct.ID)
		if !correct {
			p.parseError("The name of struct not found.", tk.Line)
			return nil, p.const_err
		}
		typdcl.Name = tk.Literal

		tk, correct = p.match(ct.STRUCT)
		if !correct {
			p.parseError("The keyword struct not found.", tk.Line)
			return nil, p.const_err
		}

		tk, correct = p.match(ct.LEFTCURLY)
		if !correct {
			p.parseError("The left curly bracket not found.", tk.Line)
			return nil, p.const_err
		}

		typdcl.Fields, err = parseFields(p)
		if err != nil {
			return nil, p.const_err
		}
		
		typs.TypDcls = append(typs.TypDcls, typdcl)

		tk, correct = p.match(ct.RIGHTCURLY)
		if !correct {
			p.parseError("The right curly bracket not found.", tk.Line)
			return nil, p.const_err
		}

		tk, correct = p.match(ct.SEMICOLON)
		if !correct {
			p.parseError("Semicolon not found.", tk.Line)
			return nil, p.const_err
		}
	}
	return nil, nil
}

func parseType(p *Parser) (tp.Type, error) {
	tk, correct := p.match(ct.INT)
	if correct {
		return tp.IntType, nil
	}

	tk, correct = p.match(ct.BOOL)
	if correct {
		return tp.BoolType, nil
	}

	tk, correct = p.match(ct.ASTEROID)
	if !correct {
		p.parseError("The type of field only contains int, bool and pointer.", tk.Line)
		return tp.UnknownType, p.const_err
	}

	tk, correct = p.match(ct.ID)
	if !correct {
		p.parseError("The name of pointer not found.", tk.Line)
		return tp.UnknownType, p.const_err
	}

	return &tp.StructType{Name: tk.Literal}, nil
}

func parseField(p *Parser) (*ast.Field, error) {
	var err error
	res := &ast.Field{}
	res.Line = p.currToken().Line
	tk, correct := p.match(ct.ID)
	if !correct {
		p.parseError("Field name not found.", tk.Line)
		return nil, p.const_err
	}
	res.Name = tk.Literal
	res.Typ, err = parseType(p)
	if err != nil {
		return nil, p.const_err
	}

	tk, correct = p.match(ct.SEMICOLON)
	if !correct {
		p.parseError("Semicolon not found.", tk.Line)
		return nil, p.const_err
	}

	return res, nil
}

func parseFields(p *Parser) ([]*ast.Field, error) {
	var res []*ast.Field
	firstField, err := parseField(p)
	if err != nil {
		return nil, p.const_err
	}
	res = append(res, firstField)

	for p.currToken().Type == ct.ID {
		
		nextField, err := parseField(p)
		if err != nil {
			return res, nil
		}
		res = append(res, nextField)
	}
	return res, nil
}

func parseVarDeclaration(p *Parser) (*ast.VarDeclaration, error) {
	var err error
	res := &ast.VarDeclaration{}
	res.Line = p.currToken().Line
	tk, correct := p.match(ct.VAR)
	if !correct {
		p.parseError("Expect var keyword.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.ID)
	if !correct {
		p.parseError("Expect variable identifier.", tk.Line)
		return nil, p.const_err
	}
	res.Names = append(res.Names, tk.Literal)

	for true {
		tk, correct := p.match(ct.COMMA)
		if !correct {
			break
		}

		tk, correct = p.match(ct.ID)
		if !correct {
			p.parseError("Expect variable identifier after comma.", tk.Line)
			return nil, p.const_err
		}

		res.Names = append(res.Names, tk.Literal)
	}

	res.Typ, err = parseType(p)
	if err != nil {
		return nil, p.const_err
	}

	tk, correct = p.match(ct.SEMICOLON)
	if !correct {
		p.parseError("Semicolon not found.", tk.Line)
		return nil, p.const_err
	}

	return res, nil
}

func parseVarDeclarations(p *Parser) ([]*ast.VarDeclaration, error) {
	var res []*ast.VarDeclaration
	for true {
		_, correct := p.match(ct.VAR)
		if !correct {
			return res, nil
		}
		p.DecIdx()
		varDcl, err := parseVarDeclaration(p)
		if err != nil {
			return nil, p.const_err
		}
		res = append(res, varDcl)
	}
	return nil, nil
}

func parseLValue(p *Parser) (ast.Expression, error) {

	tk, correct := p.match(ct.ID)
	if !correct {
		return nil, errors.New("Lvalue must start with an identifier.")
	}
	left := &ast.IdenLiteral{tk, tk.Literal}

	return parseSelectorTermPrime(p, left)
}

func parseAssignment(p *Parser) (*ast.AssignmentSt, error) {
	var err error
	res := &ast.AssignmentSt{}
	res.LVal, err = parseLValue(p)
	if err != nil {
		return nil, err
	}

	tk, correct := p.match(ct.ASSIGN)
	if !correct {
		return nil, errors.New("The symbol '=' not found")
	}

	res.Expr, err = parseExpression(p)
	if err != nil {
		return nil, p.const_err
	}

	tk, correct = p.match(ct.SEMICOLON)
	if !correct {
		p.parseError("Semicolon not found.", tk.Line)
		return nil, p.const_err
	}
	res.Line = tk.Line
	return res, nil
}

func parsePrint(p *Parser) (*ast.PrintSt, error) {
	res := &ast.PrintSt{}

	p.match(ct.FMT)
	p.match(ct.DOT)

	tk, correct := p.match(ct.PRINT)
	if !correct {
		tk, _ = p.match(ct.PRINTLN)
	}
	res.MethodName = tk.Literal

	tk, correct = p.match(ct.LEFTPARENTHESIS)
	if !correct {
		p.parseError("Expect left parathesis after print.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.ID)
	if !correct {
		p.parseError("Expect identifier after print.", tk.Line)
		return nil, p.const_err
	}
	res.Iden = tk.Literal

	tk, correct = p.match(ct.RIGHTPARENTHESIS)
	if !correct {
		p.parseError("Expect right parathesis after print.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.SEMICOLON)
	if !correct {
		p.parseError("Expect semicolon at the end.", tk.Line)
		return nil, p.const_err
	}
	res.Line = tk.Line
	return res, nil
}

func parseRead(p *Parser) (*ast.ReadSt, error) {
	res := &ast.ReadSt{}

	p.match(ct.FMT)
	p.match(ct.DOT)
	p.match(ct.SCAN)

	tk, correct := p.match(ct.LEFTPARENTHESIS)
	if !correct {
		p.parseError("Expect '(' after scan keyword.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.AMPERSAND)
	if !correct {
		p.parseError("Expect '&' before the identifier.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.ID)
	if !correct {
		p.parseError("Expect identifier for reading.", tk.Line)
		return nil, p.const_err
	}
	res.Iden = tk.Literal

	tk, correct = p.match(ct.RIGHTPARENTHESIS)
	if !correct {
		p.parseError("Expect right ')' after the identifier.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.SEMICOLON)
	if !correct {
		p.parseError("Expect semicolon at the end.", tk.Line)
		return nil, p.const_err
	}
	res.Line = tk.Line
	return res, nil
}

func parseReturn(p *Parser) (ast.Statement, error) {
	p.match(ct.RETURN)
	res := &ast.ReturnSt{}
	res.Expr = &ast.NilLiteral{p.currToken()}
	res.Line = p.currToken().Line
	if p.currToken().Type == ct.SEMICOLON {
		p.nextToken()
		return res, nil
	} else {
		expr, err := parseExpression(p)
		if err != nil {
			return nil, p.const_err
		}
		res.Expr = expr

		tk, correct := p.match(ct.SEMICOLON)
		if !correct {
			p.parseError("Expect semicolon at the end.", tk.Line)
		}

		return res, nil
	}
}

func parseStatement(p *Parser) (ast.Statement, error) {
	switch p.currToken().Type {
	case ct.LEFTCURLY:
		return parseBlock(p)
	case ct.FMT:
		p.nextToken()
		tk, correct := p.match(ct.DOT)
		if !correct {
			p.parseError("Expect '.' after fmt.", tk.Line)
			return nil, p.const_err
		}
		switch p.currToken().Type {
		case ct.PRINT:
			p.DecIdx()
			p.DecIdx()
			return parsePrint(p)
		case ct.PRINTLN:
			p.DecIdx()
			p.DecIdx()
			return parsePrint(p)
		case ct.SCAN:
			p.DecIdx()
			p.DecIdx()
			return parseRead(p)
		default:
			p.parseError("Expect 'Scan', 'Print' or 'Println' methods after fmt.", tk.Line)
			return nil, p.const_err
		}
	case ct.RETURN:
		return parseReturn(p)
	case ct.IF:
		return parseConditional(p)
	case ct.FOR:
		return parseLoop(p)
	case ct.ID:
		p.nextToken()
		_, correct := p.match(ct.LEFTPARENTHESIS)
		if correct {
			p.DecIdx()
			p.DecIdx()
			return parseInvocation(p)
		} else {
			p.DecIdx()
			return parseAssignment(p)
		}
	default:
		p.parseError("Not a valid statement.", p.currToken().Line)
		return nil, p.const_err
	}
}

func parseBlock(p *Parser) (*ast.Block, error) {
	p.match(ct.LEFTCURLY)
	res := &ast.Block{}
	for true {
		_, correct := p.match(ct.RIGHTCURLY)
		if correct {
			return res, nil
		}
		st, err := parseStatement(p)
		if err != nil {
			return nil, p.const_err
		}
		res.Stmnts = append(res.Stmnts, st)
	}
	return nil, nil
}

func parseReturnType(p *Parser) (tp.Type, error) {
	_, correct := p.match(ct.LEFTCURLY)
	if !correct {
		return parseType(p)
	} else {
		p.DecIdx()
		return tp.VoidType, nil
	}
}

func parseLoop(p *Parser) (*ast.LoopSt, error) {
	p.match(ct.FOR)

	res := &ast.LoopSt{}
	res.Line = p.currToken().Line
	tk, correct := p.match(ct.LEFTPARENTHESIS)
	if !correct {
		p.parseError("Expect '(' after for.", tk.Line)
		return nil, p.const_err
	}

	expr, err := parseExpression(p)
	if err != nil {
		return nil, p.const_err
	}
	res.Condition = expr

	tk, correct = p.match(ct.RIGHTPARENTHESIS)
	if !correct {
		p.parseError("Expect ')' after the expression.", tk.Line)
		return nil, p.const_err
	}

	blk, err := parseBlock(p)
	if err != nil {
		return nil, p.const_err
	}
	res.LoopBlock = blk

	return res, nil
}

func parseConditional(p *Parser) (*ast.ConditionalSt, error) {
	p.match(ct.IF)

	tk, correct := p.match(ct.LEFTPARENTHESIS)
	if !correct {
		p.parseError("Expect '(' after if keyword.", tk.Line)
		return nil, p.const_err
	}

	res := &ast.ConditionalSt{}
	expr, err := parseExpression(p)
	if err != nil {
		return nil, p.const_err
	}
	res.Condition = expr

	tk, correct = p.match(ct.RIGHTPARENTHESIS)
	if !correct {
		p.parseError("Expect ')' after the expression.", tk.Line)
		return nil, p.const_err
	}

	blk, err := parseBlock(p)
	if err != nil {
		return nil, p.const_err
	}
	res.IfBlock = blk

	if p.currToken().Type == ct.ELSE {
		p.nextToken()
		blk, err = parseBlock(p)
		if err != nil {
			return nil, p.const_err
		}
		res.ElseBlock = blk
	}
	res.Line = tk.Line
	return res, nil
}

func parseArguments(p *Parser) ([]ast.Expression, error) {
	tk, correct := p.match(ct.LEFTPARENTHESIS)
	if !correct {
		p.parseError("Expect '(' before arguments.", tk.Line)
		return nil, p.const_err
	}

	if p.currToken().Type == ct.RIGHTPARENTHESIS {
		p.nextToken()
		return []ast.Expression{}, nil
	} else {
		var res []ast.Expression
		for true {
			expr, err := parseExpression(p)
			if err != nil {
				return nil, p.const_err
			}
			res = append(res, expr)
			if p.currToken().Type == ct.RIGHTPARENTHESIS {
				p.match(ct.RIGHTPARENTHESIS)
				return res, nil
			}
			tk, correct = p.match(ct.COMMA)
			if !correct {
				p.parseError("Expect ',' between arguments.", tk.Line)
				return nil, p.const_err
			}
		}
	}
	return nil, nil
}

func parseParameter(p *Parser) (*ast.Parameter, error) {
	res := &ast.Parameter{}
	tk, correct := p.match(ct.ID)
	if !correct {
		p.parseError("Expect identifer for a parameter.", tk.Line)
		return nil, p.const_err
	}
	res.Name = tk.Literal
	
	typ, err := parseType(p)
	if err != nil {
		return nil, p.const_err
	}
	res.Typ = typ

	return res, nil
}

func parseParameters(p *Parser) ([]*ast.Parameter, error) {
	tk, correct := p.match(ct.LEFTPARENTHESIS)
	if !correct {
		p.parseError("Expect '(' before parameters.", tk.Line)
		return nil, p.const_err
	}

	if p.currToken().Type == ct.RIGHTPARENTHESIS {
		p.nextToken()
		return []*ast.Parameter{}, nil
	} else {
		var res []*ast.Parameter
		for true {
			dcl, err := parseParameter(p)
			if err != nil {
				return nil, p.const_err
			}
			res = append(res, dcl)
			if p.currToken().Type == ct.RIGHTPARENTHESIS {
				p.match(ct.RIGHTPARENTHESIS)
				return res, nil
			}
			tk, correct = p.match(ct.COMMA)
			if !correct {
				p.parseError("Expect ',' between parameters.", tk.Line)
				return nil, p.const_err
			}
		}
	}
	return nil, nil
}

func parseInvocation(p *Parser) (*ast.InvocationSt, error) {
	var err error
	res := &ast.InvocationSt{}
	tk, _ := p.match(ct.ID)
	res.FuncName = tk.Literal

	res.Arguments, err = parseArguments(p)
	if err != nil {
		return nil, p.const_err
	}

	tk, correct := p.match(ct.SEMICOLON)
	if !correct {
		p.parseError("Expect ';' after invocation statement.", tk.Line)
		return nil, p.const_err
	}
	res.Line = tk.Line
	return res, nil
}

func parseFunction(p *Parser) (*ast.Function, error) {
	var err error
	res := &ast.Function{}
	tk, correct := p.match(ct.FUNC)
	if !correct {
		p.parseError("Expect keyword func.", tk.Line)
		return nil, p.const_err
	}

	tk, correct = p.match(ct.ID)
	if !correct {
		p.parseError("Expect func name.", tk.Line)
		return nil, p.const_err
	}
	res.Name = tk.Literal

	res.Parameters, err = parseParameters(p)
	if err != nil {
		return nil, p.const_err
	}

	res.ReturnTyp, err = parseReturnType(p)
	if err != nil {
		return nil, p.const_err
	}

	tk, correct = p.match(ct.LEFTCURLY)
	if !correct {
		p.parseError("Expect left curly bracket after function signature.", tk.Line)
		return nil, p.const_err
	}

	res.Dcl, err = parseVarDeclarations(p)
	if err != nil {
		return nil, p.const_err
	}

	for true {
		if p.currToken().Type == ct.RIGHTCURLY {
			break
		}
		st, err := parseStatement(p)
		if err != nil {
			return nil, p.const_err
		}
		res.Stmnts = append(res.Stmnts, st)
	}

	tk, correct = p.match(ct.RIGHTCURLY)
	if !correct {
		p.parseError("Expect right curly bracket after function statements.", tk.Line)
		return nil, p.const_err
	}
	res.Line = tk.Line
	return res, nil

}

func parseFactor(p *Parser) (ast.Expression, error) {
	var err error
	switch p.currToken().Type {
	case ct.NIL:
		res := &ast.NilLiteral{p.currToken()}
		p.nextToken()
		return res, nil
	case ct.TRUE:
		res := &ast.BoolLiteral{p.currToken(), true}
		p.nextToken()
		return res, nil
	case ct.FALSE:
		res := &ast.BoolLiteral{p.currToken(), false}
		p.nextToken()
		return res, nil
	case ct.NUMBER:
		num, _ := strconv.Atoi(p.currToken().Literal)
		res := &ast.IntLiteral{p.currToken(), num}
		p.nextToken()
		return res, nil
	case ct.LEFTPARENTHESIS:
		p.match(ct.LEFTPARENTHESIS)
		expr, err := parseExpression(p)
		if err == nil {
			tk, correct := p.match(ct.RIGHTPARENTHESIS)
			if !correct {
				p.parseError("Expect ')' after the expression.", tk.Line)
				return nil, p.const_err
			} else {
				return expr, nil
			}
		} else {
			return nil, p.const_err
		}
	case ct.ID:
		tk, _ := p.match(ct.ID)
		if p.currToken().Type != ct.LEFTPARENTHESIS {
			return &ast.IdenLiteral{tk, tk.Literal}, nil
		} else {
			res := &ast.InvocationExpr{}
			res.Line = p.currToken().Line
			res.Func = &ast.IdenLiteral{tk, tk.Literal}
			res.Arguments, err = parseArguments(p)
			if err != nil {
				return nil, p.const_err
			} else {
				return res, nil
			}
		}
	default:
		p.parseError("Not a valid factor.", p.currToken().Line)
		return nil, p.const_err
	}
}

func parseSelectorTerm(p *Parser) (ast.Expression, error) {
	fac, err := parseFactor(p)
	if err != nil {
		return nil, p.const_err
	}

	return parseSelectorTermPrime(p, fac)
}

func parseSelectorTermPrime(p *Parser, left ast.Expression) (ast.Expression, error) {
	res := &ast.BinOpExpr{}
	res.Left = left
	switch p.currToken().Type {
	case ct.DOT:
		res.Token = p.currToken()
		res.Line = p.currToken().Line
		res.Operator = ast.ACCESS
		p.nextToken()
		tk, correct := p.match(ct.ID)
		if !correct {
			return nil, p.const_err
		}
		res.Right = &ast.IdenLiteral{tk, tk.Literal}
	default:
		return left, nil
	}
	return parseSelectorTermPrime(p, res)
}

func parseUnaryTerm(p *Parser) (ast.Expression, error) {
	switch p.currToken().Type {
	case ct.EXCLAMATION:
		res := &ast.UnOpExpr{p.currToken(), ast.NOT, nil, 0}
		res.Line = p.currToken().Line
		p.nextToken()
		selTerm, err := parseSelectorTerm(p)
		if err != nil {
			return nil, p.const_err
		}
		res.Left = selTerm
		return res, nil
	case ct.HYPHEN:
		res := &ast.UnOpExpr{p.currToken(), ast.SUB, nil, 0}
		res.Line = p.currToken().Line
		p.nextToken()
		selTerm, err := parseSelectorTerm(p)
		if err != nil {
			return nil, p.const_err
		}
		res.Left = selTerm
		return res, nil
	default:
		return parseSelectorTerm(p)
	}
	return nil, nil
}

func parseTerm(p *Parser) (ast.Expression, error) {
	left, err := parseUnaryTerm(p)
	if err != nil {
		return nil, p.const_err
	}
	return parseTermPrime(p, left)
}

func parseTermPrime(p *Parser, left ast.Expression) (ast.Expression, error) {
	var err error
	res := &ast.BinOpExpr{}
	res.Line = p.currToken().Line
	res.Left = left
	switch p.currToken().Type {
	case ct.ASTEROID:
		res.Token = p.currToken()
		res.Operator = ast.MULT
		p.nextToken()
		res.Right, err = parseUnaryTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	case ct.SLASH:
		res.Token = p.currToken()
		res.Operator = ast.DIV
		p.nextToken()
		res.Right, err = parseUnaryTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	default:
		return left, nil
	}
	return parseTermPrime(p, res)
}

func parseSimpleTerm(p *Parser) (ast.Expression, error) {
	left, err := parseTerm(p)
	if err != nil {
		return nil, p.const_err
	}
	return parseSimpleTermPrime(p, left)
}

func parseSimpleTermPrime(p *Parser, left ast.Expression) (ast.Expression, error) {
	var err error
	res := &ast.BinOpExpr{}
	res.Line = p.currToken().Line
	res.Left = left
	switch p.currToken().Type {
	case ct.PLUS:
		res.Token = p.currToken()
		res.Operator = ast.ADD
		p.nextToken()
		res.Right, err = parseTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	case ct.HYPHEN:
		res.Token = p.currToken()
		res.Operator = ast.SUB
		p.nextToken()
		res.Right, err = parseTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	default:
		return left, nil
	}
	return parseSimpleTermPrime(p, res)
}

func parseRelationTerm(p *Parser) (ast.Expression, error) {
	left, err := parseSimpleTerm(p)
	if err != nil {
		return nil, p.const_err
	}
	return parseRelationTermPrime(p, left)
}

func parseRelationTermPrime(p *Parser, left ast.Expression) (ast.Expression, error) {
	var err error
	res := &ast.BinOpExpr{}
	res.Line = p.currToken().Line
	res.Left = left
	res.Token = p.currToken()
	switch p.currToken().Type {
	case ct.GREATERTHAN:
		res.Operator = ast.GT
		p.nextToken()
		res.Right, err = parseSimpleTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	case ct.LESSTHAN:
		res.Operator = ast.LT
		p.nextToken()
		res.Right, err = parseSimpleTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	case ct.GREATEROREQUAL:
		res.Operator = ast.GE
		p.nextToken()
		res.Right, err = parseSimpleTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	case ct.LESSOREQUAL:
		res.Operator = ast.LE
		p.nextToken()
		res.Right, err = parseSimpleTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	default:
		return left, nil
	}
	return parseRelationTermPrime(p, res)
}

func parseEqualTerm(p *Parser) (ast.Expression, error) {
	left, err := parseRelationTerm(p)
	if err != nil {
		return nil, p.const_err
	}
	return parseEqualTermPrime(p, left)
}

func parseEqualTermPrime(p *Parser, left ast.Expression) (ast.Expression, error) {
	var err error
	res := &ast.BinOpExpr{}
	res.Line = p.currToken().Line
	res.Left = left
	res.Token = p.currToken()
	switch p.currToken().Type {
	case ct.EQUAL:
		res.Operator = ast.EQ
		p.nextToken()
		res.Right, err = parseRelationTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	case ct.NOTEQUAL:
		res.Operator = ast.NE
		p.nextToken()
		res.Right, err = parseRelationTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	default:
		return left, nil
	}
	return parseEqualTermPrime(p, res)
}

func parseBoolTerm(p *Parser) (ast.Expression, error) {
	left, err := parseEqualTerm(p)
	if err != nil {
		return nil, p.const_err
	}
	return parseBoolTermPrime(p, left)
}

func parseBoolTermPrime(p *Parser, left ast.Expression) (ast.Expression, error) {
	var err error
	res := &ast.BinOpExpr{}
	res.Line = p.currToken().Line
	res.Left = left
	res.Token = p.currToken()
	switch p.currToken().Type {
	case ct.AND:
		res.Operator = ast.AND
		p.nextToken()
		res.Right, err = parseEqualTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	default:
		return left, nil
	}
	return parseBoolTermPrime(p, res)
}

func parseExpression(p *Parser) (ast.Expression, error) {
	left, err := parseBoolTerm(p)
	if err != nil {
		return nil, p.const_err
	}
	return parseExpressionPrime(p, left)
}

func parseExpressionPrime(p *Parser, left ast.Expression) (ast.Expression, error) {
	var err error
	res := &ast.BinOpExpr{}
	res.Line = p.currToken().Line
	res.Left = left
	res.Token = p.currToken()
	switch p.currToken().Type {
	case ct.OR:
		res.Operator = ast.OR
		p.nextToken()
		res.Right, err = parseBoolTerm(p)
		if err != nil {
			return nil, p.const_err
		}
	default:
		return left, nil
	}
	return parseExpressionPrime(p, res)
}
