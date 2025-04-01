package main

import (
	"context"
	"fmt"
	"go-query/utils"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/jackc/pgx/v5"
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
func getNewConnection(newConn utils.Params) (newString string) {
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
	newConn := &utils.Params{}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the host").
				Value(&newConn.Host),
			huh.NewInput().
				Title("Enter the port").
				CharLimit(400).
				Value(&newConn.Port),
			huh.NewInput().
				Title("Enter the Database Name").
				CharLimit(400).
				Value(&newConn.DB_name),
			huh.NewInput().
				Title("Enter the username").
				CharLimit(400).
				Value(&newConn.Username),
			huh.NewInput().
				Title("Enter the Password").
				CharLimit(400).
				Value(&newConn.Password),
		),
	)
	formErr := form.Run()
	if formErr != nil {
		log.Fatal(formErr)
	}
	var database *pgx.Conn
	var choice int = 0
	var err error
	connString := getNewConnection(*newConn)
	database, err = connect(connString)
	if err != nil {
		fmt.Println("An error has occured while opening the database")
		return
	}
	for true {
		var query string
		showOptions(&choice)
		if choice == 2 {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewText().
						Title("Enter your query").
						Value(&query).
						ShowLineNumbers(true),
				),
			)
			form.Run()
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
