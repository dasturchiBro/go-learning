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

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Header() string {
	var class, quarter, examType string
	fmt.Println("Enter the class: ")
	fmt.Scanln(&class)

	fmt.Println("Enter the quarter(chorak): ")
	fmt.Scanln(&quarter)

	fmt.Println("Enter the exam type(e.g 1-BSB, CHSB, 2-BSB, etc.): ")
	fmt.Scanln(&examType)
	header := fmt.Sprintf("POP tuman IM %v-sinf o'quvchilarining %v-chorak %v natijalari", class, quarter, examType)
	return header
}

func SetCriteria() map[string]int {
	criteria := make(map[string]int) 
	for {
		fmt.Print("How many criteria(Vocabulary 5, Writing 8, Grammar 9, etc.) do you want to add? \nEnter the number:  ")
		var numberOfCriteria int
		fmt.Scanln(&numberOfCriteria)
		fmt.Println()
		if numberOfCriteria > 6 {
			fmt.Println("You can add up to 6 criteria. No more.\nPlease try again.")
			continue
		}
		for i := range numberOfCriteria {
			reader := bufio.NewReader(os.Stdin)
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