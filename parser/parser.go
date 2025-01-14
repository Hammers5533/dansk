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

	return p
}

func (p *Parser) ParseProgram() *intepreter.Program {
	program := &intepreter.Program{}
	program.Statements = []intepreter.Statement{}

	for p.hasTokens() {
		stmt := parseStatement(p)
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}
	return program
}
