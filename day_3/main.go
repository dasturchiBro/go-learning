// "To-Do List" - "Student Grades Tracker"

/* 
How does a simple to-do list work?
Main:
	Items
	a) Add
	b) Remove


*/
package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	"strconv"
)

func loader(t time.Duration) {
	time.Sleep(t * time.Millisecond)
	fmt.Print(".")

	time.Sleep(t * time.Millisecond)
	fmt.Print(".")

	time.Sleep(t * time.Millisecond)
	fmt.Print(".")
	time.Sleep((t+200) * time.Millisecond)
	fmt.Print("\n")
}

func read_file(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(filename)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return string(data)
}

func write_file(filename, message, SPLIT_KEY string) {
	file, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(SPLIT_KEY + message); err != nil {
		panic(err)
	}
}

func to_do_add(filename, SPLIT_KEY string) {
	for true {
		fmt.Println("What would you like to add to your To-Do list?\n\n[Enter 'g' to go back].")
		var answer string 
		fmt.Scanln(&answer)
		if answer == "g" {
			launch_to_do()
			loader(200)
			return
		} else if answer == "" {
			continue
		} else {
			write_file(filename, answer, SPLIT_KEY)
			loader(400)
			fmt.Println("Task added successfully!")
			loader(200)
			launch_to_do()
			return
		}
	}
}

func launch_to_do() {
	for true {
		SPLIT_KEY := "#$!@((SPLIT_KEY_876543#$%^&*"
		filename := "to_do.txt"
		data := read_file(filename)
		items := strings.Split(data, SPLIT_KEY)
		main_message := "--To-Do List--\n"
		if items[0] != "" {
			for index, value := range items {
				main_message += (strconv.Itoa(index+1) + ") " + value + "\n")
			}
			main_message += "___a) Add\n___b) Remove\n"
		} else {
			main_message += "___a) Add\n"
		}
		main_message += "___m) Main Menu\n"
		fmt.Println(main_message)
		
		var answer string  
		fmt.Scanln(&answer)
		if answer == "m" {
			loader(300)
			main()
		} else if answer == "a" {
			to_do_add(filename, SPLIT_KEY)
		}
		
	}
}

func main() {
	greetings_message := "Welcome to the world of programming stuff!\n\n\nChoose one of these programs:\n1) To-Do List\n2) Student Grades Tracker\n\n\n[Enter '0' to exit.]"
	var answer int
	for true {
		fmt.Println(greetings_message)
		fmt.Scanln(&answer)
		loader(200)
		if answer == 0 {
			fmt.Println("Goodbye!!!")
			loader(200)
			break	
		} else if answer == 1 {
			launch_to_do()
		}
	}
}


// Today: To-Do list addition finished only. Gotta finish function Remove and create the next program... loader(8.64e+7)