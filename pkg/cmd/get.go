package cmd

import (
	"github.com/spf13/cobra"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

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

		tables, err := db.Query("SELECT indexname FROM pg_catalog.pg_indexes")

		if err != nil {
			panic(err)
		}

		fmt.Printf("Indexes list: \n")
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

var viewsCmd = &cobra.Command{
	Use:   "views",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT viewname FROM pg_catalog.pg_views")

		if err != nil {
			panic(err)
		}

		fmt.Printf("Views list: \n")
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

var matViewsCmd = &cobra.Command{
	Use:   "matviews",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT matviewname FROM pg_catalog.pg_matviews")

		if err != nil {
			panic(err)
		}

		fmt.Printf("Materialized views list: \n")
		fmt.Printf("--------------------------- \n")

		for tables.Next() {
			var name string
			if err := tables.Scan(&name); err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", name)
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

		tables, err := db.Query("SELECT locktype FROM pg_catalog.pg_locks")

		if err != nil {
			panic(err)
		}

		fmt.Printf("Locks list: \n")
		fmt.Printf("------------ \n")

		for tables.Next() {
			var name string
			if err := tables.Scan(&name); err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", name)
		}

	},
}