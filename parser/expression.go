package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Hammers5533/dklang/intepreter"
	"github.com/Hammers5533/dklang/token"
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

func parseFunctionCall(p *Parser, leftExpression intepreter.Exp, bp bindingPower) intepreter.Exp {
	p.expect(token.LEFTPARENTHESIS)

	expressions := make([]intepreter.Exp, 0)

	for p.hasTokens() && p.currentTokenType() != token.RIGHTPARENTHESIS {
		exp := parseExpression(p, defaultBP)
		expressions = append(expressions, exp)

		if p.currentTokenType() != token.RIGHTPARENTHESIS {
			p.expect(token.DOT)
		}
	}
	p.expect(token.RIGHTPARENTHESIS)

	return intepreter.FuncCall{Name: leftExpression, Parameters: expressions}
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

func parseAssignmentExpression(p *Parser, left intepreter.Exp, bp bindingPower) intepreter.Exp {
	switch p.advance().Type {
	case token.ASSIGN:
		right := parseExpression(p, bp)
		return intepreter.AssignExpression{
			Name:  left,
			Value: right,
		}
	case token.PLUSASSIGN:
		right := parseExpression(p, bp)
		return intepreter.AssignExpression{
			Name: left,
			Value: intepreter.BinaryExpression{
				Left:     left,
				Operator: token.Token{Type: token.PLUS, Literal: "+"},
				Right:    right,
			},
		}
	case token.MINUSASSIGN:
		right := parseExpression(p, bp)
		return intepreter.AssignExpression{
			Name: left,
			Value: intepreter.BinaryExpression{
				Left:     left,
				Operator: token.Token{Type: token.MINUS, Literal: "-"},
				Right:    right,
			},
		}
	default:
		panic("Not a valid assignment token")
	}
}

func parseArrayExpression(p *Parser) intepreter.Exp {
	p.expect(token.LEFTBRACKET)

	var values []intepreter.Exp = []intepreter.Exp{}

	for p.currentToken().Type != token.RIGHTBRACKET && p.hasTokens() {
		exp := parseExpression(p, defaultBP)
		values = append(values, exp)

		if p.currentToken().Type != token.RIGHTBRACKET && p.hasTokens() {
			p.expect(token.DOT)
		}
	}
	p.expect(token.RIGHTBRACKET)
	return intepreter.ValueExpWrapper{Value: intepreter.List{Value: values}}
}

func parseMemberExpression(p *Parser, left intepreter.Exp, bp bindingPower) intepreter.Exp {
	p.expect(token.LEFTBRACKET)

	index := parseExpression(p, defaultBP)

	p.expect(token.RIGHTBRACKET)

	return intepreter.MemberExpression{
		Member: left,
		Index:  index,
	}
}

func parsePrefixExpression(p *Parser) intepreter.Exp {
	operatorToken := p.advance()
	expression := parseExpression(p, unary)

	return intepreter.PrefixExpression{
		Operator: operatorToken,
		Right:    expression,
	}
}
