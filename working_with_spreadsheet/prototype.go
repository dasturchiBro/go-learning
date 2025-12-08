package main

import (
	"log"
	"github.com/xuri/excelize/v2"
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)


func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Header() string {
	fmt.Println("Enter the class: ")
	class, _ := reader.ReadString('\n')

	fmt.Println("Enter the quarter(chorak): ")
	quarter, _ := reader.ReadString('\n')

	fmt.Println("Enter the exam type(e.g 1-BSB, CHSB, 2-BSB, etc.): ")
	examType, _ := reader.ReadString('\n')
	header := fmt.Sprintf("POP tuman IM %v-sinf o'quvchilarining %v-chorak %v natijalari", class, quarter, examType)
	return header
}

func SetCriteria() map[string]int {
	criteria := make(map[string]int) 
	for {
		fmt.Print("How many criteria(Vocabulary 5, Writing 8, Grammar 9, etc.) do you want to add? \nEnter the number:  ")
		numberOfCriteria_string, _ := reader.ReadString('\n')
		numberOfCriteria_string = strings.TrimSpace(numberOfCriteria_string)
		numberOfCriteria, err := strconv.Atoi(numberOfCriteria_string)
		Check(err)
		fmt.Println()
		if numberOfCriteria > 6 {
			fmt.Println("You can add up to 6 criteria. No more.\nPlease try again.")
			continue
		}
		for i := range numberOfCriteria {
			
			fmt.Print(strconv.Itoa(i+1)+") Please enter a criterion in this form: Criterion Point (e.g Writing 9).\nEnter: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			parts := strings.Split(input, " ")
			if len(parts) != 2 {
				log.Fatal(parts, "You entered the criterion in a wrong format. The format should be like Vocabulary 8 - Criterion Point.")
			} 
			point, err := strconv.Atoi(parts[1])
			Check(err)
			criteria[parts[0]] = point
		}
		break
	}
	return criteria
}

func AddStudents(f *excelize.File) {
	message := "How many students do you want to add? (enter in numbers)\nEnter: "
	fmt.Print(message)
	
	numberOfStudents, _ := reader.ReadString('\n')
	fmt.Println()
	number, err := strconv.Atoi(strings.TrimSpace(numberOfStudents))
	Check(err)
	for i := range number {
		fmt.Printf("Enter the name of the student on number %d: ", i+1)
		input, _ := reader.ReadString('\n')
		rowNum := strconv.Itoa(i+6)
		f.SetCellValue("Sheet1", "A"+rowNum, i+1)
		f.SetCellValue("Sheet1", "B"+rowNum, input)
	}
	fmt.Print("Do you want to add scores to the students you added? \nY/N:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "Y", "y", "YES", "yes":
		AddScoresToStudents(f)
	default:
		return
	}
}
func AddScoresToStudents(f *excelize.File) {
	start_index := 6

	for {
		indexOfRow := strconv.Itoa(start_index)
		cell, err := f.GetCellValue("Sheet1", "B"+indexOfRow)
		Check(err)
		cell = strings.TrimSpace(cell)
		if cell == "" {
			return
		}
		fmt.Printf("What scores would you like to add to %s?\nEnter (in this form: 10 6 9 ...): ", cell)
		scores, _ := reader.ReadString('\n')
		scores = strings.TrimSpace(scores)
		parts := strings.Split(scores, " ")
		letters := []string{"C", "D", "E", "F", "G", "H", "I", "J"}
		total := 0
		for part := range len(parts) {
			Check(f.SetCellValue("Sheet1", letters[part]+indexOfRow, parts[part]))
			num, err := strconv.Atoi(parts[part])
			Check(err)
			total += num
		}
		totalVal, err := f.GetCellValue("Sheet1", letters[len(parts)]+"5")
		Check(err)
		totalVal = strings.TrimSpace(totalVal)
		totalValue := strings.Split(totalVal, " ")[1]
		totalValue_int, err := strconv.Atoi(totalValue)
		Check(err)
		fmt.Println(totalValue_int)
		percent := fmt.Sprintf("%.2f", (float64(total)/float64(totalValue_int)) * 100)
		Check(f.SetCellValue("Sheet1", letters[len(parts)+1]+indexOfRow, (percent)))
		Check(f.SetCellValue("Sheet1", letters[len(parts)]+indexOfRow, total))
		start_index += 1
	}
}
func AddScoresToStudent(f *excelize.File) {}

func Body(f *excelize.File) {
	methods := "Choose one of these methods: \n1) Add students\n2) Add scores to a student\n3) Add scores to students\n0) Exit\n"
	for {
		fmt.Println(methods)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		switch answer {
		case "1":
			AddStudents(f)
		case "2":
			AddScoresToStudent(f)
		case "3":
			AddScoresToStudents(f)
		case "0":
			return
		default:
			fmt.Println("Undefined method. Try again. Method you typed: ", answer)
		}
	}
}

func main() {
	f := excelize.NewFile()
	defer func() {
		err := f.Close()
		Check(err)
	}()
	// Set Up Header
	Check(f.SetCellValue("Sheet1", "A1", Header()))
	Check(f.SetCellValue("Sheet1", "A5", "№"))
	Check(f.SetCellValue("Sheet1", "B5", "F.I.SH"))
	letters := []string{"C", "D", "E", "F", "G", "H", "I", "J"}
	criteria := SetCriteria()
	index := -1
	points := 0
	for key, value := range criteria {
		index+=1
		points += value
		Check(f.SetCellValue("Sheet1", letters[index]+"5", key+" "+strconv.Itoa(value)))
	}
	Check(f.SetCellValue("Sheet1", letters[index+1]+"5", "Jami: " + strconv.Itoa(points)))
	Check(f.SetCellValue("Sheet1", letters[index+2]+"5", "Foiz"))
	Check(f.MergeCell("Sheet1", "A1", letters[index+2]+"4"))
	// End Set Up Header


	// START Body
	Body(f)
	// END Body


	// Style Start
	Check(f.SetRowHeight("Sheet1", 5, 40))
	Check(f.SetColWidth("Sheet1", "A", "H", 20))
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical: "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
		},
	})
	Check(err)
	f.SetCellStyle("Sheet1", "A1", letters[index+2]+"15", style)

	headerStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type: "pattern",
			Color: []string{"BDD7EE"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical: "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold: true,
		},
	})
	Check(err)
	f.SetCellStyle("Sheet1", "A5", letters[index+2]+"5", headerStyle)
	// Style End

	Check(f.SaveAs("template.xlsx"))
}