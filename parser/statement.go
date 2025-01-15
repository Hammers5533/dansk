package parser

import (
	"dklang/intepreter"
	"dklang/token"
	"fmt"
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
	assignmentType := p.advance().Type

	expression := parseExpression(p, defaultBP)
	p.expect(token.SEMICOLON)

	switch assignmentType {
	case token.ASSIGN:
		return intepreter.Assign{
			Name:  name.Literal,
			Value: expression,
		}
	case token.PLUSASSIGN:
		return intepreter.Assign{
			Name: name.Literal,
			Value: intepreter.BinaryExpression{
				Left:     intepreter.ValueExpWrapper{Value: intepreter.Variable{Value: name.Literal}},
				Operator: token.Token{Type: token.PLUS, Literal: "+"},
				Right:    expression,
			},
		}
	case token.MINUSASSIGN:
		return intepreter.Assign{
			Name: name.Literal,
			Value: intepreter.BinaryExpression{
				Left:     intepreter.ValueExpWrapper{Value: intepreter.Variable{Value: name.Literal}},
				Operator: token.Token{Type: token.MINUS, Literal: "-"},
				Right:    expression,
			},
		}
	default:
		err := fmt.Sprintf("Not a valid assignment token %s", assignmentType)
		panic(err)
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

	return intepreter.FuncDef{Name: functionName, Parameters: functionParameters, Body: body}

}
