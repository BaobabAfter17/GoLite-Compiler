package scanner

import (
	"strings"
	ct "proj/token"
)

type Scanner struct {
	idx int                                    // Next index to read in the string
	curLine int                                // Current line number
	rawInput string                            // The input string
	inputLength int                            // The length of string
	arithmeticOpMap map[byte] ct.TokenType     // The mapping from arithmetic char to token type
	keywordMap map[string] ct.TokenType        // The mapping from keyword to token type
	numCmpOpMap map[string] ct.TokenType       // The mapping from num compare operation to token type
	boolCmpOpMap map[string] ct.TokenType      // The mapping from bool compare operation to token type
	singleSpecChar map[string] ct.TokenType    // The set of single special character
}

// New returns an initialized Scanner
func New(input string) *Scanner {
	scanner := &Scanner{}
	scanner.rawInput = input
	scanner.curLine = 1
	scanner.idx = 0
	scanner.inputLength = len(input)

	scanner.keywordMap = map[string]ct.TokenType{
		"package": ct.PACKAGE, "import": ct.IMPORT, "fmt": ct.FMT, "type": ct.TYPE, 
		"struct": ct.STRUCT, "int": ct.INT, "bool": ct.BOOL, "var": ct.VAR, "func": ct.FUNC, 
		"Scan": ct.SCAN, "Print": ct.PRINT, "Println": ct.PRINTLN, "if": ct.IF, "else": ct.ELSE, 
		"for": ct.FOR, "return": ct.RETURN, "true": ct.TRUE, "false": ct.FALSE, "nil": ct.NIL,
	}
	scanner.numCmpOpMap = map[string] ct.TokenType{
		"<=": ct.LESSOREQUAL, "<": ct.LESSTHAN, ">=": ct.GREATEROREQUAL, ">": ct.GREATERTHAN,
		"==": ct.EQUAL, "!=": ct.NOTEQUAL, "=": ct.ASSIGN, "!": ct.EXCLAMATION,
	}
	scanner.boolCmpOpMap = map[string] ct.TokenType{
		"|": ct.OR, "||": ct.OR, "&": ct.AMPERSAND, "&&": ct.AND,
	}
	scanner.singleSpecChar = map[string] ct.TokenType{
		",": ct.COMMA, ";": ct.SEMICOLON, "\"": ct.QUOTATION, "{": ct.LEFTCURLY,
		"}": ct.RIGHTCURLY, "(": ct.LEFTPARENTHESIS, ")": ct.RIGHTPARENTHESIS,
		".": ct.DOT, "+": ct.PLUS, "-": ct.HYPHEN, "*": ct.ASTEROID, "/":ct.SLASH,
	}
	return scanner
}

// NextToken returns the next token in the format of Token struct.
// This function uses the FA method.
func (l *Scanner) NextToken() ct.Token {
	var lexeme strings.Builder
	var tk string
	var nextChar byte
	var _type ct.TokenType
	var exist bool

	// ======================================================
	//   START STATE:
	//       The state prepares the empty string builder and 
	//       diretly go to the __S0__ state.
	// ======================================================
	__START__: 
	lexeme.Reset()
	goto __S0__

	// ======================================================
	//   S0 STATE:
	//       The state differentiates the possible token
	//       types such as integer, identifier (or keyword),
	//       and operators.
	// ======================================================
	__S0__:
	if l.idx >= l.inputLength {
		return ct.Token{ct.EOF, "EOF", l.curLine}
	}

	nextChar = l.rawInput[l.IncIdx()]
	if nextChar == ' ' || nextChar == '\n' || nextChar == '\t' {
		if nextChar == '\n' {
			l.curLine ++
		}
		goto __START__
	}

	lexeme.WriteByte(nextChar)

	if IsDigit(nextChar) {
		goto __S2__
	}

	if IsLetter(nextChar) {
		goto __S3__
	}

	// Special character
	goto __S4__

	// ======================================================
	//   S2 STATE:
	//       The state receives a new character of an integer
	//       and output an integer token.
	// ======================================================
	__S2__:
    if l.idx == l.inputLength {
        return ct.Token{ct.NUMBER, lexeme.String(), l.curLine}
    }
	nextChar = l.rawInput[l.IncIdx()]
	if !IsDigit(nextChar) {
		l.idx --
		return ct.Token{ct.NUMBER, lexeme.String(), l.curLine}
	} else if IsDigit(nextChar) {
		lexeme.WriteByte(nextChar)
		goto __S2__
	}

	// ======================================================
	//   S3 STATE:
	//       The state receives a new character of an identifier
	//       or a keyword.
	// ======================================================
	__S3__:
	if l.idx == l.inputLength {
		goto __S3_END__
	}
	nextChar = l.rawInput[l.IncIdx()]
	
	if IsLetter(nextChar) || IsDigit(nextChar) {
		lexeme.WriteByte(nextChar)
		goto __S3__
	} else {
		l.idx --
		goto __S3_END__
	}

	// ======================================================
	//   S3 END STATE:
	//       The state determines the specific type.
	// ======================================================
	__S3_END__:
		tk = lexeme.String()
		_type, exist = l.keywordMap[tk]
		if exist {
			return ct.Token{_type, tk, l.curLine}
		} else {
			return ct.Token{ct.ID, tk, l.curLine}
		}

	// ======================================================
	//   S4 STATE:
	//       The state determines the type of splitting
	//       characters such as ';', '=', '+', '-', '*', '/'
	// ======================================================
	__S4__:
	tk = lexeme.String()
	_type, exist = l.numCmpOpMap[tk]
	if exist {
		if l.idx < l.inputLength {
			nextChar = l.rawInput[l.IncIdx()]
			if nextChar == '=' {
				return ct.Token{l.numCmpOpMap[tk + "="], tk + "=", l.curLine}
			} else {
				l.idx --
				return ct.Token{_type, tk, l.curLine}
			}
		} else {
			return ct.Token{_type, tk, l.curLine}
		}
	}

	_type, exist = l.boolCmpOpMap[tk]
	if exist {
		if l.idx < l.inputLength {
			nextChar = l.rawInput[l.IncIdx()]
			if tk + string(nextChar) == "||" {
				return ct.Token{ct.OR, "||", l.curLine}
			} else if tk + string(nextChar) == "&&" {
				return ct.Token{ct.AND, "&&", l.curLine}
			} else if tk == "|" {
				goto __SEXIT__
			} else if tk == "&" {
				l.idx --
				return ct.Token{ct.AMPERSAND, tk, l.curLine}
			}
		} else {
			if tk == "|" {
				goto __SEXIT__
			}
			return ct.Token{_type, tk, l.curLine}
		}
	}

	_type, exist = l.singleSpecChar[tk]
	if exist {
		if _type == ct.SLASH && l.idx < l.inputLength && l.rawInput[l.IncIdx()] == '/' {
			goto __S5__ // comment state
		}
		return ct.Token{_type, tk, l.curLine}
	}

	// No special character match
	goto __SEXIT__

	// ======================================================
	//   S5 STATE:
	//       Handling Comments (starts with //, ends with \n)
	// ======================================================

	__S5__:
	for l.idx < l.inputLength {
		nextChar = l.rawInput[l.IncIdx()]
		if nextChar == '\n' {
			goto __START__
		}
	}

	// ======================================================
	//   SEXIT STATE:
	//       The state outputs illegal tokens
	// ======================================================
	__SEXIT__:
	return ct.Token{ct.ILLEGAL, lexeme.String(), l.curLine}

}

// IncIdx returns l.idx and then increment it
func (l *Scanner) IncIdx() int {
	res := l.idx
	l.idx ++
	return res
}

// IsLetter returns if c is a letter
func IsLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// IsDigit returns if c is a digit
func IsDigit(c byte) bool {
	return (c >= '0' && c <= '9')
}