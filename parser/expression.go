package parser

import (
	"dklang/intepreter"
	"fmt"
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

		leftExpression = ledFun(p, leftExpression, bp)
	}

	return leftExpression
}

func parseBinaryExpression(p *Parser, leftExpression intepreter.Exp, bp bindingPower) intepreter.Exp {
	operatorToken := p.advance()

	// Parse right side
	rightExpression := parseExpression(p, defaultBP)

	return intepreter.BinaryExpression{Left: leftExpression, Operator: operatorToken, Right: rightExpression}
}
