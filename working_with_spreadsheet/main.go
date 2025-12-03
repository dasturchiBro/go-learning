package main

import (
	"log"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f,err := excelize.OpenFile("book1.xlsx")
	check(err)
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	check(f.SetRowHeight("Sheet1", 5, 40))
	check(f.MergeCell("Sheet1", "A1", "H4"))

	f.SetCellValue("Sheet1", "A1", "POP tuman IM  o`quvchilarining III chorak 1-BSB natijalari Ingliz tili")
	f.SetCellValue("Sheet1", "A5", "№")
	f.SetCellValue("Sheet1", "B5", "F.I.SH")
	f.SetCellValue("Sheet1", "C5", "Vocabulary: 8")
	f.SetCellValue("Sheet1", "D5", "Reading: 8")
	f.SetCellValue("Sheet1", "E5", "Grammar: 8")
	f.SetCellValue("Sheet1", "F5", "Writing: 8")
	f.SetCellValue("Sheet1", "G5", "Jami: 25")
	f.SetCellValue("Sheet1", "H5", "Foizi")
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
	check(err)
	f.SetCellStyle("Sheet1", "A1", "H15", style)
	check(f.SetColWidth("Sheet1", "A", "H", 20))

	newStyle, err := f.NewStyle(&excelize.Style{
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
	check(err)
	check(f.SetCellStyle("Sheet1", "A5", "H5", newStyle))
	for i := range 10 {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+6), i+1)
		f.SetRowHeight("Sheet1", i+6, 20)
	}
	if err := f.SaveAs("book1.xlsx"); err != nil {
		log.Fatal(err)
	}
}