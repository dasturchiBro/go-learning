package main

import (
	"fmt"
	"strconv"
)

// Types
type Age int

type Person struct {
	Name string
	Age Age  
	Address string
	PhoneNum PhoneNumber
}

type PhoneNumber int

// Methods
func (p Person) GiveInfo() string {
	res := "\n***User info:\nName: "+p.Name+"\nAge: "+strconv.Itoa(int(p.Age))+"\nAddress: "+p.Address
	if p.PhoneNum != 0 {
		res += "\nPhone Number: "+strconv.Itoa(int(p.PhoneNum))+"\n\n"
	}
	return res+"\n\n"
}

func (p *Person) ChangeNum(pn PhoneNumber) {
	p.PhoneNum = pn
}


func main() {
	var Johns_num PhoneNumber
	fmt.Println("Enter John's phone number: ")
	fmt.Scanln(&Johns_num)
	fmt.Println("John's Phone Number is ", Johns_num)


	john := Person{
		Name: "John Doe",
		Age: 32,
		Address: "13th Street. 47 W 13th St, New York, NY 10011, USA",
	}
	fmt.Println(john.GiveInfo())
	john.ChangeNum(Johns_num)
	fmt.Println(john.GiveInfo())
	numChanger := (*Person).ChangeNum
	numChanger(&john, 000000)
	info := Person.GiveInfo

	fmt.Println(info(john))
}