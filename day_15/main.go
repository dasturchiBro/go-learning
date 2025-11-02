package main

import (
	"fmt"
)

func main() {
	var arr [3]int //An array - cannot add or remove values
	var slice []string //A slice -> append is used to add values to a slice

	//len() - used to find the length of an array or slice
	arr[0] = 1
	arr[1] = 2 
	arr[2] = 3
	/*for true {
		var answer string
		fmt.Print("[Enter /exit to exit]\nEnter what you want to add to the slice: ")
		fmt.Scanln(&answer)
		if answer == "/exit" {break}
		slice = append(slice, answer)
		fmt.Println(slice)
		fmt.Println("Length: ", len(slice), "\nCapacity: ", cap(slice))
	}*/
	
	fmt.Println("\nHere is your slice: ", slice)

	// make() - used to create slices, giving them capacity
	newSlice := make([]uint64, 0, 10000000000)
	newSlice2 := make([]uint64, 0, 10000000000)
	newSlice3 := make([]uint64, 0, 10000000000)
	fmt.Println("Len: ", len(newSlice), "\nCapacity: ", cap(newSlice))
	fmt.Println("Len: ", len(newSlice2), "\nCapacity: ", cap(newSlice2))
	fmt.Println("Len: ", len(newSlice3), "\nCapacity: ", cap(newSlice3))

}