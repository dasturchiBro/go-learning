package main

import (
	"fmt"
)

// Methods in Go

type Person struct {
	Name string
	Age int
}

func (p Person) HelloWorld() {
	fmt.Printf("Hello %s.\nYour age is %d\n\n", p.Name, p.Age)
}

func (p Person) SayGoodbye() {
	fmt.Printf("Goodbye %s\n", p.Name)
}

type People []Person

func (p People) HelloWorld() {
	for _, value := range p {
		fmt.Printf("Hello, %s\n", value.Name)
	}
}


func (p *Person) changeAge(newAge int) {
	p.Age = newAge
}

func (p Person) changeName(Name string) {
	p.Name = Name
}


func main() {
	john := Person{Name: "John Doe", Age: 14}
	john.HelloWorld()
	john.SayGoodbye()
	data := People{
		{"John Due", 78},
		{"David Goggins", 50},
		{"Peter Smith", 18},
	}
	data.HelloWorld()

	john.changeAge(15)
	john.changeName("John Due")
	john.HelloWorld()
}


