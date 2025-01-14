package parser

import (
	"dklang/token"
	"fmt"
)

func (p *Parser) currentToken() token.Token {
	return p.tokens[p.position]
}

func (p *Parser) currentTokenType() token.TokenType {
	return p.tokens[p.position].Type
}

func (p *Parser) advance() token.Token {
	token := p.currentToken()
	p.position++
	return token
}

func (p *Parser) previousToken() token.Token {
	return p.tokens[p.position-1]
}

func (p *Parser) nextToken() token.Token {
	return p.tokens[p.position+1]
}

func (p *Parser) hasTokens() bool {
	return p.position < len(p.tokens) && p.currentTokenType() != token.EOF
}

func (p *Parser) expectError(expectedType token.TokenType, err any) token.Token {
	token := p.currentToken()
	tokenType := token.Type

	if tokenType != expectedType {
		if err == nil {
			err = fmt.Sprintf("Expected TokenType %s but instead got %s", expectedType, tokenType)
		}
		panic(err)
	}

	return p.advance()
}

func (p *Parser) expect(expectedType token.TokenType) token.Token {
	return p.expectError(expectedType, nil)
}
