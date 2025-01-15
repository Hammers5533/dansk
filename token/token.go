package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	INVALID = "INVALID"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"
	FLOAT      = "FLOAT"
	STRING     = "STRING"

	ASSIGN      = "="
	PLUSASSIGN  = "+="
	MINUSASSIGN = "-="

	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	MODULUS  = "%"

	LT  = "<"
	GT  = ">"
	LEQ = "<="
	GEQ = ">="
	NEQ = "!="
	EQ  = "=="

	AND = "AND"
	OR  = "OR"

	NOT = "NOT"

	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."

	LEFTPARENTHESIS  = "("
	RIGHTPARENTHESIS = ")"
	LEFTBRACE        = "{"
	RIGHTBRACE       = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NONE     = "NONE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	IN       = "IN"
	WHILE    = "WHILE"
	FOR      = "FOR"
	TRY      = "TRY"
	EXCEPT   = "EXCEPT"
	FINALLY  = "FINALLY"
	THROW    = "THROW"
)

var keywords = map[string]TokenType{
	"funktion":   FUNCTION,
	"lad":        LET,
	"sandt":      TRUE,
	"falsk":      FALSE,
	"hvis":       IF,
	"ellers":     ELSE,
	"tilbagegiv": RETURN,
	"i":          IN,
	"imens":      WHILE,
	"prøv":       TRY,
	"medmindre":  EXCEPT,
	"endligt":    FINALLY,
	"kast":       THROW,
	"og":         AND,
	"eller":      OR,
	"omvendt":    NOT,
}

func CheckIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return IDENTIFIER
}
