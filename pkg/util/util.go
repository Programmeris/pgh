package util

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type ConnectionOptions struct{
	hostname string
	port string
	username string
	password string
	database string
}

func NewConnectionOptions(hostname string, port string, username string, password string, database string) *ConnectionOptions {
	connectionOptions := new(ConnectionOptions)
	connectionOptions.hostname = hostname
	connectionOptions.database = database
	connectionOptions.username = username
	connectionOptions.password = password
	connectionOptions.port = port

	return connectionOptions

}

func buildConnectionString(connectionOptions *ConnectionOptions) string {
	return "host=" + connectionOptions.hostname + " port=" + 
			connectionOptions.port + " user=" + connectionOptions.username + 
			" password=" + connectionOptions.password + " dbname=" + 
			connectionOptions.database + " sslmode=disable"
}

func GetRowsFromTable(connectionOptions ConnectionOptions, columns string, table string) *sql.Rows  {
	var connString = buildConnectionString(&connectionOptions)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT" + columns + " FROM " + table)

	if err != nil {
		panic(err)
	}

	return rows
}

func PrintRows(rows *sql.Rows, title string){

	fmt.Printf(title + " list: \n")
	fmt.Printf("--------------------- \n")

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", name)
	}

}
