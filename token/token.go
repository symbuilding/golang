package token

const (
	EOL     = "Eol"
	ILLEGAL = "Illegal"
	NULL = "NULL"

    NUM = "Int"

	PLUS    = "Plus"
	MINUS   = "Minus"
	ASTERIC = "Asteric"
	SLASH   = "Slash"
    MOD = "Mod"

    EXPONENT = "Exponent"

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
