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
