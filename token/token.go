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
)

type Token struct {
	Type    string
	Literal string
}
