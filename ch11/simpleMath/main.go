package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// note that setting this as a variable allows us to modify it during testing
var out io.Writer = os.Stdout

var n = flag.Int("n", 0, "specify an integer")

func main() {
	flag.Parse()
	addTwo(*n)
}

func addTwo(n int) {
	fmt.Fprintf(out, "got %d", n+2)

}
