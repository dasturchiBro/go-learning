package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// display("START")
	// go display("START", &wg, 1.0 * time.Second)
	// go display("END", &wg, 500 * time.Millisecond)

	go func() {
		display("START", 1.0 * time.Second)
		wg.Done()
	}()

	go func() {
		display("END", 500 * time.Millisecond)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("HELLO WORLD!!!")
}

// func display(input string, wg *sync.WaitGroup, num time.Duration) {
// 	for i := 1; i < 19; i++ {
// 		fmt.Println(i, " - ", input)
// 		time.Sleep(num)
// 	}
// 	wg.Done()
// }

func display(input string, num time.Duration) {
	for i := 1; i < 2; i++ {
		fmt.Println(i, " - ", input)
		time.Sleep(num)
	}
}