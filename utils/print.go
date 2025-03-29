package utils

import (
	"bufio"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"strings"
)

type Params struct {
	Username string
	Password string
	Host     string
	Port     string
	DB_name  string
}

func GetParamsFromUser(p *Params) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	Username, _ := reader.ReadString('\n')
	p.Username = trimNewline(Username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	p.Password = trimNewline(password)

	fmt.Print("Enter host: ")
	host, _ := reader.ReadString('\n')
	p.Host = trimNewline(host)

	fmt.Print("Enter port: ")
	port, _ := reader.ReadString('\n')
	p.Port = trimNewline(port)

	fmt.Print("Enter database name: ")
	dbName, _ := reader.ReadString('\n')
	p.DB_name = trimNewline(dbName)

	return nil // No specific error handling in this basic example
}

// Helper function to remove the trailing newline character from input
func trimNewline(s string) string {
	return strings.TrimSuffix(s, "\n")
}
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

func ReadLines() string {
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		line, err := reader.ReadString('\n') //this reads only one read
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		line_s := strings.Split(line, " ")
		lines = append(lines, line_s...)
	}
	return strings.Join(lines, " ")
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
