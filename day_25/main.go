package main

import ("fmt")

// type Employee struct {
// 	Name string
// 	Age int
// 	Email string
// }

// type Manager struct {
// 	Employee
// 	Staff []Employee
// 	Id int
// }


// func (e Employee) Work() {
// 	fmt.Println("Employee is working...")
// }

// func (m Manager) Work() {
// 	fmt.Println("Manager is working...")
// }

// type Worker interface {
// 	Work()
// }

// func startWork(w Worker) {
// 	w.Work()
// }


// Exercise =>

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	PI := 3.14
	return PI*c.Radius*c.Radius
}

func (c Circle) Perimeter() float64 {
	const PI = 3.14
	return 2*PI*c.Radius
}

type Rectangle struct {
	Height float64
	Width float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Triangle struct {
	Height float64
	Base float64
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base)/2.0
}

type Shape interface {
	Area() float64
}

func getArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}

type ShapeDetails interface {
	Shape  
	Perimeter() float64
}

func printDetails(d ShapeDetails) {
	fmt.Println("Shape Area: ", d.Area(), "\nShape Perimeter: ", d.Perimeter())
}

func main() {
	// var John Manager
	// John.Name = "John Doe"
	// John.Age = 43
	// John.Email = "john@google.com"
	// John.Id = 3454
	// var e Employee
	// // e = John - this gives as error
	// e = John.Employee // this works
	// fmt.Println(John, e)
	// e.Work()
	// John.Work()
	
	// // Interfaces
	// fmt.Println("\n")
	// startWork(John)

	circle := Circle{
		9,
	}
	rec := Rectangle{
		10,
		6,
	}
	tri := Triangle{
		5,
		10,
	}
	getArea(circle)
	getArea(rec)
	getArea(tri)
	printDetails(circle)
}