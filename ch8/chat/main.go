package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// To use on Mac
// brew install nmap
// ncat localhost 8080
func main() {
	listener, err := net.Listen("tcp", ":8080")
	fmt.Println("Server listening")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for client := range clients {
				client <- msg
			}
		case client := <-entering:
			clients[client] = true
		case client := <-leaving:
			delete(clients, client)
			close(client)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + "has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
