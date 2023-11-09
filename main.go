package main

import (
	"github.com/spf13/cobra"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func main() {
	initCommands()
	pghCmd.Execute()
}

func initCommands()  {
	pghCmd.PersistentFlags().StringVar(&hostname, "a", "localhost", "")
	pghCmd.PersistentFlags().StringVar(&port, "p", "5432", "")
	pghCmd.PersistentFlags().StringVar(&username, "U", "", "")
	pghCmd.PersistentFlags().StringVar(&password, "P", "", "")
	pghCmd.PersistentFlags().StringVar(&database, "d", "", "")

	topCmd.AddCommand(topQueriesCmd)
	getCmd.AddCommand(matViewsCmd)
	getCmd.AddCommand(tablesCmd)
	getCmd.AddCommand(indexCmd)
	getCmd.AddCommand(viewsCmd)
	getCmd.AddCommand(locksCmd)
	getCmd.AddCommand(seqCmd)
	pghCmd.AddCommand(getCmd)
	pghCmd.AddCommand(topCmd)
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

var topCmd = &cobra.Command{
	Use:   "top",
}

var topQueriesCmd = &cobra.Command{
	Use:   "queries",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		slowQueries, err := db.Query("SELECT * FROM pg_stat_statements JOIN pg_roles r ON r.oid = userid WHERE calls > 100 ORDER BY mean_time DESC")

		if err != nil {
			panic(err)
		}

		for slowQueries.Next() {
			var name string
			if err := slowQueries.Scan(&name); err != nil {
				panic(err)
			}
			fmt.Printf("this is something: %s\n", name)
		}

	},
}

var tablesCmd = &cobra.Command{
	Use:   "tables",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT tablename FROM pg_catalog.pg_tables")

		if err != nil {
			panic(err)
		}

		fmt.Printf("Tables list: \n")
		fmt.Printf("--------------------- \n")

		for tables.Next() {
			var name string
			if err := tables.Scan(&name); err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", name)
		}

	},
}

var seqCmd = &cobra.Command{
	Use:   "sequences",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT sequencename FROM pg_catalog.pg_sequences")

		if err != nil {
			panic(err)
		}

		fmt.Printf("Sequences list: \n")
		fmt.Printf("--------------------- \n")

		for tables.Next() {
			var name string
			if err := tables.Scan(&name); err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", name)
		}

	},
}

var indexCmd = &cobra.Command{
	Use:   "indexes",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT * FROM pg_catalog.pg_indexes")

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

var viewsCmd = &cobra.Command{
	Use:   "views",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT * FROM pg_catalog.pg_views")

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

var matViewsCmd = &cobra.Command{
	Use:   "matviews",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT * FROM pg_catalog.pg_matviews")

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

var locksCmd = &cobra.Command{
	Use:   "locks",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT * FROM pg_catalog.pg_locks")

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
