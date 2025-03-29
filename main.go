package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go-query/utils"
	"log"
)

func connect(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
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
func getNewConnection() (newString string) {
	newConn := &utils.Params{}
	utils.GetParamsFromUser(newConn)
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", newConn.Username, newConn.Password, newConn.Host, newConn.Port, newConn.DB_name)
	return connString
}
func showOptions(choice *int) {
	fmt.Println("Welcome!")
	fmt.Println("1 - Show all tables")
	fmt.Println("2 - Make a query")
	fmt.Println("3 - Show query history")
	fmt.Println("4 - Clear Screen")
	fmt.Println("5 - Add a connection")
	fmt.Scanln(choice)
}
func main() {
	var database *pgx.Conn
	var choice int = 0
	var err error
	connString := getNewConnection()
	database, err = connect(connString)
	if err != nil {
		fmt.Println("An error has occured while opening the database")
		return
	}
	for true {
		showOptions(&choice)
		if choice == 2 {
			fmt.Println("Enter your query")
			query := utils.ReadLines()
			queryData(database, query)
			fmt.Println("")
		} else if choice == 1 {
			fmt.Println("")
			queryData(database, "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
			fmt.Println("")
		} else if choice == 4 {
			fmt.Print("\033[H\033[2J") //clears the screen, need to find concrete solution
		} else if choice == 5 {
		}

	}
}
