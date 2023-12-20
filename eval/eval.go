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
	HIGHEST
)

type Eval struct {
	prefixFns   map[string]func() float64
	infixFns    map[string]func(float64) float64
	precedences map[string]int

	lex *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func Evaluate(lex *lexer.Lexer) float64 {
	eval := &Eval{
		lex: lex,
	}

	eval.prefixFns = map[string]func() float64{
		token.NUM:    eval.evalIntegers,
		token.LPAREN: eval.evalGroupedExpression,

		token.MINUS: eval.evalPrefixMinus,

		token.SIN: eval.evalFuncs,
		token.COS: eval.evalFuncs,
		token.TAN: eval.evalFuncs,

		token.SQRT: eval.evalFuncs,
	}

	eval.infixFns = map[string]func(float64) float64{
		token.PLUS:    eval.evalInfixExpression,
		token.MINUS:   eval.evalInfixExpression,
		token.SLASH:   eval.evalInfixExpression,
		token.ASTERIC: eval.evalInfixExpression,

		token.MOD: eval.evalInfixExpression,

		token.EXPONENT: eval.evalInfixExpression,

		token.EULER: eval.evalInfixExpression,
	}

	eval.precedences = map[string]int{
		token.PLUS:  SUM,
		token.MINUS: SUM,

		token.SLASH:   PRODUCT,
		token.ASTERIC: PRODUCT,
		token.MOD:     PRODUCT,

		token.EXPONENT: HIGHEST,
		token.EULER:    HIGHEST,

		token.LPAREN: GROUP,

		token.SIN: FUNCS,
		token.COS: FUNCS,
		token.TAN: FUNCS,

		token.SQRT: FUNCS,
	}

	eval.nextToken()
	eval.nextToken()

	return eval.evalExpression(LOWEST)
}

func (eval *Eval) evalExpression(precedence int) float64 {
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

func (eval *Eval) evalIntegers() float64 {
	val, _ := strconv.ParseFloat(eval.curToken.Literal, 64)

	return val
}

func (eval *Eval) evalGroupedExpression() float64 {
	eval.nextToken()

	val := eval.evalExpression(LOWEST)

	if !eval.peekTokenIs(token.RPAREN) {
		fmt.Println("Forgot to close the parens?")
		return -1
	}

	eval.nextToken()

	return val
}

func (eval *Eval) evalPrefixMinus() float64 {
	eval.nextToken()

	if !eval.curTokenIs(token.NUM) {
		fmt.Printf("Incorrect opearand for prefix minus, got %s\n", eval.curToken.Literal)
		return -1
	}

	return -1.00000 * eval.evalExpression(PREFIX)
}

func (eval *Eval) evalFuncs() float64 {
	funcType := eval.curToken.Type

	eval.nextToken()

	arg := eval.evalExpression(GROUP)

	var val float64

	switch funcType {
	case token.SIN, token.COS, token.TAN:
		val = eval.evalTrigFuncs(funcType, arg)
	case token.SQRT:
		val = math.Sqrt(arg)
	default:
		val = -1.00000
	}

	return val
}

func (eval *Eval) evalTrigFuncs(funcType string, angle float64) float64 {
	radian := angle * (math.Pi / 180)

	var val float64

	switch funcType {
	case token.SIN:
		val = math.Sin(radian)
	case token.COS:
		val = math.Cos(radian)
	case token.TAN:
		val = math.Tan(radian)
	}

	return val
}

func (eval *Eval) evalInfixExpression(left float64) float64 {
	operator := eval.curToken.Literal

	precendence := eval.curPrecedence()

	eval.nextToken()

	right := eval.evalExpression(precendence)

	var val float64

	switch operator {
	case "+":
		val = left + right
	case "-":
		val = left - right
	case "*":
		val = left * right
	case "/":
		val = left / right
	case "%":
		r := int(right)
		l := int(left)
		val = float64(l % r)

	case "**":
		val = math.Pow(left, right)

	case "e":
		if left != 1 {
			fmt.Printf("Non one power 10 exponent, got %f", left)
		}
		val = math.Pow10(int(right))
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
