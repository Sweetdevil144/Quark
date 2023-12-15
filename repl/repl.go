package repl

import (
	"Quark/lexer"
	"Quark/token"
	"bufio"
	"fmt"
	"io"
)

//REPL stands for “Read Eval Print Loop”
//It is a simple interactive programming environment that takes user inputs, executes them, and returns the result to the user.

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)
		for tok := l.TokenizeInputString(); tok.Type != token.EOF; tok = l.TokenizeInputString() {
			fmt.Printf("%+v\n", tok)
		}

	}
}
