package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", " "
	for i := 1; i < len(os.Args); i++ {
		if i == 1 {
			s += os.Args[i]
		} else {
			s += sep + os.Args[i]
		}
	}

	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args, " "))
	for i, v := range os.Args {
		fmt.Println(i)
		fmt.Println(v)
	}

}
