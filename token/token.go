package token

// TokenType defines token type in Monkey.
type TokenType string

// Token is a token object
type Token struct {
	Type    TokenType
	Literal string
}
