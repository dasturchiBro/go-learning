package main

import (
	"fmt"
	"net"
	"log"
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

	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())

	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Print(err)
	}
}