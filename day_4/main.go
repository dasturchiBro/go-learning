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

func write_file_not_append(filename, message string) {
	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(message); err != nil {
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
			return
		}
	}
}

func launch_to_do() {
	for true {
		SPLIT_KEY := "\n"
		filename := "to_do.txt"
		data := read_file(filename)
		items := strings.Split(data, SPLIT_KEY)
		is_there_items := true
		main_message := "--To-Do List--\n"
		if items[0] != "" {
			for index, value := range items {
				main_message += (strconv.Itoa(index+1) + ") " + value + "\n")
			}
			main_message += "+a) Add\n+r) Remove\n"
		} else {
			main_message += "+a) Add\n"
			is_there_items = false
		}
		main_message += "+m) Main Menu\n"
		fmt.Println(main_message)
		
		var answer string  
		fmt.Scanln(&answer)
		if answer == "m" {
			loader(300)
			main()
		} else if answer == "a" {
			to_do_add(filename, SPLIT_KEY)
		} else if answer == "r" && is_there_items == true {
			fmt.Print("Enter the index of the task to delete: \n\n[Enter '0' to go back.]\n\t\t")
			var index_of_item int 
			fmt.Scanln(&index_of_item)
			loader(300)
			index_of_item = index_of_item - 1
			if index_of_item >= 0 && index_of_item < len(items) {
				items = append(items[:index_of_item], items[index_of_item + 1:]...)
				var message string
				for _, value := range items {
					message += value + SPLIT_KEY
				}
				write_file_not_append(filename, message)
			} else if index_of_item == -1 {
				continue
			} else {
				fmt.Println("The task doesn't exist.")
				loader(300)
			}
		}
	}
}

func main() {
	greetings_message := "Welcome to the world of programming stuff!\n\n\nChoose one of these programs:\n1) To-Do List\n2) Student Grades Tracker\n\n\n[Enter '-1' to exit.]"
	for true {
		fmt.Println(greetings_message)
		var answer int
		fmt.Scanln(&answer)
		loader(200)
		if answer == -1 {
			fmt.Println("Goodbye!!!")
			loader(200)
			break	
		} else if answer == 1 {
			launch_to_do()
		}
	}
}


// Today: To-Do list addition finished only. Gotta finish function Remove and create the next program... loader(8.64e+7)