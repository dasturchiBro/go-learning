package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net"
	"io/ioutil"
)

func main() {
	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "DELETE / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

	resp, _ := http.Get("http://aileaders.uz")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}