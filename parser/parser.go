package parser

import (
	"github.com/budougumi0617/waiig/ast"
	"github.com/budougumi0617/waiig/lexer"
	"github.com/budougumi0617/waiig/token"
)

// Parser parses statements by Lexer rule.
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New returns Parser object.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram returns Program object with parsed result.
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
