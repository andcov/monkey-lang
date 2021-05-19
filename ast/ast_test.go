package ast

import (
	"testing"

	"github.com/andcov/monkey-lang/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foobar"},
					Value: "foobar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "hello_world"},
					Value: "hello_world",
				},
			},
		},
	}

	if program.String() != "let foobar = hello_world;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
