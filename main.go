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
		log.Println(c.String(0))
	}
}

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatalln("missing input")
	}

	input := os.Args[1]

	token := NewLexer(strings.NewReader(input)).Lex()
	debugToken(token)

	ast := NewParser(token).Parse()
	debugAst(ast)

	byteCode := Compile(ast)
	vm := Vm{trace: true}
	vm.NewVmIn(byteCode).Execute()
	fmt.Printf("=> %f\n", vm.reg[0])
}
