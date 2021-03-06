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

	broadcastNames := func() {
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
				select {
				case client.ch <- msg:
					// do nothing
				case <-time.NewTimer(2 * time.Second).C:
					// in case messages are dropped, do not block
					fmt.Println("Message skipped for " + client.who)
				}
			}
		case client := <-entering:
			clients[client] = true
			broadcastNames()
		case client := <-leaving:
			delete(clients, client)
			close(client.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	who := ""
	fmt.Fprintln(conn, "Please enter your name: ")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		who = input.Text()
		break
	}
	go clientWriter(conn, ch)
	timeoutVal := time.Duration(3600)
	client := &clientWithName{ch, who}

	reset := make(chan struct{})
	go func() {
		for alive := true; alive; {
			timer := time.NewTimer(timeoutVal * time.Second)
			select {
			case <-reset:
				timer.Stop()
			case <-timer.C:
				alive = false
				messages <- who + " has timed out"
				fmt.Fprintln(conn, "You have timed out!")
				conn.Close()
			}
		}
	}()

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client

	for input.Scan() {
		messages <- who + ": " + input.Text()
		reset <- struct{}{}
	}
	leaving <- client
	messages <- who + " has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
