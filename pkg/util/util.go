package util

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func getRowsFromDb(hostname string, port string, username string, password string, database string, columns string, table string) *sql.Rows  {
	var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
	db, err := sql.Open("postgres", connString)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT" + columns + " FROM " + table)

	if err != nil {
		panic(err)
	}

	return rows
}