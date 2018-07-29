package ast

import (
	"bytes"

	"github.com/budougumi0617/waiig/token"
)

// Node is base of AST.
type Node interface {
	TokenLiteral() string
	String() string
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

// String resurn literal string.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String resurn literal string.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement is LET statement.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns literal.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String resurn literal string.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier is identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns literal.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String resurn literal string.
func (i *Identifier) String() string { return i.Value }

// ExpressionStatement is expression statement.
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns literal.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String resurn literal string.
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
