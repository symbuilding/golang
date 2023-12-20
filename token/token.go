package token

const (
	EOL     = "Eol"
	ILLEGAL = "Illegal"
	NULL = "NULL"

    INT = "Int"

	PLUS    = "Plus"
	MINUS   = "Minus"
	ASTERIC = "Asteric"
	SLASH   = "Slash"

	RPAREN = "Rparen"
	LPAREN = "Lparen"

    SIN = "Sin"
    COS = "Cos"
    TAN = "Tan"
)

type Token struct {
	Type    string
	Literal string
}
