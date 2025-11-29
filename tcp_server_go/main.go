package main

import (
	"fmt"
	"net"
	"log"
	"time"
	"bufio"
	"strings"
)


func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	fmt.Println("Listening on :80")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		conn.SetDeadline(time.Now().Add(30 * time.Second))
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}

		fmt.Printf("New messsage: ", msg)

		var command string
		fmt.Println("Enter your command: ")
		fmt.Scanln(&command)
		_, err = writer.WriteString(command)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
		}
		writer.Flush()
	}
}