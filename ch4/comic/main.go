package main

import (
	"fmt"
	"log"

	"github.com/mtanzim/gopl-book/ch4/xkcd"
)

func main() {
	cache := xkcd.NewCache()

	var input string

	for {
		fmt.Print("\n\nEnter comic id: ")
		fmt.Scanln(&input)

		if input == "q" {
			return
		}

		result, method, err := xkcd.GetComic(input, cache)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n%s\n", method, result.Transcript)
		}
	}

}
