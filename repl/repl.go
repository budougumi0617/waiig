package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/budougumi0617/waiig/lexer"
	"github.com/budougumi0617/waiig/token"
)

// PROMPT is prefix for prompt line.
const PROMPT = ">> "

// Start starts REPL of the Monkey language.
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
