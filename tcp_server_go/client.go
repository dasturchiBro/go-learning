package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "HELLO message from the client\n")
	
	resp, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("-Response: ", resp)
}