package ast

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
