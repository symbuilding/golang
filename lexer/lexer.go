package lexer

import (
	"calc/token"
)

type Lexer struct {
	input        string
	ch           string
	readPosition int
	position     int
}

func NewLexer(input string) *Lexer {
	lex := &Lexer{
		input:        input,
		readPosition: 0,
		position:     0,
		ch:           "\x00",
	}

	lex.readChar()

	return lex
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespaces()

	switch lex.ch {
	case "+":
		tok = token.Token{Literal: lex.ch, Type: token.PLUS}
	case "-":
		tok = token.Token{Literal: lex.ch, Type: token.MINUS}
	case "*":
		tok = token.Token{Literal: lex.ch, Type: token.ASTERIC}
	case "/":
		tok = token.Token{Literal: lex.ch, Type: token.SLASH}
	case "%":
		tok = token.Token{Literal: lex.ch, Type: token.MOD}

	case "(":
		tok = token.Token{Literal: lex.ch, Type: token.LPAREN}
	case ")":
		tok = token.Token{Literal: lex.ch, Type: token.RPAREN}

	case "\x00":
		tok = token.Token{Literal: "", Type: token.EOL}
	default:
		if isDigit(lex.ch) {
			tok.Literal = lex.readInt()
			tok.Type = token.INT
			return tok
		} else if isChar(lex.ch) {
			tok = lex.readFunction()
			return tok
		} else {
			tok = token.Token{Literal: lex.ch, Type: token.ILLEGAL}
		}
	}

	lex.readChar()

	return tok
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = "\x00"
	} else {
		lex.ch = string(lex.input[lex.readPosition])
	}

	lex.position = lex.readPosition

	lex.readPosition++
}

func (lex *Lexer) skipWhitespaces() {
	for lex.ch == " " || lex.ch == "\n" || lex.ch == "\r" || lex.ch == "\t" {
		lex.readChar()
	}
}

func (lex *Lexer) readInt() string {
	pos := lex.position

	for isDigit(lex.ch) {
		lex.readChar()
	}

	return lex.input[pos:lex.position]
}

func (lex *Lexer) readFunction() token.Token {
	pos := lex.position

	for isChar(lex.ch) {
		lex.readChar()
	}

	functionLiteral := lex.input[pos:lex.position]

	return token.Token{
		Literal: functionLiteral,
		Type:    lex.getFunctionType(functionLiteral),
	}
}

func isDigit(ch string) bool {
	return ch >= "0" && ch <= "9"
}

func isChar(ch string) bool {
	return ch >= "a" && ch <= "z"
}

func (lex *Lexer) getFunctionType(literal string) string {
	switch literal {
	case "sin":
		return token.SIN
	case "cos":
		return token.COS
	case "tan":
		return token.TAN
	default:
		return token.ILLEGAL
	}
}
