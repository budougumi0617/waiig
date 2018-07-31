package ast

import (
	"testing"

	"github.com/budougumi0617/waiig/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	want := "let myVar = anotherVar;"
	if program.String() != want {
		t.Errorf("want = %q\nbut got=%q", want, program.String())
	}
}
