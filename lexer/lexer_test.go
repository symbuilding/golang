package lexer_test

import (
	"calc/lexer"
	"calc/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "+*-   ( / *)  %  123+69   sin(0)cos(0)tan(0)cot(1+2)"
	lex := lexer.NewLexer(input)

	expected_tokens := []token.Token{
		{Literal: "+", Type: token.PLUS},
		{Literal: "*", Type: token.ASTERIC},
		{Literal: "-", Type: token.MINUS},

		{Literal: "(", Type: token.LPAREN},
		{Literal: "/", Type: token.SLASH},
		{Literal: "*", Type: token.ASTERIC},
		{Literal: ")", Type: token.RPAREN},

		{Literal: "%", Type: token.ILLEGAL},

		{Literal: "123", Type: token.INT},
		{Literal: "+", Type: token.PLUS},
		{Literal: "69", Type: token.INT},

		{Literal: "sin", Type: token.SIN},
		{Literal: "(", Type: token.LPAREN},
		{Literal: "0", Type: token.INT},
		{Literal: ")", Type: token.RPAREN},

		{Literal: "cos", Type: token.COS},
		{Literal: "(", Type: token.LPAREN},
		{Literal: "0", Type: token.INT},
		{Literal: ")", Type: token.RPAREN},

		{Literal: "tan", Type: token.TAN},
		{Literal: "(", Type: token.LPAREN},
		{Literal: "0", Type: token.INT},
		{Literal: ")", Type: token.RPAREN},

		{Literal: "cot", Type: token.COT},
		{Literal: "(", Type: token.LPAREN},
		{Literal: "1", Type: token.INT},
		{Literal: "+", Type: token.PLUS},
		{Literal: "2", Type: token.INT},
		{Literal: ")", Type: token.RPAREN},

		{Literal: "", Type: token.EOL},
	}

	for _, expected_tok := range expected_tokens {
		tok := lex.NextToken()
		if tok.Literal != expected_tok.Literal {
			t.Fatalf("Unexpected token literal. Expected %s got %s.", expected_tok.Literal, tok.Literal)
		}

		if tok.Type != expected_tok.Type {
			t.Fatalf("Unexpected token type. Expected %s got %s.", expected_tok.Type, tok.Type)
		}
	}
}
