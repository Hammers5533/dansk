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

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	LT  = "<"
	GT  = ">"
	LEQ = "<="
	GEQ = ">="
	AND = "AND"
	OR  = "OR"
	NOT = "NOT"

	COMMA     = ","
	SEMICOLON = ";"

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
	TRY      = "PRØV"
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
