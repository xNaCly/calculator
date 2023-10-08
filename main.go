package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func debugToken(token []Token) {
	log.Printf("%5s | %15s | %15s \n\n", "index", "type", "raw")
	for i, t := range token {
		log.Printf("%5d | %15s | %15s \n", i, TOKEN_LOOKUP[t.Type], t.Raw)
	}
}

func debugAst(ast []Node) {
	for _, c := range ast {
		fmt.Println(c.String(0))
	}
}

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatalln("no input given")
	}

	input := os.Args[1]
	if len(input) == 0 {
		log.Fatalln("no input")
	}

	lexer := NewLexer(strings.NewReader(input))
	token := lexer.Lex()

	if len(token) == 1 {
		log.Fatalln("only got EOF, probably an error? ig?")
	}

	debugToken(token)

	// ast := []Node{
	// 	&Addition{
	// 		left: &Multiplication{
	// 			left:  &Number{token: Token{Raw: "2"}},
	// 			right: &Number{token: Token{Raw: "1"}},
	// 		},
	// 		right: &Number{token: Token{Raw: "2"}},
	// 	},
	// }

	parser := NewParser(token)
	ast := parser.Parse()
	if len(ast) == 0 {
		log.Fatalln("parsing error, i think :^)")
	}

	debugAst(ast)

	byteCode := Compile(ast)
	vm := Vm{trace: true}
	vm.NewVmIn(byteCode).Execute()
}
