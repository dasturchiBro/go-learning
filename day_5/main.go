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
	"bufio"
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
	file_data := read_file(filename)

	if file_data == "" {
		if _, err := file.WriteString(message); err != nil {
			panic(err)
		}
	} else {
		if _, err := file.WriteString(SPLIT_KEY + message); err != nil {
			panic(err)
		}
	}

}

func write_file_not_append(filename, message string) {
	file, err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		panic(err)
	}
}

func to_do_add(filename, SPLIT_KEY string) {
	for true {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("What would you like to add to your To-Do list?\n\n[Enter 'g' to go back].")
		scanner.Scan()
		answer := scanner.Text()
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
				for index_of_item, value := range items {
					message += value
					if index_of_item != len(items) - 1 {
						message += SPLIT_KEY
					}
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

type Student struct {
	id int
	name string 
	age int 
	avg_score float32
	max_score float32
	min_score float32
	scores []float32
}



func all_students(students []Student) {
	for index, value := range students {
		fmt.Printf("%d) %v \n\tAge: %d\n\tAverage Score: %.2f\n\tID: %d\n\n",index+1, value.name, value.age, value.avg_score, value.id)
		time.Sleep(300 * time.Millisecond)
	}
}

func statistics(students []Student) {
	var average_score float32
	var sum_of_average_scores float32
	for _, value := range students {
		sum_of_average_scores += value.avg_score
	}
	average_score = sum_of_average_scores / float32(len(students))
	var students_above_avg string  
	var students_below_avg string
	for _, value := range students {
		if value.avg_score >= average_score {
			students_above_avg += value.name + "\n"
		} else {
			students_below_avg += value.name + "\n"
		}
	}
	message := "*Better-performing students*\n" + students_above_avg + "\n\n" + "*Worse-performing students*\n" + students_below_avg
	fmt.Println(message)
}

func enter_results() {
	message := "*Choose the student you want to add scores*"
		
}

func add_students() {

}

func get_student_by_id() {

}

func remove_student_by_id() {
	
}


func launch_student_grades_tracker() {
	students := []Student{
		{
			1,
			"Fayozbek Erkinjonov",
			17,
			67,
			78.4,
			34.5,
			[]float32{14.5,43,23.67,56.7,87.5,98.3},
		},

		{
			2,
			"Kai Chen",
			19,
			345,
			3.15,
			75.3,
			[]float32{78.0, 72.5, 80.1, 69.5, 76.8, 81.1},
		},
		{
			3,
			"Liam Rodriguez",
			17,
			678,
			4.0,
			98.7,
			[]float32{99.0, 97.5, 98.2, 99.5, 96.0, 98.9},
		},
		{
			4,
			"Sofia Amani",
			20,
			211,
			2.90,
			65.4,
			[]float32{62.5, 70.0, 68.3, 60.1, 71.5, 64.9},
		},
		{
			5,
			"Noah Miller",
			18,
			890,
			3.55,
			85.6,
			[]float32{83.2, 88.1, 85.0, 87.9, 84.5, 86.6},
		},
		{
			6,
			"Chloe Bellwether",
			17,
			432,
			3.95,
			95.0,
			[]float32{94.5, 96.0, 93.8, 97.1, 95.5, 92.9},
		},
		{
			7,
			"Ethan Hunt",
			19,
			765,
			3.30,
			79.2,
			[]float32{80.5, 75.5, 82.0, 77.1, 78.9, 81.5},
		},
		{
			8,
			"Ava Sharma",
			18,
			109,
			3.70,
			88.8,
			[]float32{87.0, 90.5, 89.1, 86.8, 91.0, 88.5},
		},
		{
			9,
			"Marcus Jones",
			20,
			543,
			2.50,
			55.9,
			[]float32{58.0, 50.5, 60.1, 52.5, 59.9, 53.0},
		},
		{
			10,
			"Isabelle Tran",
			17,
			987,
			3.60,
			86.1,
			[]float32{85.5, 87.0, 84.5, 88.1, 83.9, 87.5},
		},
	}
	for true {
		fmt.Println("Welcome to the Student Grades Tracker program. Choose options below: \n+++(a) See all students\n+++(b) See statistics\n+++(c) Enter results\n+++(d) Add students\n+++(e) Get Student by ID\n+++(f) Remove Student by ID\n\n\n[Enter 'g' to go back.]")
		var answer string  
		fmt.Scanln(&answer)
		if answer == "a" {
			all_students(students)
		} else if answer == "b" {
			statistics(students)
		} else if answer == "c" {
			enter_results()
		} else if answer == "d" {
			add_students()
		} else if answer == "e" {
			get_student_by_id()
		} else if answer == "f" {
			remove_student_by_id()
		} else if answer == "g" {
			return
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
		} else if answer == 2 {
			launch_student_grades_tracker()
		}
	}
}

