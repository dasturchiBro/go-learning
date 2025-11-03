package main

import (
	"fmt"
)

// Working with slicing tricks and strings.

func main() {
	x := []int{1,2,3,4,5}
	y := x[1:3]
	fmt.Println(y, "\nLen: ", len(y), "\nCap: ", cap(y))
	fmt.Println("__________________ __________________")

	y = append(y, 89)
	fmt.Println("X: ", x, "\nY: ", y) // addition in Y can change X because they share the same memory

	fmt.Println("__________________ __________________")

	number := 109
	fmt.Println("In number: ", number, "\nIn string: ", string(number)) //string(number) - shows the letter in this number in the ASCII table
	fmt.Println("__________________ __________________")

	var message string = "Hello World!"
	var messageByte []byte = []byte(message)
	var messageRune []rune = []rune(message)

	fmt.Println("Message: ", message, "\nMessage in Bytes: ", messageByte, "\nMessage in Runes: ", messageRune)
	fmt.Println("__________________ __________________")

	students := make(map[string][]string, 2)
	students["A Class"] = []string{"John Smith", "Bob Hanes", "David Goggins"}
	students["B Class"] = []string{"Lisa Johnson", "Peter Parker", "William Holmes"}

	knownItem, ok := students["B Class"]
	unknownItem, ok2 := students["C Class"]
	fmt.Println("Map: ", students, "\nUnknown value: ", unknownItem, "\nOk: ", ok2,"\n\nKnown value: ", knownItem, "\nOK: ", ok)

	delete(students, "B Class") //used to delete an item from a map
	fmt.Println(students)
	fmt.Println("__________________ __________________")

	numbers := []int{1,2,4,2,3,1,2,3,2,3,1,5}
	fmt.Println("The length of the slice numbers: ", len(numbers))
	intSet := map[int]bool{}

	for _, value := range numbers {
		intSet[value] = true // Go doesn't have sets. We gotta implement that using map.
	}

	fmt.Println("The length of the map: ", len(intSet))	
	fmt.Println("100 is in the set: ", intSet[100])
	fmt.Println("1 is in the set: ", intSet[1])
	fmt.Println("5 is in the set: ", intSet[5])
	fmt.Println("__________________ __________________")

	type Person struct {
		Name string  
		Age int  
		IsEmployed bool
		FirstName string
	}
	john := Person{Name: "John Holmes", Age: 21, IsEmployed: true, FirstName: "John"}
	fmt.Println(john)
	john.Age = 22
	fmt.Printf("%s's age is %d\n", john.FirstName, john.Age)

	company := struct{ //anonymous struct
		Name string
		EstablishedYear int
	}{
		"Softer Hardware Inc.", 1997,
	}

	fmt.Println(company)
}