package main

import (
	"fmt"
	"time"
)

func increment(num *int) {
	(*num)++
}

type Person struct {
	Name string
	Age int
}

func (p *Person) birthday() {
	p.Age++
}

func good() *int {
	x := 10 // normally on stack
	return &x // x escapes and moves to heap
}

func slices() *[]int {
	s := []int{341, 143, 532, 2}
	return &s
}

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Wooof Woof!!!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow... Meow."
}

func Announce(s Speaker) {
	fmt.Println(s.Speak())
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	A float64
	B float64
}

func (r Rectangle) Area() float64 {
	return r.A * r.B
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return 3.14 * (c.r)*(c.r)
}

func PrintArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}

func say(s string) {
	time.Sleep(1 * time.Second)
	fmt.Println(s)
}

func main() {
	num := 43
	fmt.Println("Number before: ", num)
	increment(&num)
	increment(&num)
	increment(&num)

	fmt.Println("Number num: ", num, "\n")

	person := Person{"Fibbo", 16}
	fmt.Println(person)

	person.birthday()
	fmt.Println(person)


	fmt.Println(*(good()), " - ", good())
	fmt.Println(*(slices()), " - ", slices())

	var d Dog
	var c Cat

	Announce(d)
	Announce(c)

	r := Rectangle{
		13,
		14,
	}

	ci := Circle{
		45,
	}

	PrintArea(r)
	PrintArea(ci)


	go say("Hello")
	go say("Hi")
	go say("Hey")
	time.Sleep(2 * time.Second)

}


/* 
	Memory Management in Golang
There are two main memory zones: Stack and Heap

The stack:
	Holds local variables
	Fast (de)allocation
	Automatic cleaning
	Small, short-lived stuff

The heap: 
	Stores variables outside the stack
	Larger, slower to allocate
	Garbage collected by Goroutine

Use the go build -gcflags="-m" to check escape analysis

Result:
# tech
./main.go:7:6: can inline increment
./main.go:16:6: can inline (*Person).birthday
./main.go:20:6: can inline good
./main.go:25:6: can inline slices
./main.go:32:13: inlining call to fmt.Println
./main.go:33:11: inlining call to increment
./main.go:34:11: inlining call to increment
./main.go:35:11: inlining call to increment
./main.go:37:13: inlining call to fmt.Println
./main.go:40:13: inlining call to fmt.Println
./main.go:42:17: inlining call to (*Person).birthday
./main.go:43:13: inlining call to fmt.Println
./main.go:46:20: inlining call to good
./main.go:46:36: inlining call to good
./main.go:46:13: inlining call to fmt.Println
./main.go:47:22: inlining call to slices
./main.go:47:40: inlining call to slices
./main.go:47:13: inlining call to fmt.Println
./main.go:7:16: num does not escape
./main.go:16:7: p does not escape
./main.go:21:2: moved to heap: x
./main.go:26:2: moved to heap: s
./main.go:26:12: []int{...} escapes to heap
./main.go:46:36: moved to heap: x
./main.go:47:40: moved to heap: s
./main.go:32:13: ... argument does not escape
./main.go:32:14: "Number before: " escapes to heap
./main.go:32:33: num escapes to heap
./main.go:37:13: ... argument does not escape
./main.go:37:14: "Number num: " escapes to heap
./main.go:37:30: num escapes to heap
./main.go:37:35: "\n" escapes to heap
./main.go:40:13: ... argument does not escape
./main.go:40:14: person escapes to heap
./main.go:43:13: ... argument does not escape
./main.go:43:14: person escapes to heap
./main.go:46:13: ... argument does not escape
./main.go:46:14: *(~r0) escapes to heap
./main.go:46:25: " - " escapes to heap
./main.go:47:13: ... argument does not escape
./main.go:47:14: *(~r0) escapes to heap
./main.go:47:22: []int{...} escapes to heap
./main.go:47:27: " - " escapes to heap
./main.go:47:40: []int{...} escapes to heap

*/