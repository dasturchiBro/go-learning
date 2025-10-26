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
	avg_score float64
	max_score float64
	min_score float64
	scores []float64
}



func all_students(students []Student) {
	for index, value := range students {
		fmt.Printf("%d) %v \n\tAge: %d\n\tAverage Score: %.2f\n\tMax Score: %.2f\n\tMin Score: %.2f\n\tID: %d\n\n",index+1, value.name, value.age, value.avg_score, value.max_score, value.min_score, value.id)
		time.Sleep(300 * time.Millisecond)
	}
}

func statistics(students []Student) {
	var average_score float64
	var sum_of_average_scores float64
	for _, value := range students {
		sum_of_average_scores += value.avg_score
	}
	average_score = sum_of_average_scores / float64(len(students))
	var students_above_avg string  
	var students_below_avg string
	for _, value := range students {
		if value.avg_score >= average_score {
			students_above_avg += value.name + "\n"
		} else {
			students_below_avg += value.name + "\n"
		}
	}
	loader(300)
	message := "*Better-performing students*\n" + students_above_avg + "\n\n" + "*Worse-performing students*\n" + students_below_avg
	fmt.Println(message)
	loader(300)
}

func enter_results(students []Student) {
	for true {
		message := "*Enter the student's ID you want to add scores*\n"
		fmt.Println(message)
		loader(500)
		all_students(students)
		fmt.Println("\n\n\n[Enter '0' to go back]")
		var user_id int 
		fmt.Print("\t\t")
		fmt.Scanln(&user_id)
		loader(100)
		if user_id == 0 {
			break
		}
		isThereUser := false
		for i, value := range students {
			if value.id == user_id {
				isThereUser = true 
				newVal, isHome := enter_results_user_found(value)
				if isHome == 1 {
					return
				} 
				students[i] = newVal
				loader(300)
				fmt.Println("***Scores Added Successfully***")
				loader(500)
				return
			}
		}
		if !isThereUser {
			fmt.Println("A student with this ID doesn't exist. Please try again.")
			loader(300)
		}
	}
}

func change_arrstr_to_float(items []string, length int) ([]float64, int) {
	newArr := []float64{}
	for _, value := range items {
		value, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return []float64{0.0}, 1
		}
		newArr = append(newArr, value)
	}
	return newArr, 0
}

func enter_results_user_found(value Student) (Student, int) {
	for true {
		message := "Enter the results in this form: 12.4 89.4 12.3 45.3 ...\n\n\n[Enter '0' to go back]"
		fmt.Println(message)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := scanner.Text()
		if answer == "0" {
			return value, 1
		}
		items := strings.Split(answer, " ")
		values, indicator := change_arrstr_to_float(items, len(items))
		if indicator == 1 {
			fmt.Println("Warning: All values should be numbers! Please try again.")
			continue
		}
		value.scores = append(value.scores, values...)
		var scores_sum float64
		max_score, min_score := value.scores[0], value.scores[0] 
		for _, score := range value.scores {
			scores_sum += score
			if max_score < score {max_score = score} 
			if min_score > score {min_score = score} 
		}
		value.avg_score = scores_sum / float64(len(value.scores))
		value.min_score = min_score
		value.max_score = max_score
		return value, 0
	}
	return value, 0
}



func add_students(students *[]Student) {
	fmt.Println("Enter the name of the student you want to add: \n\n[Enter '0' to go back]")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	if name == "0" {
		return
	}
	var newStudent Student
	newStudent.name = name 
	fmt.Println("Enter the age of the student: [Enter '0' to go back]")
	var age int  
	fmt.Scanln(&age)
	if age == 0 {
		return
	}
	newStudent.age = age 
	newStudent.id = (*students)[len(*students) - 1].id + 1
	*students = append(*students, newStudent)
	fmt.Println("***Student Added Successfully**")

	loader(500)
}

func get_student_by_id(students []Student) {
	for true {
		fmt.Println("Enter the id of a student you want to see: \n\n[Enter '0' to go back]")
		var id int
		fmt.Scanln(&id)

		if id == 0 {
			return
		}
		var isThereStudent bool = false
		for _, student := range students {
			if student.id == id {
				isThereStudent = true
				loader(200)
				fmt.Printf("%v \n\tAge: %d\n\tAverage Score: %.2f\n\tMax Score: %.2f\n\tMin Score: %.2f\n\tID: %d\n\n",student.name, student.age, student.avg_score, student.max_score, student.min_score, student.id)
				return
			}
		}
		if !isThereStudent {
			fmt.Println("Student by ID not found. Please try again.\n\n")
			loader(400)
		}
	}
}

func remove_student_by_id(students *[]Student) {
	for true {
		fmt.Println("Enter the id of a student you want to remove: \n\n[Enter '0' to go back]")
		var id int
		fmt.Scanln(&id)

		if id == 0 {
			return
		}
		var isThereStudent bool = false
		for index, student := range *students {
			if student.id == id {
				isThereStudent = true
				loader(200)
				fmt.Printf("%v \n\tAge: %d\n\tAverage Score: %.2f\n\tMax Score: %.2f\n\tMin Score: %.2f\n\tID: %d\n***Student was removed successfully***\n\n\n",student.name, student.age, student.avg_score, student.max_score, student.min_score, student.id)
				*students = append((*students)[:index], (*students)[index+1:]...)
				loader(500)
				return
			}
		}
		if !isThereStudent {
			fmt.Println("Student by ID not found. Please try again.")
			loader(400)
			fmt.Print("\n\n")
		}
	}
}


func launch_student_grades_tracker() {
	students := []Student{
		{
			5,
			"Noah Miller",
			18,
			890,
			3.55,
			85.6,
			[]float64{83.2, 88.1, 85.0, 87.9, 84.5, 86.6},
		},
		{
			6,
			"Chloe Bellwether",
			17,
			432,
			3.95,
			95.0,
			[]float64{94.5, 96.0, 93.8, 97.1, 95.5, 92.9},
		},
		{
			7,
			"Ethan Hunt",
			19,
			765,
			3.30,
			79.2,
			[]float64{80.5, 75.5, 82.0, 77.1, 78.9, 81.5},
		},
		{
			8,
			"Ava Sharma",
			18,
			109,
			3.70,
			88.8,
			[]float64{87.0, 90.5, 89.1, 86.8, 91.0, 88.5},
		},
		{
			9,
			"Marcus Jones",
			20,
			543,
			2.50,
			55.9,
			[]float64{58.0, 50.5, 60.1, 52.5, 59.9, 53.0},
		},
		{
			10,
			"Isabelle Tran",
			17,
			987,
			3.60,
			86.1,
			[]float64{85.5, 87.0, 84.5, 88.1, 83.9, 87.5},
		},
	}
	for true {
		fmt.Println("Welcome to the Student Grades Tracker program. Choose options below: \n+++(a) See all students\n+++(b) See statistics\n+++(c) Enter results\n+++(d) Add student\n+++(e) Get Student by ID\n+++(f) Remove Student by ID\n\n\n[Enter 'g' to go back.]")
		var answer string  
		fmt.Scanln(&answer)
		if answer == "a" {
			all_students(students)
		} else if answer == "b" {
			statistics(students)
		} else if answer == "c" {
			enter_results(students)
		} else if answer == "d" {
			add_students(&students)
		} else if answer == "e" {
			get_student_by_id(students)
		} else if answer == "f" {
			remove_student_by_id(&students)
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

