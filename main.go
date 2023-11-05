package main

import (
	"github.com/spf13/cobra"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func main() {
	initCommandFlags()
	pghCmd.Execute()
}

func initCommandFlags()  {
	pghCmd.PersistentFlags().StringVar(&hostname, "a", "localhost", "")
	pghCmd.PersistentFlags().StringVar(&port, "p", "5432", "")
	pghCmd.PersistentFlags().StringVar(&username, "U", "", "")
	pghCmd.PersistentFlags().StringVar(&password, "P", "", "")
	pghCmd.PersistentFlags().StringVar(&database, "d", "", "")

	getCmd.AddCommand(tablesCmd)
	getCmd.AddCommand(seqCmd)
	pghCmd.AddCommand(getCmd)
}

var ( 
	hostname string
	port string
	username string
	password string
	database string
)

var pghCmd = &cobra.Command{
	Use:   "pgh",
	Short: "Simple utility for interactive with PostgreSQL",
	Long:  "Simple utility for interactive with PostgreSQL",
}

var getCmd = &cobra.Command{
	Use:   "get",
}

var tablesCmd = &cobra.Command{
	Use:   "tables",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT * FROM pg_catalog.pg_tables")

		if err != nil {
			panic(err)
		}

		for tables.Next() {
			var name string
			if err := tables.Scan(&name); err != nil {
				panic(err)
			}
			fmt.Printf("this is something: %s\n", name)
		}

	},
}

var seqCmd = &cobra.Command{
	Use:   "sequences",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT * FROM pg_catalog.pg_sequences")

		if err != nil {
			panic(err)
		}

		for tables.Next() {
			var name string
			if err := tables.Scan(&name); err != nil {
				panic(err)
			}
			fmt.Printf("this is something: %s\n", name)
		}

	},
}
