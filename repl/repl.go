package repl

import (
	"bufio"
	"fmt"
	"io"
	"language_x/lexer"
	"language_x/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	
scanner := bufio.NewScanner(in)
for {
	fmt.Printf(PROMPT)
	scanned := scanner.Scan()
	if (!scanned) {
		return
	}
	input := scanner.Text()

l := lexer.New(input)

for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){
	fmt.Printf("%+v\n", tok)
}

}

}