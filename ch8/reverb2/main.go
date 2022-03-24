package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// the `go` here allows multiple clients (ie: netcat)
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		// `go` here allows a single client to interleave multiple inputs, as oppposed to hearing back one set of echo at a time
		go echo(c, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Println(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Println(c, "\t", shout)
	time.Sleep(delay)
	fmt.Println(c, "\t", strings.ToLower(shout))

}
