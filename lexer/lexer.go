package lexer

import (
	"dklang/token"
	"fmt"
	"unicode"
)

type Lexer struct {
	input           []rune
	currentPosition int
	readingPosition int
	ch              rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.CheckIdentifier(tok.Literal)
			fmt.Println(tok)
			return tok
		} else if unicode.IsDigit(l.ch) {
			tok.Literal = l.readDigit()
			tok.Type = token.INTEGER

		} else {
			tok = newToken(token.INVALID, l.ch)
		}

	}

	l.readChar()
	fmt.Println(tok)
	return tok

}

func (l *Lexer) readDigit() string {
	position := l.currentPosition
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.currentPosition])
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.currentPosition])
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readingPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readingPosition])
	}
	l.currentPosition = l.readingPosition
	l.readingPosition += 1
}

// Takes a type and chara
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
