package main

import (
	"Quark/lexer"
	"fmt"
)

func main() {
	input := "123 + 456 - 789"

	l := lexer.NewLexer(input)

	for {
		tok := l.NextToken()
		if tok.Type == lexer.EOF {
			break
		}
		fmt.Printf("%+v\n", tok)
	}
}
