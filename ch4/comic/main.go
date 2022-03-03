package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mtanzim/gopl-book/ch4/xkcd"
)

func main() {
	result, err := xkcd.GetComicFromRemote(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}
