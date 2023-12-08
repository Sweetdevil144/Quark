package main

import (
	"Quark/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Welcome to Quark ", currentUser)
	repl.Start(os.Stdin, os.Stdout)
}
