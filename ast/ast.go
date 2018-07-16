package ast

import (
	"github.com/budougumi0617/waiig/token"
)

// Node is base of AST.
type Node interface {
	TokenLiteral() string
}

// Statement is base statement.
type Statement interface {
	Node
	statementNode()
}

// Expression is base expression.
type Expression interface {
	Node
	expressionNode()
}

// Program imprements Node interface.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns literal.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is LET statement.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns literal.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns literal.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
