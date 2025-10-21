package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

func loader(t time.Duration) {
	fmt.Print(".")
	time.Sleep((t * time.Millisecond))


	fmt.Print(".")
	time.Sleep((t * time.Millisecond))


	fmt.Print(".")
	time.Sleep((t * time.Millisecond))

	fmt.Print(".")
	time.Sleep((t * time.Millisecond))


	fmt.Print(".")
	time.Sleep((t * time.Millisecond))

	fmt.Print(".\n")
	time.Sleep((t * time.Millisecond))	
}

func guesser_game() {
	loader(200)
	var max, min int  
	fmt.Println("--Welcome to the Guesser game.-- \nEnter the range(max min): ")
	fmt.Scanln(&max, &min)
	var random_number int = rand.IntN((max - min)) + min
	for true {
		var guessed_num int
		fmt.Println("Guess the number: ")
		fmt.Scanln(&guessed_num)
		loader(200)
		if guessed_num == random_number {
			fmt.Println("YOU ARE ROCK!!!")
			break
		} else if guessed_num > random_number {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}
	}
}

func temperature_converter() {
	loader(200)
	var type_of string
	var num float64
	fmt.Println("--Enter the type of conversion:\na) Celcius to Farenheit\nb) Farenheit to Celcius")
	fmt.Scanln(&type_of)


	switch type_of {
		case "a":
			fmt.Println("Enter the value: ")
			fmt.Scanln(&num)
			loader(200)
			fmt.Println(strconv.FormatFloat(from_celcius_to_far(num), 'f', -1, 64))
		case "b":
			fmt.Println("Enter the value: ")
			fmt.Scanln(&num)
			loader(200)
			fmt.Println(strconv.FormatFloat(from_far_to_celcius(num), 'f', -1, 64))
		default:
			fmt.Println("Undefined type!")
			main_menu()
	} 
}


func r_p_s_game() {
	loader(200)
	plays := map[int]string{
		1: "Rock",
		2: "Paper",
		3: "Scissors",
	}
	fmt.Println("Welcome to the Rock-Paper-Scissors game.")
	
	for true {
		loader(300)
		fmt.Println("Choose:\n-- 1) Rock\n-- 2) Paper\n-- 3) Scissors\n\n\n[Enter '0' to exit.]\n")
		var answer int
		fmt.Scanln(&answer)
		if answer == 0 {
			break
		} else if plays[answer] == "" {
			continue
		}

		var randomized_play int = (rand.IntN(3 - 1) + 1)
		loader(300)
		if randomized_play == answer {
			fmt.Printf("\n%v and %v: Draw!!!\n", plays[answer], plays[answer])
		} else if randomized_play < answer || (answer == 3 && randomized_play == 1) {
			fmt.Printf("\n%v and %v: You won the game!!!\n", plays[randomized_play], plays[answer])
		} else {
			fmt.Printf("\n%v and %v: I won the game!!! \nStay Hard!\n", plays[randomized_play], plays[answer])
		}
	}

}

func from_celcius_to_far(value float64) float64 {
	return value * (9/5) + 32
}

func from_far_to_celcius(value float64) float64 {
	return (value - 32) * (9/5)
}

func main_menu() {
	for true {
		loader(200)
		var index int  
		fmt.Println("Enter the index of one of the programs:\n - 1) Guess and Win\n - 2) Temperature Converter\n - 3) Rock-Paper-Scissors Game")
		fmt.Scanln(&index)

		switch index {
			case 1:
				guesser_game()
			case 2:
				temperature_converter()
			case 3:
				r_p_s_game()
			default:
				fmt.Println("Undefined index. Please try again.")
		}
	}
}

func main() {
	main_menu()
}