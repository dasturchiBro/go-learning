package main

import ("fmt")


type Person struct {
	Name string
	Message *string
}

func (p Person) changeMessage(newMessage string) {
	*p.Message = newMessage
}

func main() {
	// Pointers
	var x int = 1005
	var y bool = true
	var z = []int{1,3,5,6,7}

	var pointerX = &x
	var pointerY = &y
	var pointerZ = &z
	var pointerOfpointerZ = &pointerZ
	var pointerH *string

	fmt.Println("***Pointers: \n", pointerX, "\n", pointerY, "\n", **&*&pointerOfpointerZ, "\n", pointerH)

	// &: values -> address
	// *: address -> value
	message := "HELLO WORLD!"
	john := Person{
		"John",
		&message,
	}
	fmt.Println(john, "\tHere is the message variable: ", message)
	john.changeMessage("New Message HERE!!!")
	fmt.Println(john, "\tHere is the message variable: ", message)
	
}