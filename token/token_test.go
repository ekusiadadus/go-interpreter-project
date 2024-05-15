package token

import "testing"

func TestTokenTypeString(t *testing.T) {
	tests := []struct {
		kind     TokenType
		expected string
	}{
		{Illegal, "Illegal"},
		{Eof, "Eof"},
		{Ident, "Ident"},
		{Int, "Int"},
		{Assign, "="},
		{Eq, "=="},
		{NotEq, "!="},
		{Plus, "+"},
		{Minus, "-"},
		{Bang, "!"},
		{Asterisk, "*"},
		{Slash, "/"},
		{Lt, "<"},
		{Gt, ">"},
		{Comma, ","},
		{Semicolon, ";"},
		{LParen, "("},
		{RParen, ")"},
		{LBrace, "{"},
		{RBrace, "}"},
		{Function, "fn"},
		{Let, "let"},
		{If, "if"},
		{Else, "else"},
		{True, "true"},
		{False, "false"},
		{Return, "return"},
		{TokenType(100), "TokenType(100)"},
	}

	for _, tt := range tests {
		if got := tt.kind.String(); got != tt.expected {
			t.Errorf("TokenType.String() = %v, want %v", got, tt.expected)
		}
	}
}

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		identifier string
		expected   TokenType
	}{
		{"fn", Function},
		{"let", Let},
		{"if", If},
		{"else", Else},
		{"true", True},
		{"false", False},
		{"return", Return},
		{"foo", Ident},
		{"bar", Ident},
	}

	for _, tt := range tests {
		if got := LookupIdent(tt.identifier); got != tt.expected {
			t.Errorf("LookupIdent(%v) = %v, want %v", tt.identifier, got, tt.expected)
		}
	}
}
