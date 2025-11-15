package main

import (
	"fmt"
	"errors"
)

type Student struct {
	Name string
	Age int  
	AvgScore float64
}

func addStudent(students map[string]Student, id string, student Student) error {
	if int(student.AvgScore) < 55 {
		return errors.New("can't add a student with an average score below 55")
	}
	students[id] = student
	return nil
}

func addStudent_second(students map[string]Student, id string, student Student) error {
	if int(student.AvgScore) < 55 {
		return fmt.Errorf("the student's average score (%v) is below 55", student.AvgScore)
	}
	students[id] = student
	return nil
}

type Sentinel string
func (s Sentinel) Error() string {
	return string(s)
}

func main() {
	// First method of printing an error
	students := make(map[string]Student)
	students["#2hf33"] = Student{"John Doe", 19, 67.5}
	students["#1fg13"] = Student{"Hurry Smith", 21, 89.1}
	students["#4ad59"] = Student{"Thomas Holmes", 34, 77}
	students["#0hj98"] = Student{"Peter Parker", 24, 55}


	func() {err := addStudent(students, "$4JK23", Student{"Tom Web", 22, 67.6})
	if err != nil {
		fmt.Println("Something went wrong: ", err)
	} else {
		fmt.Println("Student added successfully!\n", students)
	}}()

	func() {
		err := addStudent(students, "#4JK23", Student{"Hanes Thompson", 32, 54.6})
		if err != nil {
			fmt.Println("Something went wrong: ", err)
		} else {
			fmt.Println("Student added successfully!\n", students)
		}		
	}()


	// Second method of printing an error

	func() {
		err := addStudent_second(students, "#4J313", Student{"Hanes Thompson", 32, 54.6})
		if err != nil {
			fmt.Println("Something went wrong: ", err)
		} else {
			fmt.Println("Student added successfully!\n", students)
		}		
	}()


	// Sentinel Errors
	ErrBar := Sentinel("bar error")
	fmt.Println(ErrBar)

	fmt.Println(ErrBar.Error())
}