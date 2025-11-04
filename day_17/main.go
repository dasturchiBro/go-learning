package main

import ("fmt")

func getSomething() string {
	// return "Something something something something ..."
	return "Hello"
}

func main() {
	//Blocks
	x := 100
	if x < 101 {
		fmt.Println("Current: ", x)
		x := 200
		fmt.Println("Current X:", x)
	}
	fmt.Println("Final X: ", x)
	fmt.Println("_____________________")

	//if-else-if-else
	if answer := getSomething(); answer == "Hello" {
		fmt.Println("Hello Woooooooooorld!!!!")
	} else if answer == "World" {
		fmt.Println("Heeeeeeeeeeello World!")
	} else {
		fmt.Println("Say something...")
	}
	fmt.Println("_____________________")

	// fmt.Println(answer) - .\main.go:30:14: undefined: answer


	// A complete for statement
	for i := 1; i < 15; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// The Condition-Only for statement
	i := 0.01
	for i < 0.09 {
		fmt.Printf("%.4f - ", i)
		i *= 1.53234234276543456
	}


	// The Infinite for Statement
	for {
		fmt.Println(" ")
		fmt.Println("enter something to exit the infinite loop... ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "continue" {
			continue
		}

		if answer != "" {
			break
		} 
	}

	// The for-range loop
	mySlice := []int{1432,2341,442,960, 560, 4154, 829, 4189}
	mySlice = append(mySlice, (mySlice)...)
	for i, val := range mySlice {
		if i%2 == 0 {
			val *= i/2*21
		} else {
			val *= (3+i-(i/2))
		}
		fmt.Print(string(val), " ")
	}
	fmt.Println()

	// Labeling loops
outer:
	for _, s := range []string{"John", "Sherlock Holmes", "William Shakespear", "Bob", "Al Pachino", "Ludovic"} {
		for _, letter := range s {
			fmt.Println(string(letter))
			if string(letter) == "l" || string(letter) == "h" {
				fmt.Println("\n")
				continue outer
			}
		}
	}
}