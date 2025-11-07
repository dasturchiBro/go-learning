package main

import (
	"fmt"
	"sort"
	"os"
	"log"
)

type Person struct {
	Firstname string  
	Age int
	Height int
}

func mainFunc(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
// Closures //
	people := []Person{{"John", 17, 175}, {"Shakespear", 65, 156}, {"Thomas", 21, 186}, {"Emily", 19, 162}, {"Robert", 34, 179}, {"Sophia", 23, 168}, {"Michael", 28, 182}, {"Olivia", 16, 158}, {"Daniel", 40, 190}, {"Isabella", 22, 165}}
	fmt.Println(people)

	// Sort by Height
	sort.Slice(people, func(i, j int) bool {
		return people[i].Height > people[j].Height
	})
	fmt.Println("\nSorted by Height: ")
	fmt.Println(people)

	// Sort by Age
	sort.Slice(people, func(j, i int) bool {
		return people[j].Age > people[i].Age
	})
	fmt.Println("\nSorted by Age: ")
	fmt.Println(people)

	// Return closures
	fmt.Println("\n***Base two values: ")
	baseTwo := mainFunc(2)
	for i := 1; i < 10; i++ {
		fmt.Println(baseTwo(i))
	}

	fmt.Println("\n***Base three values: ")
	baseThree := mainFunc(3)
	for i := -10; i != 0; i++ {
		fmt.Println(baseThree(i))
	}


// defer
	fmt.Println("\n***Working with defer:")
	file_name := "main.txt"
	f, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		fmt.Println(string(data[:count]))
		if err != nil {
			break
		}
	}
}