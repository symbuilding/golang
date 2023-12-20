package eval

import (
	"calc/lexer"
	"calc/token"
	"fmt"
	"math"
	"strconv"
)

const (
	_ = iota
	LOWEST
	FUNCS
	SUM
	PRODUCT
	PREFIX
	GROUP
)

type Eval struct {
	prefixFns   map[string]func() int
	infixFns    map[string]func(int) int
	precedences map[string]int

	lex *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func Evaluate(lex *lexer.Lexer) int {
	eval := &Eval{
		lex: lex,
	}

	eval.prefixFns = map[string]func() int{
		token.INT:    eval.evalIntegers,
		token.LPAREN: eval.evalGroupedExpression,

		token.MINUS: eval.evalPrefixMinus,

		token.SIN: eval.evalTrigsFuncs,
		token.COS: eval.evalTrigsFuncs,
		token.TAN: eval.evalTrigsFuncs,
	}

	eval.infixFns = map[string]func(int) int{
		token.PLUS:    eval.evalInfixExpression,
		token.MINUS:   eval.evalInfixExpression,
		token.SLASH:   eval.evalInfixExpression,
		token.ASTERIC: eval.evalInfixExpression,
	}

	eval.precedences = map[string]int{
		token.PLUS:  SUM,
		token.MINUS: SUM,

		token.SLASH:   PRODUCT,
		token.ASTERIC: PRODUCT,

		token.LPAREN: GROUP,
	}

	eval.nextToken()
	eval.nextToken()

	return eval.evalExpression(LOWEST)
}

func (eval *Eval) evalExpression(precedence int) int {
	prefix, ok := eval.prefixFns[eval.curToken.Type]

	if !ok {
		fmt.Println("You fked up :/")
		return -1
	}

	left := prefix()

	for !eval.peekTokenIs(token.EOL) && precedence < eval.peekPrecedence() {
		infix, ok := eval.infixFns[eval.peekToken.Type]

		if !ok {
			return left
		}

		eval.nextToken()

		left = infix(left)
	}

	return left
}

func (eval *Eval) nextToken() {
	eval.curToken = eval.peekToken

	eval.peekToken = eval.lex.NextToken()
}

func (eval *Eval) evalIntegers() int {
	val, _ := strconv.Atoi(eval.curToken.Literal)

	return val
}

func (eval *Eval) evalGroupedExpression() int {
	eval.nextToken()

	val := eval.evalExpression(LOWEST)

	if !eval.peekTokenIs(token.RPAREN) {
		fmt.Println("Forgot to close the parens?")
		return -1
	}

	eval.nextToken()

	return val
}

func (eval *Eval) evalPrefixMinus() int {
	eval.nextToken()

	if !eval.curTokenIs(token.INT) {
		fmt.Printf("Incorrect opearand for prefix minus, got %s\n", eval.curToken.Literal)
		return -1
	}

	return -1 * eval.evalExpression(PREFIX)
}

func (eval *Eval) evalTrigsFuncs() int {
	funcType := eval.curToken.Type

	eval.nextToken()

	angle := eval.evalExpression(GROUP)

	radian := float64(angle) * (math.Pi / 180)

	var val int

	switch funcType {
	case token.SIN:
		val = int(math.Sin(radian))
	case token.COS:
		val = int(math.Cos(radian))
	case token.TAN:
		val = int(math.Tan(radian))
	default:
		val = -1
	}

	return val
}

func (eval *Eval) evalInfixExpression(left int) int {
	operator := eval.curToken.Literal

	precendence := eval.curPrecedence()

	eval.nextToken()

	right := eval.evalExpression(precendence)

	var val int

	switch operator {
	case "+":
		val = left + right
	case "-":
		val = left - right
	case "*":
		val = left * right
	case "/":
		val = left / right
	}

	return val
}

func (eval *Eval) peekTokenIs(tokType string) bool {
	return eval.peekToken.Type == tokType
}

func (eval *Eval) curTokenIs(tokType string) bool {
	return eval.curToken.Type == tokType
}

func (eval *Eval) peekPrecedence() int {
	precedence, ok := eval.precedences[eval.peekToken.Type]

	if !ok {
		return LOWEST
	}

	return precedence
}

func (eval *Eval) curPrecedence() int {
	precedence, ok := eval.precedences[eval.curToken.Type]

	if !ok {
		return LOWEST
	}

	return precedence
}
