package main

import (
	"fmt"
	"language_x/repl"
	"os"
	"os/user"
)

func main() {
	// var t *testing.T
	// lexer.TextCaseSymbols(t)

	user, err := user.Current()
if err != nil {
panic(err)
}
fmt.Printf("Hello %s! This is the XLang programming language!\n",
user.Username)
fmt.Printf("Feel free to type in commands\n")
repl.Start(os.Stdin, os.Stdout)
}