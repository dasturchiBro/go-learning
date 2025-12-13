package services

import (
	"github.com/xuri/excelize/v2"
	"xlsx/models"
	"strings"
	"errors"
	"strconv"
	"fmt"
)

func header(class, quarter, examType string) string {
	header := fmt.Sprintf("Pop tumani ixtisoslashtirilgan maktabi %v-sinf o'quvchilarining %v-chorak %v natijalari", class, quarter, examType)
	return header
}

func colName(i int) string {
	col, _ := excelize.ColumnNumberToName(i)
	return col
}

func BuildXLSX(req models.XLSXRequest) (*excelize.File, error) {
	f := excelize.NewFile()
	if err := f.SetCellValue("Sheet1", "A1", header(req.Header[0], req.Header[1], req.Header[2])); err != nil {
		return nil, err
	}

	if err := f.SetCellValue("Sheet1", "A5", "№"); err != nil {
		return nil, err
	}

	if err := f.SetCellValue("Sheet1", "B5", "F.I.SH"); err != nil {
		return nil, err
	}

	totalPoints := 0
	criteriaStart := 3
	for index, criterion := range req.Criteria {

		col := colName(index + criteriaStart)
		criterion = strings.TrimSpace(criterion)
		parts := strings.Split(criterion, " ")
		if len(parts) != 2 {
			return nil, errors.New("point was not assigned to a criterion")
		} 
		point, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		totalPoints += point
		if err := f.SetCellValue("Sheet1", col+"5", criterion); err != nil {
			return nil, err
		}
	}
	if err := f.SetCellValue("Sheet1", colName(criteriaStart + len(req.Criteria)) + "5", "Jami: " + strconv.Itoa(totalPoints)); err != nil {
		return nil, err
	}
	if err := f.SetCellValue("Sheet1", colName(criteriaStart + len(req.Criteria) + 1) + "5", "Foiz"); err != nil {
		return nil, err
	}
	if err := f.MergeCell("Sheet1", "A1", colName(criteriaStart + len(req.Criteria) + 1)+"4"); err != nil {
		return nil, err
	}

	// // START - Adding students //
	// for index, student := range req.Students {
	// 	if len(req.Criteria) != len(student[1:]) {
	// 		return nil, errors.New("number of criteria should be the same")
	// 	}
	// 	if err := f.SetCellValue("Sheet1", "A" + strconv.Itoa(2+(index + 1)), index+1); err != nil {
	// 		return nil, err
	// 	}
	// 	if err := f.SetCellValue("Sheet1", "B" + strconv.Itoa(2+(index + 1)), student[0]); err != nil {
	// 		return nil, err
	// 	}
	// 	for index, c := range student[1:] {
	// 		col := colName(index + criteriaStart)
	// 		if err := f.SetCellValue("Sheet1", col + strconv.Itoa(2+(index + 1)), c); err != nil {
	// 			return nil, err
	// 		}
	// 	}
		
	// }
	// // END - Adding students //

	// Style Start
	if err := f.SetRowHeight("Sheet1", 5, 40); err != nil {
		return nil, err
	}
	if err := f.SetColWidth("Sheet1", "A", colName(criteriaStart + len(req.Criteria) + 1), 20); err != nil {
		return nil, err
	}
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
	if err != nil {
		return nil, err
	}
	lastCellNumber := strconv.Itoa(5 + len(req.Students))
	f.SetCellStyle("Sheet1", "A1", colName(criteriaStart + len(req.Criteria) + 1) + lastCellNumber, style)

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
	if err != nil {
		return nil, err
	}
	f.SetCellStyle("Sheet1", "A5", colName(criteriaStart + len(req.Criteria) + 1)+"5", headerStyle)
	// Style End
	return f, nil		
}