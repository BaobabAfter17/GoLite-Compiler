package semantic

import (
	"fmt"
	"proj/ast"
	st "proj/symTable"
	tp "proj/type"
)

type SemanticAnalysis struct {
	globalST  *st.SymbolTable
	currScope *st.SymbolTable
	root      *ast.Program
	filePath  string
}

func New(root *ast.Program, filePath string) *SemanticAnalysis {
	globalST := &st.SymbolTable{
		Attributes: make(map[string]st.Entry),
	}
	
	globalST.Insert("new", &st.FuncEntry{})
	globalST.Insert("delete", &st.FuncEntry{})

	return &SemanticAnalysis{globalST, globalST, root, filePath}
}

func (sa *SemanticAnalysis) String() string {
	return sa.globalST.String()
}

func (sa *SemanticAnalysis) GetGlobalST() *st.SymbolTable {
	return sa.globalST
}

func (sa *SemanticAnalysis) Analyse() bool {
	return sa.buildSymTable() && sa.typeCheck()
}

func (sa *SemanticAnalysis) buildSymTable() bool {

	root := sa.root
	res := true
	res = res && sa.parsePackage(root.Pkg)
	res = res && sa.parseTypes(root.Typs)
	res = res && sa.parseDeclarations(root.Dcl)
	res = res && sa.parseFunctions(root.Funcs)
	return res
}

func (sa *SemanticAnalysis) parsePackage(pkg *ast.Package) bool {
	if pkg == nil {
		sa.semanticError("Package definition not found", 0)
		return false
	}
	if pkg.Name != "main" {
		sa.semanticError("only package main is allowed", pkg.Line)
		return false
	}
	return true
}

// parse root.Typs
func (sa *SemanticAnalysis) parseTypes(tps *ast.Types) bool {

	if tps == nil {
		return true
	}
	var td *ast.TypeDeclaration
	for i := 0; i < len(tps.TypDcls); i++ {
		td = tps.TypDcls[i]
		if !sa.parseTypeDeclaration(td) {
			return false
		}
	}
	return true
}

func (sa *SemanticAnalysis) parseTypeDeclaration(td *ast.TypeDeclaration) bool {
	if td == nil {
		return true
	}

	name := td.Name
	if sa.lookupAll(name) != nil {
		sa.semanticError(name + " already defined.", td.Line)
		return false
	}

	newStructType := &tp.StructType{}
	newStructType.Name = name
	newStructType.Fields = make(map[string]tp.Type)
	strct := &st.StructEntry{newStructType}
	sa.insert(name, strct)
	if !sa.structFromFields(td.Fields, strct) { return false }

	return true
}

func (sa *SemanticAnalysis) insert(name string, entry st.Entry) {
	sa.currScope.Insert(name, entry)
}

func (sa *SemanticAnalysis) lookupLocal(name string) st.Entry {
	return sa.currScope.LookUpLocal(name)
}

func (sa *SemanticAnalysis) lookupAll(name string) st.Entry {
	return sa.currScope.LookUpAll(name)
}

func (sa *SemanticAnalysis) structFromFields(fields []*ast.Field, strct *st.StructEntry) bool {

	newStructType := strct.Typ
	
	for _, f := range fields {
		fieldName := f.Name
		if newStructType.Contains(fieldName) {
			sa.semanticError(fieldName + " already defined in struct", f.Line)
			return false
		}
		
		switch f.Typ.(type) {
		case *tp.StructType:
			completeStructEntry := sa.lookupAll(f.Typ.(*tp.StructType).Name)
			if completeStructEntry == nil {
				sa.semanticError("type " + f.Typ.(*tp.StructType).Name + " not defined", f.Line)
				return false
			} else {
				newStructType.Insert(fieldName, completeStructEntry.(*st.StructEntry).Typ)
			}
		default:
			newStructType.Insert(fieldName, f.Typ)
		}
	}

	return true
}

// parse root.Dcl
func (sa *SemanticAnalysis) parseDeclarations(dcls []*ast.VarDeclaration) bool {
	if dcls == nil {
		return true
	}
	res := true
	for _, dcl := range dcls {
		if !sa.parseDeclaration(dcl) {
			res = false
		}
	}
	return res
}

func (sa *SemanticAnalysis) parseDeclaration(dcl *ast.VarDeclaration) bool {
	varType := dcl.Typ
	for _, name := range dcl.Names {
		entry := sa.lookupAll(name)
		if entry != nil {
			sa.semanticError("variable " + name + " already defined.", dcl.Line)
			return false
		}

		switch varType.(type) {
		case *tp.StructType:
			completeStructEntry := sa.lookupAll(varType.(*tp.StructType).Name[1:])
			if completeStructEntry == nil {
				sa.semanticError("type " + varType.(*tp.StructType).Name + " not defined.", dcl.Line)
				return false
			} else {
				sa.insert(name, &st.VarEntry{Typ: completeStructEntry.(*st.StructEntry).Typ})
			}
		default:
			sa.insert(name, &st.VarEntry{Typ: varType})
		}

	}
	return true
}

func (sa *SemanticAnalysis) parseParameters(paras []*ast.Parameter) bool {
	for _, p := range paras {
		if sa.parseParameter(p) != true {
			return false
		}
	}
	return true
}

func (sa *SemanticAnalysis) parseParameter(para *ast.Parameter) bool {
	entry := sa.lookupLocal(para.Name)
	if entry != nil {
		sa.semanticError("parameter " + para.Name + " already defined.", para.Line)
		return false
	} else {
		typ := para.Typ
		switch typ.(type) {
		case *tp.StructType:
			completeStructEntry := sa.lookupAll(typ.(*tp.StructType).Name)
			if completeStructEntry != nil {
				sa.insert(para.Name, &st.VarEntry{Typ: completeStructEntry.(*st.StructEntry).Typ})
			} else {
				sa.semanticError("Type " + typ.(*tp.StructType).Name + " not defined.", para.Line)
				return false
			}
		default:
			sa.insert(para.Name, &st.VarEntry{Typ: typ})
		}
	}

	return true
}


// parse root.Funcs
func (sa *SemanticAnalysis) parseFunctions(funcs []*ast.Function) bool {
	if funcs == nil {
		return true
	}
	for _, fn := range funcs {
		if !sa.parseFunction(fn) {
			return false
		}
	}
	return true
}

func (sa *SemanticAnalysis) parseFunction(fn *ast.Function) bool {
	name := fn.Name
	entry := sa.lookupAll(name)
	switch entry.(type) {
	case *st.FuncEntry:
		sa.semanticError("function " + name + " has already been defined!", fn.Line)
		return false
	case *st.VarEntry:
		sa.semanticError("function " + name + " has been used as a variable!", fn.Line)
		return false
	}

	sa.currScope = sa.currScope.InitializeScope(name)
	fentry := sa.lookupAll(name).(*st.FuncEntry)
	fentry.RetTyp = fn.ReturnTyp
	for _, p := range fn.Parameters {
		fentry.ParaNames = append(fentry.ParaNames, p.Name)
	}
	
	defer func() { sa.currScope = sa.currScope.FinalizeScope() }()
	ok := sa.parseParameters(fn.Parameters)
	if !ok { return false }

	ok = sa.parseLocalDeclarations(fn.Dcl)
	if !ok { return false }

	return true
}

func (sa *SemanticAnalysis) parseLocalDeclarations(dcls []*ast.VarDeclaration) bool {
	res := true
	for _, dcl := range dcls {
		ok := sa.parseLocalDeclaration(dcl)
		if !ok { res = false }
	}
	return res
}

func (sa *SemanticAnalysis) parseLocalDeclaration(dcl *ast.VarDeclaration) bool {
	varType := dcl.Typ
	for _, name := range dcl.Names {
		entry := sa.lookupLocal(name)
		if entry != nil {
			sa.semanticError("local variable " + name + " already defined.", dcl.Line)
			return false
		}

		switch varType.(type) {
		case *tp.StructType:
			completeStructEntry := sa.lookupAll(varType.(*tp.StructType).Name)
			if completeStructEntry == nil {
				sa.semanticError("type " + varType.(*tp.StructType).Name + " not defined.", dcl.Line)
				return false
			} else {
				sa.insert(name, &st.VarEntry{Typ: completeStructEntry.(*st.StructEntry).Typ})
			}
		default:
			sa.insert(name, &st.VarEntry{Typ: varType})
		}

	}
	return true
}

// ===================
// Type Check
func (sa *SemanticAnalysis) typeCheck() bool {
	return sa.checkProgram(sa.root)
}

func (sa *SemanticAnalysis) checkProgram(prog *ast.Program) bool {
	return sa.checkFunctions(prog.Funcs)
}

func (sa *SemanticAnalysis) checkFunctions(funcs []*ast.Function) bool {
	res := true
	if mainf := sa.getMainFunction(funcs); mainf == nil {
		sa.semanticError("main function not defined!", 0)
		res = false
	} else if mainf.Parameters != nil {
		if len(mainf.Parameters) != 0 {
			sa.semanticError("main function should not have parameters.", mainf.Line)
			res = false
		}
	}
	for _, fn := range funcs {
		if ok := sa.checkFunction(fn); !ok {
			res = false
		}
	}
	return res
}

func (sa *SemanticAnalysis) getMainFunction(funcs []*ast.Function) *ast.Function {
	if funcs == nil { return nil }
	for _, fn := range funcs {
		if fn.Name == "main" {
			return fn
		}
	}
	return nil
}

func (sa *SemanticAnalysis) getRealRetType(stmnts []ast.Statement, expectedType tp.Type) tp.Type {
	length := len(stmnts)
	var realRetType, tmpRetType tp.Type
	realRetType = nil
	tmpRetType = nil
	
	for i, _ := range stmnts {
		st := stmnts[length - i - 1]
		switch st.(type) {
		case *ast.ReturnSt:
			retSt := st.(*ast.ReturnSt)
			tmpRetType = sa.getExpressionType(retSt.Expr)
			if tmpRetType != expectedType {
				sa.semanticError(
					"Expected return type: " + expectedType.Literal() + 
					", but got: " +  tmpRetType.Literal(), retSt.Line)
				return tmpRetType
			} else {
				realRetType = tmpRetType
			}

			return realRetType
		case *ast.Block:
			return sa.getRealRetType(st.(*ast.Block).Stmnts, expectedType)
		case *ast.LoopSt:
			return sa.getRealRetType(st.(*ast.LoopSt).LoopBlock.Stmnts, expectedType)
		case *ast.ConditionalSt:
			var elseRetType tp.Type
			ifRetType := sa.getRealRetType(st.(*ast.ConditionalSt).IfBlock.Stmnts, expectedType)
			if st.(*ast.ConditionalSt).ElseBlock == nil {
				elseRetType = tp.VoidType
			} else {
				elseRetType = sa.getRealRetType(st.(*ast.ConditionalSt).ElseBlock.Stmnts, expectedType)
			}

			if (ifRetType != tp.VoidType && elseRetType == tp.VoidType) ||
			   (ifRetType == tp.VoidType && elseRetType != tp.VoidType) {
				if realRetType == nil {
					sa.semanticError(
						"All control-flow paths should have a return statement", 
						st.(*ast.ConditionalSt).Line)
					return tp.VoidType
				}
			}

			if ifRetType != expectedType { return ifRetType }
			if elseRetType != expectedType { return elseRetType }
			realRetType = expectedType
		default:
			continue
		}
	}

	if realRetType == nil {
		return tp.VoidType
	} else {
		return realRetType
	}
}

func (sa *SemanticAnalysis) checkFunction(fn *ast.Function) bool {
	var realRetTyp tp.Type
	fnEntry := sa.lookupAll(fn.Name).(*st.FuncEntry)
	sa.currScope = fnEntry.Scope
	defer func() { sa.currScope = sa.currScope.FinalizeScope() }()
	realRetTyp = sa.getRealRetType(fn.Stmnts, fn.ReturnTyp)

	if realRetTyp == nil {
		realRetTyp = tp.VoidType
	}
	res := true
	res = res && (realRetTyp == fn.ReturnTyp)
	res = res && sa.checkStatements(fn.Stmnts)

	return res
}

func (sa *SemanticAnalysis) checkStatements(stmts []ast.Statement) bool {
	if stmts == nil {
		return true
	}

	res := true
	for _, stmt := range stmts {
		if ok := sa.checkStatement(stmt); !ok {
			res = false
		}
	}
	return res
}

func (sa *SemanticAnalysis) checkStatement(stmt ast.Statement) bool {
	switch stmt.(type) {
	case *ast.Block:
		return sa.checkBlock(stmt.(*ast.Block))
	case *ast.AssignmentSt:
		return sa.checkAssignment(stmt.(*ast.AssignmentSt))
	case *ast.PrintSt:
		return sa.checkPrint(stmt.(*ast.PrintSt))
	case *ast.ConditionalSt:
		return sa.checkConditional(stmt.(*ast.ConditionalSt))
	case *ast.LoopSt:
		return sa.checkLoop(stmt.(*ast.LoopSt))
	case *ast.ReturnSt:
		return sa.checkReturn(stmt.(*ast.ReturnSt))
	case *ast.ReadSt:
		return sa.checkRead(stmt.(*ast.ReadSt))
	case *ast.InvocationSt:
		return sa.checkInvocation(stmt.(*ast.InvocationSt))
	default:
		return true
	}
}

func (sa *SemanticAnalysis) checkBlock(blc *ast.Block) bool {
	return sa.checkStatements(blc.Stmnts)
}

func (sa *SemanticAnalysis) checkAssignment(asmnt *ast.AssignmentSt) bool {
	LValueType := sa.getExpressionType(asmnt.LVal)
	if LValueType == nil {
		sa.semanticError(asmnt.LVal.String() + " not defined.", asmnt.Line)
		return false
	}
	ExprType := sa.getExpressionType(asmnt.Expr)
	if ExprType == nil {
		return false
	}
	if !LValueType.Equals(ExprType) {
		sa.semanticError(asmnt.LVal.String() + " = " + asmnt.Expr.String() + " assigment does not have same types on both sides.", asmnt.Line)
		return false
	}
	return true
}

func (sa *SemanticAnalysis) checkPrint(prt *ast.PrintSt) bool {
	entry := sa.lookupAll(prt.Iden)
	if entry == nil {
		sa.semanticError(prt.Iden + " is not defined.", prt.Line)
		return false
	}

	switch entry.(type) {
	case *st.VarEntry:
		if !entry.(*st.VarEntry).Typ.Equals(tp.IntType) {
			sa.semanticError("Only integer variable can be printed", prt.Line)
			return false
		}
	default:
		sa.semanticError(prt.Iden + " is not a variable", prt.Line)
		return false
	}
	return true
}

func (sa *SemanticAnalysis) checkConditional(condtn *ast.ConditionalSt) bool {
	if sa.getExpressionType(condtn.Condition) != tp.BoolType {
		sa.semanticError(condtn.Condition.String() + " - invalid boolean expression.", condtn.Line)
		return false
	}
	return true
}

func (sa *SemanticAnalysis) checkLoop(lp *ast.LoopSt) bool {
	if sa.getExpressionType(lp.Condition) != tp.BoolType {
		sa.semanticError(lp.Condition.String() + " - invalid boolean expression.", lp.Line)
		return false
	}
	return true
}

func (sa *SemanticAnalysis) checkReturn(rtrn *ast.ReturnSt) bool {
	return true
}

func (sa *SemanticAnalysis) checkRead(rd *ast.ReadSt) bool {
	entry := sa.lookupAll(rd.Iden)
	if entry == nil {
		sa.semanticError(rd.Iden + " is not defined.", rd.Line)
		return false
	}

	switch entry.(type) {
	case *st.VarEntry:
		if !entry.(*st.VarEntry).Typ.Equals(tp.IntType) {
			sa.semanticError("Only integer variable can be put in the variable", rd.Line)
			return false
		}
	default:
		sa.semanticError(rd.Iden + "is not a variable", rd.Line)
		return false
	}
	return true
}

func (sa *SemanticAnalysis) checkInvocation(invcn *ast.InvocationSt) bool {
	fnName := invcn.FuncName
	funcEntry := sa.lookupAll(fnName)
	if funcEntry == nil {
		sa.semanticError("Function " + fnName + " is not defined.", invcn.Line)
		// unknownFuncEntry := &st.FuncEntry{}
		// unknownFuncEntry.RetTyp = tp.UnknownType
		// sa.currScope.Parent.Insert(fnName, unknownFuncEntry)
		return false
	}

	if fnName == "delete" {
		return sa.checkDeleteInvocation(invcn)
	}

	if len(invcn.Arguments) != len(funcEntry.(*st.FuncEntry).ParaNames) {
		sa.semanticError("The number of arguments does not match the function definition.", invcn.Line)
		return false
	}

	for i := range invcn.Arguments {
		callTyp := sa.getExpressionType(invcn.Arguments[i])
		pName := funcEntry.(*st.FuncEntry).ParaNames[i]
		realEntry := sa.lookupAll(pName)
		
		switch callTyp.(type) {
		case *tp.StructType:
			if !callTyp.(*tp.StructType).Equals(realEntry.(*tp.StructType)) {
				sa.semanticError("The type of argument" + callTyp.Literal() + " does not match the definition.", invcn.Line)
				return false
			}
		default:
			if callTyp != realEntry.(*st.VarEntry).Typ {
				sa.semanticError("The type of argument" + callTyp.Literal() + " does not match the definition.", invcn.Line)
				return false
			}
		}
	}
	
	return true
}

func (sa *SemanticAnalysis) checkDeleteInvocation(invcn *ast.InvocationSt) bool {
	if invcn.Arguments == nil || len(invcn.Arguments) != 1 { 
		sa.semanticError("delete built-in function must have exactly one argument", invcn.Line)
		return false 
	}

	argStr := invcn.Arguments[0].String()
	entry := sa.lookupAll(argStr)
	switch entry.(type) {
	case *st.VarEntry:
		typ := entry.(*st.VarEntry).Typ
		switch typ.(type) {
		case *tp.StructType:
			return true
		default:
			sa.semanticError("Built-in function delete's argument must be an struct type", invcn.Line)
			return false 
		}
	default:
		sa.semanticError("Built-in function delete's argument must be an variable", invcn.Line)
		return false 
	}
}


// ================
// Get Type of Ast Node

func (sa *SemanticAnalysis) getExpressionType(expr ast.Expression) tp.Type {
	switch expr := expr.(type) {
	case *ast.BinOpExpr:
		return sa.getBinOpExprType(expr)
	case *ast.UnOpExpr:
		return sa.getUnOpExprType(expr)
	case *ast.InvocationExpr:
		return sa.getInvocationExprType(expr)
	case *ast.IdenLiteral:
		return sa.getIdenLiteralType(expr)
	case *ast.IntLiteral:
		return sa.getIntLiteralType(expr)
	case *ast.BoolLiteral:
		return sa.getBoolLiteralType(expr)
	case *ast.NilLiteral:
		return sa.getNilLiteralType(expr)
	default:
		return nil // Not reachable
	}
}

func (sa *SemanticAnalysis) getBinOpExprType(binexpr *ast.BinOpExpr) tp.Type {
	leftType := sa.getExpressionType(binexpr.Left)
	if leftType == tp.UnknownType {
		return tp.UnknownType
	}

	switch binexpr.Operator {
	case ast.EQ, ast.NE:
		if !leftType.Equals(sa.getExpressionType(binexpr.Right)) {
			sa.semanticError(binexpr.String() + " - types do not match", binexpr.Line)
			return tp.UnknownType
		}
		return tp.BoolType
	case ast.ADD, ast.MULT, ast.SUB, ast.DIV, ast.GE, ast.GT, ast.LE, ast.LT:
		var res tp.Type
		switch binexpr.Operator {
		case ast.ADD, ast.MULT, ast.SUB, ast.DIV:
			res = tp.IntType
		default:
			res = tp.BoolType
		}

		if !leftType.Equals(tp.IntType) {
			sa.semanticError(binexpr.String() + " - the left of " + ast.OpString(binexpr.Operator) + " must be an integer type.", binexpr.Line)
			return res
		}

		rightType := sa.getExpressionType(binexpr.Right)
		if !rightType.Equals(tp.IntType) {
			sa.semanticError(binexpr.String() + " - the right of " + ast.OpString(binexpr.Operator) + " must be an integer type.", binexpr.Line)
			return res
		}

		return res
		
	case ast.AND, ast.OR:
		if !leftType.Equals(tp.BoolType) {
			sa.semanticError(binexpr.String() + " - the left of " + ast.OpString(binexpr.Operator) + " must be a bool type.", binexpr.Line)
			return tp.BoolType
		}
		rightType := sa.getExpressionType(binexpr.Right)
		if !rightType.Equals(tp.BoolType) {
			fmt.Println(rightType.Literal())
			sa.semanticError(binexpr.String() + " - the right of " + ast.OpString(binexpr.Operator) + " must be a bool type.", binexpr.Line)
			return tp.BoolType
		}
		return tp.BoolType
	case ast.ACCESS:
		switch leftType.(type) {
		case *tp.StructType:
			realTyp := leftType.(*tp.StructType)
			rightExp := binexpr.Right
			switch rightExp.(type) {
			case *ast.IdenLiteral:
				t, ok := realTyp.Fields[rightExp.String()]
				if !ok {
					sa.semanticError(rightExp.String() + " must be a field of this struct type.", binexpr.Line)
				} else {
					return t
				}
			default:
				sa.semanticError(rightExp.String() + " must be a field of this struct type.", binexpr.Line)
			}
		default:
			return tp.UnknownType
		}

		return tp.UnknownType
	default:
		return tp.UnknownType
	}
}

func (sa *SemanticAnalysis) getUnOpExprType(unexpr *ast.UnOpExpr) tp.Type {
	leftExpType := sa.getExpressionType(unexpr.Left)
	switch unexpr.Operator {
	case ast.NOT:
		if !leftExpType.Equals(tp.BoolType) {
			sa.semanticError(unexpr.String() + " - Only bool type can follow '!' operator", unexpr.Line)
		} else {
			return tp.BoolType
		}
	case ast.SUB:
		if !leftExpType.Equals(tp.IntType) {
			sa.semanticError(unexpr.String() + " - Only integer type can follow '-' operator", unexpr.Line)
		} else {
			return tp.IntType
		}
	}
	return tp.UnknownType
}

func (sa *SemanticAnalysis) getNewInvocationExprType(inv *ast.InvocationExpr) tp.Type {
	if inv.Arguments == nil || len(inv.Arguments) != 1 {
		sa.semanticError("Built-in function new must have exactly one argument", inv.Line)
	}
	typName := inv.Arguments[0].String()
	completeStructEntry := sa.lookupAll(typName)
	if completeStructEntry == nil {
		return tp.UnknownType
	} else {
		entry := completeStructEntry.(*st.StructEntry)
		return entry.Typ
	}
}

func (sa *SemanticAnalysis) getInvocationExprType(inv *ast.InvocationExpr) tp.Type {
	fname := inv.Func.String()
	if fname == "new" {
		return sa.getNewInvocationExprType(inv)
	}
	entry := sa.lookupAll(fname)
	if entry == nil {
		sa.semanticError("Function " + fname + " is not defined.", inv.Line)
		// unknownFuncEntry := &st.FuncEntry{}
		// unknownFuncEntry.RetTyp = tp.UnknownType
		// sa.currScope.Parent.Insert(fname, unknownFuncEntry)
		return tp.UnknownType
	} else {
		fentry := entry.(*st.FuncEntry)
		return fentry.RetTyp
	}
}

func (sa *SemanticAnalysis) getIdenLiteralType(id *ast.IdenLiteral) tp.Type {
	entry := sa.lookupAll(id.Value)
	if entry == nil {
		sa.semanticError(id.String() + " is not defined.", id.Token.Line)
		sa.insert(id.Value, &st.VarEntry{Typ: tp.UnknownType})
		return tp.UnknownType
	}

	return entry.(*st.VarEntry).Typ
}

func (sa *SemanticAnalysis) getIntLiteralType(i *ast.IntLiteral) tp.Type {
	return tp.IntType
}

func (sa *SemanticAnalysis) getBoolLiteralType(bl *ast.BoolLiteral) tp.Type {
	return tp.BoolType
}

func (sa *SemanticAnalysis) getNilLiteralType(nl *ast.NilLiteral) tp.Type {
	return tp.VoidType
}

// Error Handling
func (sa *SemanticAnalysis) semanticError(msg string, errLine int) {
	fmt.Printf("%s:%d: semantic error: %s\n", sa.filePath, errLine, msg)
}
