package main

import (
	"fmt"
	"os"

	"github.com/akash1047/crepel/repel"
)

func main() {
	fmt.Println("Hi User, This is crepel.")
	repel.Start(os.Stdin, os.Stdout)
}
