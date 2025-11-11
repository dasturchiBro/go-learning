package main

import (
	"fmt"
)

type Score int 
type HighScore Score
type Number int

func (s Score) printScore() {
	fmt.Println("Your score is: ", s)
}

func (n Number) decreaseBy(b Number) Number {
	return n - b
}

func (n Number) increaseBy(b Number) Number {
	return n + b
}

type WeekDay int

type Person struct {
	Name string
	Age int  
}

type Employee struct{
	Person
	Salary int
	PhoneNumber int
	Company string
}

func main() {
	var myscore Score = 101
	var myhighestscore HighScore = 120
	myscore.printScore()
	fmt.Println("Your highest score: ", myhighestscore)
	fmt.Println("Is it true: ", Score(myhighestscore) == myscore)
	// myhighestscore.printScore() - this gives an error

	var aNumber Number = 324
	fmt.Println("Current number: ", aNumber)
	fmt.Println("Increased by 401: ", aNumber.increaseBy(401))
	fmt.Println("Decreased by 1000: ", aNumber.decreaseBy(1000))

	const (
		Monday WeekDay = 1 << iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday, " ", Tuesday, " ", Wednesday, " and so on")


	// Embedded fields
	john := Employee{
		Person: Person{
			"John Doe",
			24,
		},
		Salary: 10_000,
		PhoneNumber: 1_569_739_638,
		Company: "Google",
	}

	fmt.Println(john.Name, "works at", john.Company+".", "He is", john.Age, "years old and earns", john.Salary, "every month.")
}