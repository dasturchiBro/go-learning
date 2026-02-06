package main

import (
	"fmt"
	"time"
	"sync"
)

func CheckServer(server string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second) // ping simulation
	ch <- fmt.Sprintf("Server %s is UP!", server)
}
func CheckServer2(server string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second) // ping simulation
	ch <- fmt.Sprintf("Server %s is UP!", server)
}

func main() {
	servers := []string{"Google", "AWS", "UzCloud", "GitHub"}
	servers2 := []string{"DOO", "Dash", "ServeZ", "Proce"}

	results := make(chan string)
	var wg sync.WaitGroup

	for _, server := range servers {
		wg.Add(1)
		go CheckServer(server, results, &wg)
	}

	for _, server := range servers2 {
		wg.Add(1)
		go CheckServer2(server, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("\n===All checks are complete===")
}