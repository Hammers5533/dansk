package lexer

import (
	"dklang/token"
	"unicode"
)

type Lexer struct {
	input           []rune
	currentPosition int
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input), currentPosition: 0}
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	Rune := l.currentRune()
	switch Rune {
	case '+':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.PLUSASSIGN, Literal: "+="}
			l.advance()
		} else {
			tok = newToken(token.PLUS, Rune)
		}
	case '-':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.MINUSASSIGN, Literal: "+="}
			l.advance()
		} else {
			tok = newToken(token.MINUS, Rune)
		}
	case '*':
		tok = newToken(token.MULTIPLY, Rune)
	case '=':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.EQ, Literal: "=="}
			l.advance()
		} else {
			tok = newToken(token.ASSIGN, Rune)
		}
	case '!':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.NEQ, Literal: "!="}
			l.advance()
		} else {
			tok = newToken(token.NOT, Rune)
		}
	case '<':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.LEQ, Literal: "<="}
			l.advance()
		} else {
			tok = newToken(token.LT, Rune)
		}
	case '>':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.GEQ, Literal: ">="}
			l.advance()
		} else {
			tok = newToken(token.GT, Rune)
		}
	case ';':
		tok = newToken(token.SEMICOLON, Rune)
	case '"':
		l.advance()
		tok.Literal = l.readString()
		tok.Type = token.STRING
	case '.':
		tok = newToken(token.DOT, Rune)
	case ',':
		tok = newToken(token.COMMA, Rune)
	case '(':
		tok = newToken(token.LEFTPARENTHESIS, Rune)
	case ')':
		tok = newToken(token.RIGHTPARENTHESIS, Rune)
	case '{':
		tok = newToken(token.LEFTBRACE, Rune)
	case '}':
		tok = newToken(token.RIGHTBRACE, Rune)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(Rune) || Rune == '_' {
			tok.Literal = l.readIdentifier()
			tok.Type = token.CheckIdentifier(tok.Literal)
			return tok
		} else if unicode.IsDigit(Rune) {
			tok.Literal, tok.Type = l.readDigit()
			return tok
		} else {
			tok = newToken(token.INVALID, Rune)
		}

	}

	l.advance()
	return tok

}

func (l *Lexer) advance() rune {
	Rune := l.currentRune()
	l.currentPosition++
	return Rune
}

// Takes a type and character
func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func Tokenize(s string) []token.Token {
	l := New(s)
	var tokens []token.Token
	tok := token.Token{}
	for tok.Type != token.EOF {
		tok = l.NextToken()
		tokens = append(tokens, tok)
	}
	return tokens
}
