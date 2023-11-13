package main

import (
	"log"
)

// Grammar:
// expression ::= term
// term       ::= factor ( ( '+' | '-' ) factor ) *
// factor     ::= unary ( ( '*' | '/' ) unary ) *
// unary      ::= ('-') unary | primary
// primary    ::= NUMBER | '(' expression ')'

type Parser struct {
	token []Token
	pos   int
}

func NewParser(token []Token) *Parser {
	p := &Parser{
		token: token,
		pos:   0,
	}
	return p
}

func (p *Parser) Parse() []Node {
	o := make([]Node, 0)
	for !p.atEnd() {
		o = append(o, p.expression())
	}
	return o
}

func (p *Parser) expression() Node {
	return p.term()
}

func (p *Parser) term() Node {
	lhs := p.factor()

	for p.match(TOKEN_MINUS, TOKEN_PLUS) {
		op := p.previous()
		rhs := p.factor()
		lhs = &Binary{
			token: op,
			left:  lhs,
			right: rhs,
		}
	}

	return lhs
}

func (p *Parser) factor() Node {
	lhs := p.unary()

	for p.match(TOKEN_SLASH, TOKEN_ASTERISK) {
		op := p.previous()
		rhs := p.unary()
		lhs = &Binary{
			token: op,
			left:  lhs,
			right: rhs,
		}
	}

	return lhs
}

func (p *Parser) unary() Node {
	if p.match(TOKEN_MINUS) {
		op := p.previous()
		right := p.unary()
		return &Unary{token: op, right: right}
	}

	return p.primary()
}

func (p *Parser) primary() Node {
	if p.match(TOKEN_NUMBER) {
		op := p.previous()
		return &Number{token: op}
	} else if p.match(TOKEN_BRACE_LEFT) {
		node := p.expression()
		p.consume(TOKEN_BRACE_RIGHT, "Expected '('")
		return node
	}

	panic("Expected expression")
}

func (p *Parser) match(tokenTypes ...int) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(tokenType int, error string) {
	if p.check(tokenType) {
		p.advance()
		return
	}
	log.Panicf("Wanted %q, got %q: %s", TOKEN_LOOKUP[p.peek().Type], TOKEN_LOOKUP[tokenType], error)
}

func (p *Parser) check(tokenType int) bool {
	if p.atEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.atEnd() {
		p.pos++
	}
	return p.previous()
}

func (p *Parser) atEnd() bool {
	return p.peek().Type == TOKEN_EOF
}

func (p *Parser) peek() Token {
	return p.token[p.pos]
}

func (p *Parser) previous() Token {
	return p.token[p.pos-1]
}
