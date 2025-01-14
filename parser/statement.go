package parser

import (
	"dklang/intepreter"
	"dklang/token"
)

func parseStatement(p *Parser) intepreter.Statement {
	statementFun, exists := statementLU[p.currentTokenType()]

	if exists {
		return statementFun(p)
	}

	return parseStatement(p)

}

func parseExpressionStatement(p *Parser) intepreter.ExpStatementWrapper {
	expression := parseExpression(p, defaultBP)
	p.expect(token.SEMICOLON)

	return intepreter.ExpStatementWrapper{
		Exp: expression,
	}
}
