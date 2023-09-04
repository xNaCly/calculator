package main

type Parser struct {
	token []Token
	cur   Token
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

func (p *Parser) peek() Token {
	return p.token[p.pos]
}

func (p *Parser) advance() {
	if p.peek().Type != EOF {
		p.pos++
	}
}
