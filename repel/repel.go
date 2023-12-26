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
			// fmt.Fprintf(out, "%+v\n", tok)
			if !ok {
				err := l.LastError()

				fmt.Fprintf(out, "opps! \x1b[1;31merror:\x1b[0m %s\n", err.Message)
				fmt.Fprintf(out, "    | %s\x1b[1;31m%s\x1b[0m%s\n", err.Line[:err.Span[0]], err.Line[err.Span[0]:err.Span[1]], err.Line[err.Span[1]:])
				fmt.Fprintf(out, "    | ")

				for i := 0; i < err.Span[0]; i++ {
					fmt.Fprintf(out, " ")
				}
				fmt.Fprintf(out, "\x1b[1;31m^")

				for i := err.Span[0] + 1; i < err.Span[1]; i++ {
					fmt.Fprintf(out, "~")
				}

				fmt.Fprintf(out, "\x1b[0m\n")
			}
		}
	}
}
