package main

import (
	"log"
	"strings"
)

// Grammar:
// expression ::= term
// term       ::= factor ( ( '+' | '-' ) factor ) *
// factor     ::= unary ( ( '*' | '/' ) unary ) *
// unary      ::= ('-') unary | primary
// primary    ::= NUMBER | '(' expression ')'

type Parser struct {
	token    []Token
	pos      int
	dbgDepth int
}

func NewParser(token []Token) *Parser {
	p := &Parser{
		token: token,
		pos:   0,
	}
	return p
}

func (p *Parser) debug(prefix string) {
	log.Println(strings.Repeat(" ", p.dbgDepth), prefix, "\n", strings.Repeat(" ", p.dbgDepth), " ", TOKEN_LOOKUP[p.peek().Type])
	p.dbgDepth++
}

func (p *Parser) Parse() []Node {
	p.debug("Parse")
	defer func() {
		p.dbgDepth--
	}()

	o := make([]Node, 0)
	for !p.atEnd() {
		o = append(o, p.expression())
	}
	return o
}

func (p *Parser) expression() Node {
	p.debug("expression")
	defer func() {
		p.dbgDepth--
	}()
	return p.term()
}

func (p *Parser) term() Node {
	p.debug("term")
	defer func() {
		p.dbgDepth--
	}()
	lhs := p.factor()

	for p.match(TOKEN_MINUS, TOKEN_PLUS) {
		op := p.previous()
		rhs := p.factor()
		if op.Type == TOKEN_MINUS {
			lhs = &Subtraction{
				token: op,
				left:  lhs,
				right: rhs,
			}
		} else if op.Type == TOKEN_PLUS {
			lhs = &Addition{
				token: op,
				left:  lhs,
				right: rhs,
			}
		}
	}

	return lhs
}

func (p *Parser) factor() Node {
	p.debug("factor")
	defer func() {
		p.dbgDepth--
	}()
	lhs := p.unary()

	for p.match(TOKEN_MINUS, TOKEN_PLUS) {
		op := p.previous()
		rhs := p.unary()
		if op.Type == TOKEN_SLASH {
			lhs = &Division{
				token: op,
				left:  lhs,
				right: rhs,
			}
		} else if op.Type == TOKEN_ASTERISK {
			lhs = &Multiplication{
				token: op,
				left:  lhs,
				right: rhs,
			}
		}
	}

	return lhs
}

func (p *Parser) unary() Node {
	p.debug("unary")
	defer func() {
		p.dbgDepth--
	}()

	if p.match(TOKEN_MINUS) {
		op := p.previous()
		right := p.unary()
		return &Unary{token: op, right: right}
	}

	return p.primary()
}

func (p *Parser) primary() Node {
	p.debug("primary")
	defer func() {
		p.dbgDepth--
	}()
	if p.match(TOKEN_NUMBER) {
		op := p.previous()
		return &Number{token: op}
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
