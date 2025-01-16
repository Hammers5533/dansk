package lexer

import (
	"unicode"

	"github.com/Hammers5533/dklang/token"
)

func (l *Lexer) currentRune() rune {
	if l.currentPosition < len(l.input) {
		return l.input[l.currentPosition]
	} else {
		return 0
	}
}

func (l *Lexer) nextRune() rune {
	if l.currentPosition < len(l.input)-1 {
		return l.input[l.currentPosition+1]
	} else {
		return 0
	}
}

func (l *Lexer) readString() string {
	position := l.currentPosition
	for l.currentRune() != '"' {
		l.advance()
	}
	return string(l.input[position:l.currentPosition])
}

func (l *Lexer) readDigit() (string, token.TokenType) {
	position := l.currentPosition
	for unicode.IsDigit(l.currentRune()) {
		l.advance()
	}
	if l.currentRune() == ',' {
		l.advance()
		for unicode.IsDigit(l.currentRune()) {
			l.advance()
		}
		return string(l.input[position:l.currentPosition]), token.FLOAT
	} else {
		return string(l.input[position:l.currentPosition]), token.INTEGER
	}
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.currentPosition
	for unicode.IsLetter(l.currentRune()) || unicode.IsDigit(l.currentRune()) || l.currentRune() == '_' {
		l.advance()
	}
	return string(l.input[startPosition:l.currentPosition])
}

func (l *Lexer) skipWhitespace() {
	for l.currentRune() == ' ' || l.currentRune() == '\t' || l.currentRune() == '\n' || l.currentRune() == '\r' {
		l.advance()
	}
}
