package main

import ("fmt")

type myNum int

func main() {
	var i interface{}
	var number myNum = 100
	i = number
	fmt.Println(i.(myNum))

	theNum, ok := i.(int)
	if ok {
		fmt.Println("There is the number: ", theNum)
	} else {
		fmt.Println("Value with this type doesn't exist!")
	}

	switch i.(type) {
	case int: fmt.Println("Type INT")
	case myNum: fmt.Println("Type myNum")
	case string: fmt.Println("Type string") 
	}

}