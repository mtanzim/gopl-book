package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

type clientWithName struct {
	ch  client
	who string
}

var (
	entering = make(chan *clientWithName)
	leaving  = make(chan *clientWithName)
	messages = make(chan string)
)

func broadcaster() {

	clients := make(map[*clientWithName]bool)

	broadcastNames := func(clients map[*clientWithName]bool) {
		clientNames := []string{}
		for client := range clients {
			clientNames = append(clientNames, client.who)
		}
		for client := range clients {
			msg := fmt.Sprintf("The following clients are present %v\n", clientNames)
			client.ch <- msg
		}
	}

	for {
		select {
		case msg := <-messages:
			for client := range clients {
				client.ch <- msg
			}
		case client := <-entering:
			clients[client] = true
			broadcastNames(clients)
		case client := <-leaving:
			delete(clients, client)
			close(client.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	// go clientTimer(conn, ch)
	who := conn.RemoteAddr().String()
	client := &clientWithName{ch, who}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- client
	messages <- who + "has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func clientTimer(conn net.Conn, ch chan string) {
	timer := time.NewTimer(2 * time.Second)
	who := conn.RemoteAddr().String()
	for {
		select {
		case <-timer.C:
			leaving <- &clientWithName{ch, who}
			messages <- who + "has timed out"
			conn.Close()
			return
		case <-ch:
			timer.Stop()
			timer = time.NewTimer(2 * time.Second)
		default:
			// do nothing
		}
	}
}
