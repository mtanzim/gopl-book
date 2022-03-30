package main

import (
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

func main() {
	url := "https://www.gamespot.com"
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(list)

}
