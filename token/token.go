package token

// TokenType defines token type in Monkey.
type TokenType string

// Token is a token object.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL means unknown keyword.
	ILLEGAL = "ILLEGAL"
	// EOF means "End Of File".
	EOF = "EOF"

	// Identifiers + literals

	// IDENT means identifiers.
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT means number.
	INT = "INT" // 1343456

	// Operators

	// ASSIGN means "=".
	ASSIGN = "="
	// PLUS means "+".
	PLUS = "+"
	// MINUS means "-".
	MINUS = "-"
	// BANG means "!".
	BANG = "!"
	// ASTERISK means "*".
	ASTERISK = "*"
	// SLASH    means "/".
	SLASH = "/"

	// LT means "<".
	LT = "<"
	// GT means ">".
	GT = ">"

	// EQ means "==".
	EQ = "=="
	// NOTEQ means "!=".
	NOTEQ = "!="

	// Delimiters

	// COMMA means ",".
	COMMA = ","
	// SEMICOLON means ";".
	SEMICOLON = ";"

	// LPAREN means "(".
	LPAREN = "("
	// RPAREN means ")".
	RPAREN = ")"
	// LBRACE means "{".
	LBRACE = "{"
	// RBRACE means "}".
	RBRACE = "}"

	// Keywords

	// FUNCTION means "FUNCTION".
	FUNCTION = "FUNCTION"
	// LET means "LET".
	LET = "LET"
	// TRUE means "TRUE".
	TRUE = "TRUE"
	// FALSE means "FALSE".
	FALSE = "FALSE"
	// IF means "IF".
	IF = "IF"
	// ELSE  means "ELSE".
	ELSE = "ELSE"
	// RETURN  means "RETURN".
	RETURN = "RETURN"
)
