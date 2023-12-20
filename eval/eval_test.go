package eval_test

import (
	"calc/eval"
	"calc/lexer"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "834224",
			expected: 834224,
		},
		{
			input:    "-56432",
			expected: -56432,
		},
		{
			input:    "(-1)",
			expected: -1,
		},
		{
			input:    "(1)",
			expected: 1,
		},
		{
			input:    "1 +    2",
			expected: 3,
		},
		{
			input:    "7*6",
			expected: 42,
		},
		{
			input:    "7/6",
			expected: 1,
		},
		{
			input:    "7-6",
			expected: 1,
		},
		{
			input:    "1+(2*3)",
			expected: 7,
		},
		{
			input:    "(1+2)*3",
			expected: 9,
		},
		{
			input:    "(1+2)*3-(1+2)",
			expected: 6,
		},
		{
			input:    "(1+(6+(-4)))*3-(1+2)",
			expected: 6,
		},
		{
			input:    "(1+(6+(-4)))*3-(((-9999999)+10000000)+2)",
			expected: 6,
		},
	}

	for _, tt := range tests {
		lex := lexer.NewLexer(tt.input)

		res := eval.Evaluate(lex)

		if res != tt.expected {
			t.Fatalf("expected %d got %d", tt.expected, res)
		}
	}
}
