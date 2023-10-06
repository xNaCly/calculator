// Expression   -> Expression + Term | Expression - Term | Term
// Term         -> Term * Factor | Term / Factor | Factor
// Factor       -> Number | ( Expression )
// Number       -> [0-9e_.]
package main

type Parser struct {
	token []Token
	pos   int
}

func NewParser(token []Token) *Parser {
	p := &Parser{
		token: token,
		pos:   0,
	}
	p.advance()
	return p
}

func (p *Parser) Parse() []Node {
	r := make([]Node, 0)
	for p.peek().Type != EOF {
		p.advance()
	}
	return r
}

// get current rune
func (p *Parser) peek() Token {
	return p.token[p.pos]
}

// advance cursor to the next token in the input
func (p *Parser) advance() {
	if p.peek().Type != EOF {
		p.pos++
	}
}
