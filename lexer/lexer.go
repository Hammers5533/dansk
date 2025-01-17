package lexer

import (
	"unicode"

	"github.com/Hammers5533/dklang/token"
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
	RuneStartPosition := l.currentPosition
	switch Rune {
	case '+':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.PLUSASSIGN, Literal: "+=", Position: RuneStartPosition}
			l.advance()
		} else {
			tok = newToken(token.PLUS, Rune, RuneStartPosition)
		}
	case '-':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.MINUSASSIGN, Literal: "+="}
			l.advance()
		} else {
			tok = newToken(token.MINUS, Rune, RuneStartPosition)
		}
	case '*':
		tok = newToken(token.MULTIPLY, Rune, RuneStartPosition)
	case '=':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.EQ, Literal: "=="}
			l.advance()
		} else {
			tok = newToken(token.ASSIGN, Rune, RuneStartPosition)
		}
	case '!':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.NEQ, Literal: "!="}
			l.advance()
		} else {
			tok = newToken(token.NOT, Rune, RuneStartPosition)
		}
	case '<':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.LEQ, Literal: "<="}
			l.advance()
		} else {
			tok = newToken(token.LT, Rune, RuneStartPosition)
		}
	case '>':
		if l.nextRune() == '=' {
			tok = token.Token{Type: token.GEQ, Literal: ">="}
			l.advance()
		} else {
			tok = newToken(token.GT, Rune, RuneStartPosition)
		}
	case ';':
		tok = newToken(token.SEMICOLON, Rune, RuneStartPosition)
	case '"':
		l.advance()
		tok.Literal = l.readString()
		tok.Type = token.STRING
	case '.':
		tok = newToken(token.DOT, Rune, RuneStartPosition)
	case ',':
		tok = newToken(token.COMMA, Rune, RuneStartPosition)
	case '(':
		tok = newToken(token.LEFTPARENTHESIS, Rune, RuneStartPosition)
	case ')':
		tok = newToken(token.RIGHTPARENTHESIS, Rune, RuneStartPosition)
	case '{':
		tok = newToken(token.LEFTBRACE, Rune, RuneStartPosition)
	case '}':
		tok = newToken(token.RIGHTBRACE, Rune, RuneStartPosition)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Position = RuneStartPosition
	default:
		if unicode.IsLetter(Rune) || Rune == '_' {
			tok.Literal = l.readIdentifier()
			tok.Type = token.CheckIdentifier(tok.Literal)
			tok.Position = RuneStartPosition
			return tok
		} else if unicode.IsDigit(Rune) {
			tok.Literal, tok.Type = l.readDigit()
			tok.Position = RuneStartPosition
			return tok
		} else {
			tok = newToken(token.INVALID, Rune, RuneStartPosition)
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
func newToken(tokenType token.TokenType, ch rune, start int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Position: start}
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
