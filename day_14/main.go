package main

import (
	"fmt"
)

/* type Person struct {
	Name string
	Age int 
}

type Contact struct {
	Email string
	PhoneNumber int 
}

type Employee struct{
	Person
	Contact
	EmployeeID int
}

func (p Person) greet() {
	fmt.Println("Hello, ", p.Name)
}
*/

type Vector struct {
	X, Y float64
}

func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y}
}

type Shape []Vector

func (s Shape) Transform(offset Vector, add bool) {
	var fn func(p, q Vector) Vector
	if add {
		fn = Vector.Add
	} else {
		fn = Vector.Sub
	}

	for i := range s {
		s[i] = fn(s[i], offset)
	}
}

func get_vectors() (Vector, Vector) {
	var a, b, c, d float64
	fmt.Println("Enter the first vector in this form: 0 0")
	fmt.Scan(&a, &b)
	fmt.Println("Enter the second vector in this form: 0 0")
	fmt.Scan(&c, &d)
	v1 := Vector{}
	v1.X = a
	v1.Y = b  
	v2 := Vector{}
	v2.X = c   
	v2.Y = d
	return v1, v2
}

func main() {
	shape := Shape{{0,0}, {1, 0}, {0, 1}}
	fmt.Println(shape)

	tran := shape.Transform
	offset := Vector{2, 3}
	tran(offset, true)
	fmt.Println(shape)
	tran(offset, false)
	fmt.Println(shape)


	ops := map[string]func(a, b Vector) Vector{
		"sub": Vector.Sub,
		"add": Vector.Add,
	}

	fmt.Println("\n***Actions** \nadd - [add vectors]\nsub - [subtract vectors]")
	var answer string
	fmt.Scanln(&answer)
	if answer != "add" && answer != "sub" {
		return
	}
	a, b := get_vectors()
	fmt.Println(ops[answer](a, b))


/*
	john := Employee{}
	john.Name = "John Doe"
	john.Age = 25
	john.Email = "john@motiff.com"
	john.PhoneNumber = 99898998989
	john.EmployeeID = 1322
	john.greet()
*/

}