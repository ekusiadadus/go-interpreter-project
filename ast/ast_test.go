package ast

import (
	"testing"

	"github.com/ekusiadadus/go-interpreter-project/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Kind: token.Let, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Kind: token.Ident, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Kind: token.Ident, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestTokenLiteral(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Kind: token.Let, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Kind: token.Ident, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Kind: token.Ident, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.TokenLiteral() != "let" {
		t.Errorf("program.TokenLiteral() wrong. got=%q", program.TokenLiteral())
	}

	letStmt, ok := program.Statements[0].(*LetStatement)
	if !ok {
		t.Errorf("program.Statements[0] is not *ast.LetStatement. got=%T", program.Statements[0])
	}

	if letStmt.TokenLiteral() != "let" {
		t.Errorf("letStmt.TokenLiteral() wrong. got=%q", letStmt.TokenLiteral())
	}

	if letStmt.Name.TokenLiteral() != "myVar" {
		t.Errorf("letStmt.Name.TokenLiteral() wrong. got=%q", letStmt.Name.TokenLiteral())
	}

	if letStmt.Name.String() != "myVar" {
		t.Errorf("letStmt.Name.String() wrong. got=%q", letStmt.Name.String())
	}
}
