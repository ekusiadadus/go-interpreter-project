// src/lexer/lexer_test.go
package lexer

import (
	"testing"

	"github.com/ekusiadadus/go-interpreter-project/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	expected := []token.Token{
		{Kind: token.Assign, Literal: "="},
		{Kind: token.Plus, Literal: "+"},
		{Kind: token.LParen, Literal: "("},
		{Kind: token.RParen, Literal: ")"},
		{Kind: token.LBrace, Literal: "{"},
		{Kind: token.RBrace, Literal: "}"},
		{Kind: token.Comma, Literal: ","},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Eof, Literal: ""},
	}

	lexer := New(input)

	for idx, expectedToken := range expected {
		recvToken := lexer.NextToken()

		if recvToken.Kind != expectedToken.Kind {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				idx, expectedToken.Kind, recvToken.Kind)
		}

		if recvToken.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
				idx, expectedToken.Literal, recvToken.Literal)
		}
	}
}

func TestNextToken1(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`

	expected := []token.Token{
		{Kind: token.Let, Literal: "let"},
		{Kind: token.Ident, Literal: "five"},
		{Kind: token.Assign, Literal: "="},
		{Kind: token.Int, Literal: "5"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Let, Literal: "let"},
		{Kind: token.Ident, Literal: "ten"},
		{Kind: token.Assign, Literal: "="},
		{Kind: token.Int, Literal: "10"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Let, Literal: "let"},
		{Kind: token.Ident, Literal: "add"},
		{Kind: token.Assign, Literal: "="},
		{Kind: token.Function, Literal: "fn"},
		{Kind: token.LParen, Literal: "("},
		{Kind: token.Ident, Literal: "x"},
		{Kind: token.Comma, Literal: ","},
		{Kind: token.Ident, Literal: "y"},
		{Kind: token.RParen, Literal: ")"},
		{Kind: token.LBrace, Literal: "{"},
		{Kind: token.Ident, Literal: "x"},
		{Kind: token.Plus, Literal: "+"},
		{Kind: token.Ident, Literal: "y"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.RBrace, Literal: "}"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Let, Literal: "let"},
		{Kind: token.Ident, Literal: "result"},
		{Kind: token.Assign, Literal: "="},
		{Kind: token.Ident, Literal: "add"},
		{Kind: token.LParen, Literal: "("},
		{Kind: token.Ident, Literal: "five"},
		{Kind: token.Comma, Literal: ","},
		{Kind: token.Ident, Literal: "ten"},
		{Kind: token.RParen, Literal: ")"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Eof, Literal: ""},
	}

	lexer := New(input)

	for idx, expectedToken := range expected {
		recvToken := lexer.NextToken()

		if recvToken.Kind != expectedToken.Kind {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				idx, expectedToken.Kind, recvToken.Kind)
		}

		if recvToken.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
				idx, expectedToken.Literal, recvToken.Literal)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `!-/*5;
5 < 10 > 5;`

	expected := []token.Token{
		{Kind: token.Bang, Literal: "!"},
		{Kind: token.Minus, Literal: "-"},
		{Kind: token.Slash, Literal: "/"},
		{Kind: token.Asterisk, Literal: "*"},
		{Kind: token.Int, Literal: "5"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Int, Literal: "5"},
		{Kind: token.Lt, Literal: "<"},
		{Kind: token.Int, Literal: "10"},
		{Kind: token.Gt, Literal: ">"},
		{Kind: token.Int, Literal: "5"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Eof, Literal: ""},
	}

	lexer := New(input)

	for idx, expectedToken := range expected {
		recvToken := lexer.NextToken()

		if recvToken.Kind != expectedToken.Kind {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				idx, expectedToken.Kind, recvToken.Kind)
		}

		if recvToken.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
				idx, expectedToken.Literal, recvToken.Literal)
		}
	}
}

func TestNextToken3(t *testing.T) {
	input := `if (5 < 10) {
	return true;
} else {
	return false;
}
`

	expected := []token.Token{
		{Kind: token.If, Literal: "if"},
		{Kind: token.LParen, Literal: "("},
		{Kind: token.Int, Literal: "5"},
		{Kind: token.Lt, Literal: "<"},
		{Kind: token.Int, Literal: "10"},
		{Kind: token.RParen, Literal: ")"},
		{Kind: token.LBrace, Literal: "{"},
		{Kind: token.Return, Literal: "return"},
		{Kind: token.True, Literal: "true"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.RBrace, Literal: "}"},
		{Kind: token.Else, Literal: "else"},
		{Kind: token.LBrace, Literal: "{"},
		{Kind: token.Return, Literal: "return"},
		{Kind: token.False, Literal: "false"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.RBrace, Literal: "}"},
		{Kind: token.Eof, Literal: ""},
	}

	lexer := New(input)

	for idx, expectedToken := range expected {
		recvToken := lexer.NextToken()

		if recvToken.Kind != expectedToken.Kind {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				idx, expectedToken.Kind, recvToken.Kind)
		}

		if recvToken.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
				idx, expectedToken.Literal, recvToken.Literal)
		}
	}
}

func TestNextToken4(t *testing.T) {
	input := `10 == 10;
10 != 9;
`
	expected := []token.Token{
		{Kind: token.Int, Literal: "10"},
		{Kind: token.Eq, Literal: "=="},
		{Kind: token.Int, Literal: "10"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Int, Literal: "10"},
		{Kind: token.NotEq, Literal: "!="},
		{Kind: token.Int, Literal: "9"},
		{Kind: token.Semicolon, Literal: ";"},
		{Kind: token.Eof, Literal: ""},
	}

	lexer := New(input)

	for idx, expectedToken := range expected {
		recvToken := lexer.NextToken()

		if recvToken.Kind != expectedToken.Kind {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				idx, expectedToken.Kind, recvToken.Kind)
		}

		if recvToken.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
				idx, expectedToken.Literal, recvToken.Literal)
		}
	}
}
