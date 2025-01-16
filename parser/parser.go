package parser

import (
	"github.com/Hammers5533/dklang/intepreter"
	"github.com/Hammers5533/dklang/token"
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
	program.Body = intepreter.Body{Body: statements}

	return program
}
