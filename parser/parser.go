package parser

import (
	"dklang/intepreter"
	"dklang/token"
)

type Parser struct {
	tokens   []token.Token
	position int
}

func New(tokens []token.Token) *Parser {
	p := &Parser{tokens: tokens}
	createTokenLookups()
	return p
}

func ParseProgram(tokens []token.Token) *intepreter.Program {
	p := New(tokens)

	program := &intepreter.Program{}
	statements := []intepreter.Statement{}

	for p.hasTokens() {
		stmt := parseStatement(p)
		if stmt != nil {
			statements = append(statements, stmt)
		}
	}
	program.Body = statements

	return program
}
