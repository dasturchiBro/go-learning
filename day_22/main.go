package main

import (
	"fmt"
)

func changeValue(name string, num int, mapForChange map[string]int) {
	_, ok := mapForChange[name]
	if !ok {
		fmt.Println("A value with this name not found!")
	}
	mapForChange[name] = num
}

func changeSlice(sliceForChange []string, index int, value string) {
	sliceForChange[index] = value

	sliceForChange = append(sliceForChange, "New Val 1", "New Person 2", "New Person 3")
}

func main() {
	// Is Zero value zero?
	var num *int 
	fmt.Printf("\nType of num: %T\nValue of num: %d\n%v\n%v\n\n", num, num, num==nil, num==func(i int) *int {return &i}(0))

	// Maps change inside functions
	var books = map[string]int{
		"Sherlock Holmes": 324,
		"Atomic Habits": 254,
		"Anna Karenina": 1018,
		"A confession": 60,
		"Think and Grow Rich": 354,
	}

	fmt.Println("***Current map: ", books)
	changeValue("Sherlock Holmes", 0, books)
	changeValue("Atomic Habits", 1000, books)
	fmt.Println("***Changed map: ", books, "\n")
	changeValue("aaaaaa", 0, books)

	// What about slices?
	people := []string{"John", "Richard", "James", "Johnson", "Hurry", "Pet", "Michael", "Hanes"}
	fmt.Println("\n\n***Current slice: ", people, "\nCurrent len and cap: ", len(people), " ", cap(people), "\n")
	changeSlice(people, 0, "Sherlock Holmes  - changed value")
	fmt.Println("\n\n***Changed slice: ", people, "\nCurrent len and cap: ", len(people), " ", cap(people), "\n")

}