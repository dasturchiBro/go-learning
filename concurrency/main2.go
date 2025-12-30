package main 

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)
	// channel <- "John"
	// channel <- "Peter"

	// <- channel
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)

	go func() {
		channel <- "John"
		channel <- "Peter"
	}()

	fmt.Println(<- channel)
	// fmt.Println(<- channel)
	close(channel)
}