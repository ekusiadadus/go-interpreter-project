package token

import "fmt"

type Token struct {
	Kind    TokenType
	Literal string
}

type TokenType int

const (
	Illegal TokenType = iota
	Eof
	Ident
	Int
	Assign
	Plus
	Minus
	Bang
	Asterisk
	Slash
	Lt
	Gt
	Eq
	NotEq
	Comma
	Semicolon
	LParen
	RParen
	LBrace
	RBrace
	Function
	Let
	If
	Else
	Return
	True
	False
)

func (tk TokenType) String() string {
	switch tk {
	case Illegal:
		return "Illegal"
	case Eof:
		return "Eof"
	case Ident:
		return "Ident"
	case Int:
		return "Int"
	case Assign:
		return "="
	case Eq:
		return "=="
	case NotEq:
		return "!="
	case Plus:
		return "+"
	case Minus:
		return "-"
	case Bang:
		return "!"
	case Asterisk:
		return "*"
	case Slash:
		return "/"
	case Lt:
		return "<"
	case Gt:
		return ">"
	case Comma:
		return ","
	case Semicolon:
		return ";"
	case LParen:
		return "("
	case RParen:
		return ")"
	case LBrace:
		return "{"
	case RBrace:
		return "}"
	case Function:
		return "fn"
	case Let:
		return "let"
	case If:
		return "if"
	case Else:
		return "else"
	case True:
		return "true"
	case False:
		return "false"
	case Return:
		return "return"
	default:
		return fmt.Sprintf("TokenType(%d)", tk)
	}
}

func LookupIdent(identifier string) TokenType {
	switch identifier {
	case "fn":
		return Function
	case "let":
		return Let
	case "if":
		return If
	case "else":
		return Else
	case "true":
		return True
	case "false":
		return False
	case "return":
		return Return
	default:
		return Ident
	}
}
