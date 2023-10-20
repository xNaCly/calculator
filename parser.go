package main

import "log"

// Grammar:
// expression ::= additive-expression
// additive-expression ::= multiplicative-expression ( ( '+' | '-' ) multiplicative-expression ) *
// multiplicative-expression ::= primary ( ( '*' | '/' ) primary ) *
// primary ::= NUMBER

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
	r := make([]Node, 0)
	for p.peek().Type != TOKEN_EOF {
		r = append(r, p.expression())
	}
	return r
}

func (p *Parser) expression() Node {
	return p.additiveExpression()
}

func (p *Parser) additiveExpression() Node {
	lhs := p.multiplicativeExpression()
	op := p.peek()
	if p.matchT(op, TOKEN_PLUS, TOKEN_MINUS) {
		p.advance()
		if op.Type == TOKEN_PLUS {
			return &Addition{
				token: op,
				left:  lhs,
				right: p.multiplicativeExpression(),
			}
		} else if op.Type == TOKEN_MINUS {
			return &Subtraction{
				token: op,
				left:  lhs,
				right: p.multiplicativeExpression(),
			}
		} else {
			log.Panicf("Unexpected %q", TOKEN_LOOKUP[p.peek().Type])
		}
	}
	return lhs
}

func (p *Parser) multiplicativeExpression() Node {
	lhs := p.primary()
	op := p.peek()
	if p.matchT(op, TOKEN_ASTERISK, TOKEN_SLASH) {
		p.advance()
		if op.Type == TOKEN_ASTERISK {
			return &Multiplication{
				token: op,
				left:  lhs,
				right: p.primary(),
			}
		} else if op.Type == TOKEN_SLASH {
			return &Division{
				token: op,
				left:  lhs,
				right: p.primary(),
			}
		} else {
			log.Panicf("Unexpected %q", TOKEN_LOOKUP[p.peek().Type])
		}
	}
	return lhs
}

func (p *Parser) primary() Node {
	if !p.matchT(p.peek(), TOKEN_NUMBER) {
		log.Panicf("Expected %q, got %q", TOKEN_LOOKUP[TOKEN_NUMBER], TOKEN_LOOKUP[p.peek().Type])
	}
	op := p.peek()
	p.advance()
	return &Number{token: op}
}

func (p *Parser) matchT(t Token, tokenTypes ...int) bool {
	for _, tokenType := range tokenTypes {
		if t.Type == tokenType {
			return true
		}
	}
	return false
}

func (p *Parser) peekEquals(tokenType int) bool {
	return p.peek().Type == tokenType
}

// get current rune
func (p *Parser) peek() Token {
	return p.token[p.pos]
}

// advance cursor to the next token in the input
func (p *Parser) advance() {
	if p.peek().Type != TOKEN_EOF {
		p.pos++
	}
}
