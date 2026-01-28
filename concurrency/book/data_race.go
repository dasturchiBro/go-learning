package main

import "fmt"


func main() {
	var data int
	go func() {
		data+=100
	}()

	if data == 0 {
		fmt.Println("The value of data is ", data)
	}

	var data2 int
	go func() {
		data2+=100
	}()

	if data2 == 0 {
		fmt.Println("The value of data is ", data2)
	}

	var data3 int
	go func() {
		data3+=100
	}()

	if data3 == 0 {
		fmt.Println("The value of data3 is ", data3)
	}

	var data4 int
	go func() {
		data3+=100
	}()

	if data4 == 0 {
		fmt.Println("The value of data3 is ", data4)
	}

	var data5 int
	go func() {
		data4+=100
	}()

	if data5 == 0 {
		fmt.Println("The value of data4 is ", data5)
	}
}

// A data race in Golang