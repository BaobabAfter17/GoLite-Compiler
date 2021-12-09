package token

type TokenType string

const (
	ILLEGAL          = "ILLEGAL"
	EOF              = "EOF"
	PACKAGE          = "PACKAGE"          // package
	ID               = "ID"               // Identifiers
	SEMICOLON        = "SEMICOLON"        // ;
	IMPORT           = "IMPORT"           // import
	QUOTATION        = "QUOTATION"        // "
	FMT              = "FMT"              // fmt
	TYPE             = "TYPE"             // type
	STRUCT           = "STRUCT"           // struct
	LEFTCURLY        = "LEFTCURLY"        // {
	RIGHTCURLY       = "RIGHTCURLY"       // }
	INT              = "INT"              // int
	BOOL             = "BOOL"             // bool
	ASTEROID         = "ASTEROID"         // *
	VAR              = "VAR"              // var
	COMMA            = "COMMA"            // ,
	FUNC             = "FUNC"             // func
	LEFTPARENTHESIS  = "LEFTPARENTHESIS"  // (
	RIGHTPARENTHESIS = "RIGHTPARENTHESIS" // )
	ASSIGN           = "ASSIGN"           // =
	DOT              = "DOT"              // .
	SCAN             = "SCAN"             // Scan
	AMPERSAND        = "AMPERSAND"        // &
	PRINT            = "PRINT"            // Print
	PRINTLN          = "PRINTLN"          // Println
	IF               = "IF"               // if
	ELSE             = "ELSE"             // else
	FOR              = "FOR"              // for
	RETURN           = "RETURN"           // return
	OR               = "OR"               // ||
	AND              = "AND"              // &&
	EQUAL            = "EQUAL"            // ==
	NOTEQUAL         = "NOTEQUAL"         // !=
	GREATERTHAN      = "GREATERTHAN"      // >
	LESSTHAN         = "LESSTHAN"         // <
	GREATEROREQUAL   = "GREATEROREQUAL"   // >=
	LESSOREQUAL      = "LESSOREQUAL"      // <=
	PLUS             = "PLUS"             // +
	HYPHEN           = "HYPHEN"           // -
	SLASH            = "SLASH"            // /
	EXCLAMATION      = "EXCLAMATION"      // !
	NUMBER           = "NUMBER"           // Numbers
	TRUE             = "TRUE"             // true
	FALSE            = "FALSE"            // false
	NIL              = "NIL"              // nil
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}
