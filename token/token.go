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
	// ASSIGN means "="
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)
