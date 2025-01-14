package parser

import (
	"dklang/intepreter"
	"dklang/token"
)

type bindingPower int

const (
	defaultBP bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type statementHandler func(p *Parser) intepreter.Statement
type nudHandler func(p *Parser) intepreter.Exp
type ledHandler func(p *Parser, left intepreter.Exp, bp bindingPower) intepreter.Exp

type statementLookup map[token.TokenType]statementHandler
type nudLookup map[token.TokenType]nudHandler
type ledLookup map[token.TokenType]ledHandler
type bpLookup map[token.TokenType]bindingPower

var bpLU = bpLookup{}
var nudLU = nudLookup{}
var ledLU = ledLookup{}
var statementLU = statementLookup{}

func led(tokenType token.TokenType, bp bindingPower, ledFun ledHandler) {
	bpLU[tokenType] = bp
	ledLU[tokenType] = ledFun
}

func nud(tokenType token.TokenType, nudFun nudHandler) {
	bpLU[tokenType] = primary
	nudLU[tokenType] = nudFun
}

func statement(tokenType token.TokenType, statementFun statementHandler) {
	bpLU[tokenType] = defaultBP
	statementLU[tokenType] = statementFun
}

func createTokenLookups() {

	// Binary Operators
	led(token.PLUS, additive, parseBinaryExpression)
	led(token.MINUS, additive, parseBinaryExpression)
}
