package parser

import (
	"github.com/Hammers5533/dklang/intepreter"
	"github.com/Hammers5533/dklang/token"
)

func parseManyStatements(p *Parser) intepreter.Statement {
	var statements = []intepreter.Statement{}
	for p.hasTokens() && p.currentTokenType() != token.RIGHTBRACE {
		stmt := parseStatement(p)
		if stmt != nil {
			statements = append(statements, stmt)
		}
	}
	return intepreter.Body{Body: statements}
}

func parseStatement(p *Parser) intepreter.Statement {
	statementFun, exists := statementLU[p.currentTokenType()]

	if exists {
		return statementFun(p)
	}
	return parseExpressionStatement(p)

}

func parseExpressionStatement(p *Parser) intepreter.ExpStatementWrapper {
	expression := parseExpression(p, defaultBP)
	p.expect(token.SEMICOLON)

	return intepreter.ExpStatementWrapper{
		Exp: expression,
	}
}

func parseAssignment(p *Parser) intepreter.Statement {
	p.expect(token.LET)

	name := p.expect(token.IDENTIFIER)
	p.expect(token.ASSIGN)

	expression := parseExpression(p, defaultBP)
	p.expect(token.SEMICOLON)
	return intepreter.AssignStatement{
		Name:  name.Literal,
		Value: expression,
	}
}

func parseFunctionDeclaration(p *Parser) intepreter.Statement {
	p.expect(token.FUNCTION)

	functionParameters := make([]string, 0)

	functionName := p.expect(token.IDENTIFIER).Literal
	p.expect(token.LEFTPARENTHESIS)
	for p.hasTokens() && p.currentTokenType() != token.RIGHTPARENTHESIS {
		parameterName := p.expect(token.IDENTIFIER).Literal
		functionParameters = append(functionParameters, parameterName)
		if p.currentTokenType() != token.RIGHTPARENTHESIS {
			p.expect(token.DOT)
		}
	}
	p.expect(token.RIGHTPARENTHESIS)
	p.expect(token.LEFTBRACE)

	body := parseManyStatements(p)

	p.expect(token.RIGHTBRACE)

	return intepreter.ExpStatementWrapper{Exp: intepreter.ValueExpWrapper{Value: intepreter.FuncDef{Name: functionName, Parameters: functionParameters, Body: body}}}
}

func parseReturn(p *Parser) intepreter.Statement {
	p.expect(token.RETURN)

	returnExp := parseExpression(p, defaultBP)
	p.expect(token.SEMICOLON)
	return intepreter.ReturnStatement{Exp: returnExp}
}

func parseIf(p *Parser) intepreter.Statement {
	p.expect(token.IF)
	p.expect(token.LEFTPARENTHESIS)

	conditionExp := parseExpression(p, defaultBP)

	p.expect(token.RIGHTPARENTHESIS)
	p.expect(token.LEFTBRACE)

	IfBody := parseManyStatements(p)

	p.expect(token.RIGHTBRACE)

	var ElseBody intepreter.Statement = intepreter.Body{}
	if p.currentTokenType() == token.ELSE {
		p.expect(token.ELSE)
		p.expect(token.LEFTBRACE)

		ElseBody = parseManyStatements(p)

		p.expect(token.RIGHTBRACE)
	}

	return intepreter.IfStatement{
		Condition: conditionExp,
		IfBody:    IfBody,
		ElseBody:  ElseBody,
	}
}

func parseWhile(p *Parser) intepreter.Statement {
	p.expect(token.WHILE)
	p.expect(token.LEFTPARENTHESIS)

	conditionExp := parseExpression(p, defaultBP)

	p.expect(token.RIGHTPARENTHESIS)
	p.expect(token.LEFTBRACE)

	Body := parseManyStatements(p)

	p.expect(token.RIGHTBRACE)

	return intepreter.WhileStatement{
		Condition: conditionExp,
		Body:      Body,
	}
}
