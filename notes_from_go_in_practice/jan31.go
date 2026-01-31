package main

import (
	"fmt"
	"time"
)

// func count() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		time.Sleep(1 * time.Second)
// 	}
// }

func PrintCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
		time.Sleep(500 * time.Millisecond)
	}
}

func getName() string {
	return "John"
}

func main() {
	// go count()
	// time.Sleep(1 * time.Second)
	// fmt.Println("Hello World!!!")
	// time.Sleep(5 * time.Second)

	c := make(chan int)
	s := []int{23,12,12,4,2,5,1,3,7,954,21,63,-1}

	go PrintCount(c)

	fmt.Println("Middle")

	go func() {
		for _, v := range s {
			c <- v
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("END")
	fmt.Println("Hello", getName())
	time.Sleep(6 * time.Second)
}