package utils

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"strings"
)

func PrintRowsAsTable(rows pgx.Rows) {
	fieldDescriptions := rows.FieldDescriptions()
	columnNames := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columnNames[i] = fd.Name
	}
	numCols := len(columnNames)

	var allRows [][]string
	columnWidths := make([]int, numCols)

	for i, name := range columnNames {
		columnWidths[i] = len(name)
	}

	for rows.Next() {
		values := make([]any, numCols)
		valuePtrs := make([]any, numCols)
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatal(err)
		}

		rowData := make([]string, numCols)
		for i, v := range values {
			rowData[i] = fmt.Sprintf("%v", v)
			width := len(rowData[i])
			if width > columnWidths[i] {
				columnWidths[i] = width
			}
		}
		allRows = append(allRows, rowData)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	printHeaderDynamic(columnNames, columnWidths)
	printSeparatorDynamic(columnWidths)

	for _, rowData := range allRows {
		printRowDynamic(rowData, columnWidths)
	}
}

func printHeaderDynamic(columnNames []string, columnWidths []int) {
	for i, name := range columnNames {
		fmt.Printf("%-*s  ", columnWidths[i], name)
	}
	fmt.Println()
}

func printSeparatorDynamic(columnWidths []int) {
	for _, width := range columnWidths {
		fmt.Printf("%s  ", strings.Repeat("-", width))
	}
	fmt.Println()
}

func printRowDynamic(rowData []string, columnWidths []int) {
	for i, data := range rowData {
		fmt.Printf("%-*s  ", columnWidths[i], data)
	}
	fmt.Println()
}
