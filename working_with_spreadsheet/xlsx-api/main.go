package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type XLSXRequest struct {
	Headers []string `json:"headers"`
	Rows [][]any `json:"rows"`
}


func ColName(i int) string {
	col, _ := excelize.ColumnNumberToName(i+1)
	return col
}

func BuildXLSX(req XLSXRequest) (*excelize.File, error) {
	f := excelize.NewFile()

	// ["Name", "Age"]
	for i, _ := range req.Headers {
		cell := ColName(i) + "1"
		err := f.SetCellValue("Sheet1", cell, req.Headers[i])
		if err != nil {
			return nil, err
		}
	}

	/* 
	rows = [
		["John", 15], 
		["Sherlock", 31],
	]
	*/
	for row, val := range req.Rows {
		for column := range val {
			cell := ColName(column) + strconv.Itoa(row+2)
			err := f.SetCellValue("Sheet1", cell, req.Rows[row][column])
			if err != nil {
				return nil, err
			}
		}
	}
	return f, nil
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func xlsxHandler(c *gin.Context) {
	var req XLSXRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	f, err := BuildXLSX(req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}
	defer f.Close()
	if err := f.SaveAs("new.xlsx"); err != nil {
		c.JSON(400, gin.H{"error": "Couldn't save the file. Please try again."})
		return
	}

	c.JSON(200, req)
}

func main() {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	v1 := r.Group("/v1")
	{
		v1.POST("/xlsx", xlsxHandler)
	}

	r.Run(":8080")
}