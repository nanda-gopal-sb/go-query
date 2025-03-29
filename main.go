package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go-query/utils"
	"log"
	"os"
)

func connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:<password>du@localhost:5432/college_db")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func queryData(conn *pgx.Conn, query string) {
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()
	utils.PrintRowsAsTable(rows)
}
func showOptions(choice *int) {
	fmt.Println("Welcome!")
	fmt.Println("1 - Show all tables")
	fmt.Println("2 - Make a query")
	fmt.Println("3 - Show query history")
	fmt.Println("4 - Clear Screen")
	fmt.Scanln(choice)
}
func main() {
	var database *pgx.Conn
	reader := bufio.NewReader(os.Stdin)
	var choice int = 0
	var err error
	database, err = connect()
	if err != nil {
		fmt.Println("An error has occured while opening the database")
		return
	}
	for true {
		showOptions(&choice)
		fmt.Println(choice)
		if choice == 2 {
			fmt.Println("Enter your query")
			query, error := reader.ReadString('\n')
			if error != nil {
				fmt.Println("Error reading input:", error)
				continue
			}
			queryData(database, query)
			fmt.Println("")
		} else if choice == 1 {

			queryData(database, "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
		} else if choice == 4 {
			fmt.Print("\033[H\033[2J") //clears the screen, need to find concrete solution
		}
	}
}
