package main

import (
	"fmt"
	"os"
)

func debugToken(token []Token) {
	fmt.Printf("%5s | %15s | %15s \n\n", "index", "type", "raw")
	for i, t := range token {
		fmt.Printf("%5d | %15s | %15s \n", i, TOKEN_LOOKUP[t.Type], t.Raw)
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("no file given")
	}

	input := os.Args[1]
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	lexer := NewLexer(file)
	token := lexer.Lex()
	debugToken(token)
	parser := NewParser(token)
	ast := parser.Parse()
	Eval(ast)
}
