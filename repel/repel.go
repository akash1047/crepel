package repel

import (
	"bufio"
	"fmt"
	"os"
)

func Start(in, out *os.File) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, "> ")

		if scanned := scanner.Scan(); !scanned {
			return
		}

		input := scanner.Text()

		fmt.Fprintln(out, input)
	}
}
