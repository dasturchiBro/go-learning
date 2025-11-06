package main

import (
	"errors"
	"fmt"
)


type FunOpts struct {
	Name string  
	Age int  
}

func sayHello(person FunOpts) {
	fmt.Printf("Hello, %s\nI know that your age is %d\n", person.Name, person.Age)
}


func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

/*&func sayHello(name string, age int) {
	fmt.Printf("Hello, %s\nI know that your age is %d\n", name, age)
}*/


func DivAndRem(num, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("Can't be divided by zero!")
	}
	return num/den, num%den, nil
}

func add(first, second int) int {
	return first + second
}

func sub(first, second int) int {
	return first - second
}

type MathFunc func(int, int) int


func main() {
	// Working with functions
	// sayHello("John", 34)

	// sayHello("John") - this does not work because we have to give value to age
	sayHello(FunOpts{Name: "John"})


	//Variadic Parameters
	fmt.Println(sum(1,343,542,234,524,1))
	fmt.Println(sum())
	fmt.Println(sum(0,1,1,1,1,))

	// Multiple Return Values
	fmt.Println(DivAndRem(123,123))
	fmt.Println(DivAndRem(123,0))
	fmt.Println(DivAndRem(-1,2))


	// Functions like values
	opts := map[string]MathFunc{
		"+": add,
		"-": sub,
	}
	var first, second int  
	fmt.Println("Enter the first and second numbers: ")
	fmt.Scanln(&first, &second)

	var opt string
	fmt.Println("Choose the option: + -")
	fmt.Scanln(&opt)

	optFunc, ok := opts[opt]
	if !ok {
		fmt.Println(errors.New("Such option does not exist!"))
	}
	fmt.Println(optFunc(first, second))

	// Anonymous Functions
	for i := range []int{1,34,113,143,514} {
		fmt.Println(func(j int) string {
			if i % 2 == 0 {
				return "Even"
			}
			return "Odd"
		}(i))
	}

}