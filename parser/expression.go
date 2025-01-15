package parser

import (
	"dklang/intepreter"
	"dklang/token"
	"fmt"
	"strconv"
	"strings"
)

func parseExpression(p *Parser, bp bindingPower) intepreter.Exp {
	tokenType := p.currentTokenType()
	nudFun, exists := nudLU[tokenType]

	// First token is always NUD
	if !exists {
		panic(fmt.Sprintf("Expected NUD handler for token %s\n", tokenType))
	}

	leftExpression := nudFun(p)
	// check if binding power is increasing otherwise end reccursion
	for bpLU[p.currentTokenType()] > bp {

		// Expecting a LED token
		tokenType = p.currentTokenType()
		ledFun, exists := ledLU[tokenType]

		if !exists {
			panic(fmt.Sprintf("Expected LED handler for token %s\n", tokenType))
		}

		leftExpression = ledFun(p, leftExpression, bpLU[p.currentTokenType()])
	}

	return leftExpression
}

func parseBinaryExpression(p *Parser, leftExpression intepreter.Exp, bp bindingPower) intepreter.Exp {
	operatorToken := p.advance()

	// Parse right side
	rightExpression := parseExpression(p, bp)

	return intepreter.BinaryExpression{Left: leftExpression, Operator: operatorToken, Right: rightExpression}
}

func parseGroupExpression(p *Parser) intepreter.Exp {
	p.expect(token.LEFTPARENTHESIS)
	expression := parseExpression(p, defaultBP)
	p.expect(token.RIGHTPARENTHESIS)
	return expression
}

func parseValue(p *Parser) intepreter.Exp {
	valueToken := p.advance()

	switch valueToken.Type {
	case token.INTEGER:
		value, _ := strconv.Atoi(valueToken.Literal)
		return intepreter.ValueExpWrapper{Value: intepreter.Integer{Value: value}}
	case token.FLOAT:
		floatVal, _ := strconv.ParseFloat(strings.Replace(valueToken.Literal, ",", ".", 1), 64)
		return intepreter.ValueExpWrapper{Value: intepreter.Float{Value: floatVal}}
	case token.IDENTIFIER:
		return intepreter.ValueExpWrapper{Value: intepreter.Variable{Value: valueToken.Literal}}
	case token.STRING:
		return intepreter.ValueExpWrapper{Value: intepreter.String{Value: valueToken.Literal}}
	case token.TRUE:
		return intepreter.ValueExpWrapper{Value: intepreter.Bool{Value: true}}
	case token.FALSE:
		return intepreter.ValueExpWrapper{Value: intepreter.Bool{Value: false}}
	default:
		err := fmt.Sprintf("Extected token type %s to be a litteral\n", valueToken.Type)
		panic(err)
	}
}
