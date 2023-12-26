package repel

import (
	"bufio"
	"fmt"
	"os"

	"github.com/akash1047/crepel/lexer"
	"github.com/akash1047/crepel/token"
)

func Start(in, out *os.File) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, "> ")

		if scanned := scanner.Scan(); !scanned {
			return
		}

		input := scanner.Text()

		l := lexer.New(input)

		for tok, ok := l.NextToken(); tok.Type != token.EOF; tok, ok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
			if !ok {
				fmt.Fprintf(out, "\x1b[31merror:\x1b[0m %s\n", l.LastError())
			}
		}
	}
}
